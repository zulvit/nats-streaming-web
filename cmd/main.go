package main

import (
	"context"
	"nats-streaming-web/dbclient"
	"nats-streaming-web/internal/config"
	"nats-streaming-web/internal/logger"
	"nats-streaming-web/pkg/cache"
	"nats-streaming-web/pkg/client"
	"nats-streaming-web/pkg/model"
	"nats-streaming-web/pkg/publisher"
	"nats-streaming-web/service"
	"time"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	// Инициализация клиентов
	natsClient, err := client.NewNatsClient(cfg)
	if err != nil {
		logger.ErrorLogger.Println("Ошибка при создании клиента NATS:", err)
		return
	}
	defer natsClient.Close()

	redisClient := cache.NewRedisClient(cfg.RedisURL)
	dbClient, err := dbclient.NewDBClient(cfg.DBConnectionString)
	if err != nil {
		logger.ErrorLogger.Println("Ошибка при создании клиента БД:", err)
		return
	}

	// Инициализация сервисов
	dataProcessor := service.NewDataProcessor(publisher.NewPublisher(natsClient.Conn))
	clientOrderService := service.NewClientOrderService(natsClient, redisClient, dbClient)

	// Подписка на обработку сообщений
	clientOrderService.SubscribeAndProcess(ctx, cfg.Subject)

	// Создание тестового объекта OrderData
	testOrder := model.OrderData{
		OrderUID:    "test123",
		TrackNumber: "TK123456789",
		Entry:       "EntryExample",
		Delivery: model.Delivery{
			Name:    "John Doe",
			Phone:   "+1234567890",
			Zip:     "123456",
			City:    "Sample City",
			Address: "123 Sample Street",
			Region:  "Sample Region",
			Email:   "johndoe@example.com",
		},
		Payment: model.Payment{
			Transaction:  "TX123456789",
			RequestID:    "RQ123456789",
			Currency:     "USD",
			Provider:     "ProviderExample",
			Amount:       100,
			PaymentDt:    1609459200,
			Bank:         "Sample Bank",
			DeliveryCost: 10,
			GoodsTotal:   90,
			CustomFee:    5,
		},
		Items: []model.Item{
			{
				ChrtID:      123456,
				TrackNumber: "TK123456789",
				Price:       50,
				Rid:         "RID123456789",
				Name:        "Sample Item 1",
				Sale:        5,
				Size:        "M",
				TotalPrice:  45,
				NmID:        654321,
				Brand:       "BrandExample",
				Status:      1,
			},
		},
		Locale:            "en",
		InternalSignature: "InternalSignatureExample",
		CustomerID:        "Customer123",
		DeliveryService:   "DeliveryServiceExample",
		Shardkey:          "ShardkeyExample",
		SmID:              123,
		DateCreated:       "2021-01-01T00:00:00Z",
		OofShard:          "OofShardExample",
	}

	// Публикация тестового объекта
	if err := dataProcessor.ProcessAndPublish(ctx, &testOrder, cfg.Subject); err != nil {
		logger.ErrorLogger.Println("Ошибка при публикации заказа:", err)
	}

	// Даем время для обработки сообщения
	time.Sleep(5 * time.Second)
}
