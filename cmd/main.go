package main

import (
	"context"
	"nats-streaming-web/dbclient"
	"nats-streaming-web/internal/config"
	"nats-streaming-web/internal/logger"
	service2 "nats-streaming-web/internal/service"
	"nats-streaming-web/pkg/cache"
	"nats-streaming-web/pkg/client"
	"nats-streaming-web/pkg/publisher"
	"nats-streaming-web/pkg/transport/rest"
	"nats-streaming-web/testutils"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
	dataProcessor := service2.NewDataProcessor(publisher.NewPublisher(natsClient.Conn))
	clientOrderService := service2.NewClientOrderService(natsClient, redisClient, dbClient)
	clientOrderService.SubscribeAndProcess(ctx, cfg.Subject)

	// Запуск мониторинга в отдельной горутине
	go clientOrderService.MonitorDBAndRestoreFromCache(ctx)

	// Создание и публикация тестовых данных
	if err := testutils.CreateAndPublishTestOrder(ctx, dataProcessor, cfg.Subject); err != nil {
		logger.ErrorLogger.Println("Ошибка при публикации тестового заказа:", err)
	}
	time.Sleep(5 * time.Second)

	// Настройка и запуск HTTP-сервера
	orderHandler := rest.NewOrderHandler(dbClient)
	http.HandleFunc("/order", orderHandler)
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.ErrorLogger.Println("Ошибка http", err)
	}
}
