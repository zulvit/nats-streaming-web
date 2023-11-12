package main

func main() {
	//fixme сделать тут работу с web-сервисом
	//// Загрузка конфигурации
	//cfg := config.NewConfig()
	//
	//// Создание клиента NATS
	//natsClient, err := client.NewNatsClient(cfg)
	//if err != nil {
	//	logger.ErrorLogger.Fatalf("Ошибка подключения к NATS: %v\n", err)
	//}
	//defer natsClient.Close()
	//
	//// Создание публикатора
	//pub := service.NewPublisher(natsClient.Conn)
	//
	//orderData := &model.OrderData{
	//	OrderUID: "8f380f40",
	//}
	//
	//// Публикация данных
	//err = pub.PublishOrder(context.Background(), orderData, cfg.Subject)
	//if err != nil {
	//	logger.ErrorLogger.Printf("Ошибка при публикации данных: %v\n", err)
	//}
	//
	//// Ожидание прерывания программы
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//logger.InfoLogger.Println("Завершение работы приложения")
}
