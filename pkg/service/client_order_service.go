package service

import (
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"nats-streaming-web/internal/logger"
	"nats-streaming-web/pkg/cache"
	"nats-streaming-web/pkg/client"
	"nats-streaming-web/pkg/model"
)

type ClientOrderService struct {
	NatsClient  *client.NatsClient
	RedisClient *cache.RedisClient
}

func NewClientOrderService(natsClient *client.NatsClient, redisClient *cache.RedisClient) *ClientOrderService {
	return &ClientOrderService{
		RedisClient: redisClient,
		NatsClient:  natsClient}
}

func (cos *ClientOrderService) SubscribeAndProcess(ctx context.Context, subject string) {
	cos.NatsClient.Conn.Subscribe(subject, func(msg *stan.Msg) {
		var orderData model.OrderData
		err := json.Unmarshal(msg.Data, &orderData)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка десериализации данных: %v\n", err)
			return
		}

		// todo реализовать тут сохранение в базу данных

		err = cos.RedisClient.Set(ctx, "order:"+orderData.OrderUID, msg.Data)
		if err != nil {
			logger.ErrorLogger.Printf("Ошибка при записи в Redis: %v\n", err)
		}
	}, stan.DurableName("my-durable"))
}
