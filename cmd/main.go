package main

import (
	"context"
	"nats-streaming-web/internal/logger"
	"nats-streaming-web/pkg/config"
	"nats-streaming-web/pkg/model"
	"nats-streaming-web/pkg/natsclient"
	"nats-streaming-web/pkg/publisher"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Загрузка конфигурации
	cfg := config.NewConfig()

	// Создание клиента NATS
	natsClient, err := natsclient.NewNatsClient(cfg)
	if err != nil {
		logger.ErrorLogger.Fatalf("Ошибка подключения к NATS: %v\n", err)
	}
	defer natsClient.Close()

	// Создание публикатора
	pub := publisher.NewPublisher(natsClient.Conn)

	orderData := &model.OrderData{
		OrderUID: "8f380f40",
	}

	// Публикация данных
	err = pub.PublishOrder(context.Background(), orderData, cfg.Subject)
	if err != nil {
		logger.ErrorLogger.Printf("Ошибка при публикации данных: %v\n", err)
	}

	// Ожидание прерывания программы (например, сигнала SIGINT)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.InfoLogger.Println("Завершение работы приложения")
}
