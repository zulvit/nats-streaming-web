package service

import (
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"nats-streaming-web/dbclient"
	"nats-streaming-web/internal/config"
	"nats-streaming-web/internal/logger"
	"nats-streaming-web/pkg/cache"
	"nats-streaming-web/pkg/client"
	"nats-streaming-web/pkg/model"
	"time"
)

type ClientOrderService struct {
	NatsClient  *client.NatsClient
	RedisClient *cache.RedisClient
	DBClient    *dbclient.DBClient
}

func NewClientOrderService(natsClient *client.NatsClient, redisClient *cache.RedisClient, dbClient *dbclient.DBClient) *ClientOrderService {
	return &ClientOrderService{
		NatsClient:  natsClient,
		RedisClient: redisClient,
		DBClient:    dbClient,
	}
}

func (s *ClientOrderService) SubscribeAndProcess(ctx context.Context, subject string) {
	s.NatsClient.Conn.Subscribe(subject, func(msg *stan.Msg) {
		var orderData model.OrderData
		err := json.Unmarshal(msg.Data, &orderData)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка десериализации данных: %v\n", err)
			return
		}

		_, err = s.DBClient.SaveOrderData(ctx, &orderData)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка при сохранении данных в базе данных: %v\n", err)

			cacheErr := s.RedisClient.Set(ctx, "order:"+orderData.OrderUID, msg.Data)
			if cacheErr != nil {
				logger.ErrorLogger.Printf("Ошибка при записи в Redis: %v\n", cacheErr)
			} else {
				logger.InfoLogger.Printf("Данные успешно закешированы в Redis для OrderUID: %s\n", orderData.OrderUID)
			}
			return
		}

		err = s.RedisClient.Set(ctx, "order:"+orderData.OrderUID, msg.Data)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка при записи в Redis: %v\n", err)
		} else {
			logger.InfoLogger.Printf("Данные успешно закешированы в Redis для OrderUID: %s\n", orderData.OrderUID)
		}

	}, stan.DurableName("my-durable"))
}

func (s *ClientOrderService) MonitorDBAndRestoreFromCache(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 10)
	logger.InfoLogger.Println("Проверяем успешность соединения с БД")
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if s.DBClient.IsAvailable() {
				s.restoreOrdersFromCache(ctx)
			}
		}
	}
}

func (s *ClientOrderService) restoreOrdersFromCache(ctx context.Context) {
	cfg := config.NewConfig()
	dbClient, err := dbclient.NewDBClient(cfg.DBConnectionString)
	if err != nil {
		logger.ErrorLogger.Println("Ошибка применения миграции БД")
	}

	dbClient.SaveOrderData(ctx, &model.OrderData{})

	keys, err := s.RedisClient.Keys(ctx, "order:*")
	if err != nil {
		logger.ErrorLogger.Printf("Ошибка при получении ключей из Redis: %v\n", err)
		return
	}

	for _, key := range keys {
		data, err := s.RedisClient.Get(ctx, key)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка при получении данных из Redis для ключа %s: %v\n", key, err)
			continue
		}

		var orderData model.OrderData
		err = json.Unmarshal([]byte(data), &orderData)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка десериализации данных заказа: %v\n", err)
			continue
		}

		_, err = s.DBClient.SaveOrderData(ctx, &orderData)
		if err != nil {
			logger.ErrorLogger.Printf("Не удалось сохранить данные заказа в БД: %v\n", err)
			continue
		}

		s.RedisClient.Del(ctx, key)
	}
}
