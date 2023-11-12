package client

import (
	"github.com/nats-io/stan.go"
	"nats-streaming-web/internal/logger"
	"nats-streaming-web/pkg/config"
)

// NatsClient структура для управления соединением с NATS.
type NatsClient struct {
	Conn stan.Conn
}

// NewNatsClient создает новый клиент NATS.
func NewNatsClient(cfg *config.Config) (*NatsClient, error) {
	natsConn, err := stan.Connect(
		"test-cluster",
		"publisher",
		stan.NatsURL(cfg.NatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logger.ErrorLogger.Printf("Соединение с NATS потеряно: %v\n", reason)
		}),
		stan.Pings(10, 5),
	)
	if err != nil {
		return nil, err
	}

	return &NatsClient{Conn: natsConn}, nil
}

// Close закрывает соединение с NATS.
func (nc *NatsClient) Close() {
	if nc.Conn != nil {
		if err := nc.Conn.Close(); err != nil {
			logger.ErrorLogger.Println("Ошибка при закрытии соединения с NATS:", err)
		}
	}
}
