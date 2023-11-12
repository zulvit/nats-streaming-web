package service

import (
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"nats-streaming-web/dbclient"
	"nats-streaming-web/internal/logger"
	"nats-streaming-web/pkg/cache"
	"nats-streaming-web/pkg/client"
	"nats-streaming-web/pkg/model"
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

func (s ClientOrderService) SubscribeAndProcess(ctx context.Context, subject string) {
	s.NatsClient.Conn.Subscribe(subject, func(msg *stan.Msg) {
		var orderData model.OrderData
		err := json.Unmarshal(msg.Data, &orderData)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка десериализации данных: %v\n", err)
			return
		}

		// Сохранение данных заказа в базу данных
		_, err = s.DBClient.SaveOrderData(ctx, &orderData)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка при сохранении данных в базе данных: %v\n", err)
			// Вы можете здесь решить, стоит ли продолжать кэширование, если не удалось сохранить в БД
			return
		}

		// Кэширование данных в Redis
		err = s.RedisClient.Set(ctx, "order:"+orderData.OrderUID, msg.Data)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка при записи в Redis: %v\n", err)
		}
	}, stan.DurableName("my-durable"))
}
