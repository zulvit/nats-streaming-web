package publisher

import (
	"context"
	"github.com/nats-io/stan.go"
	"nats-streaming-web/internal/logger"
)

type Serializable interface {
	Serializable() ([]byte, error)
}

type Publisher struct {
	NatsConn stan.Conn
}

func NewPublisher(natsConn stan.Conn) *Publisher {
	return &Publisher{
		NatsConn: natsConn,
	}
}

func (p Publisher) Publish(ctx context.Context, data Serializable, subject string) error {
	serialize, err := data.Serializable()
	if err != nil {
		logger.ErrorLogger.Println("Ошибка при сериализации данных заказа:", err)
		return err
	}

	err = p.NatsConn.Publish(subject, serialize)
	if err != nil {
		logger.ErrorLogger.Println("Ошибка при публикации в NATS:", err)
		return err
	}

	logger.InfoLogger.Println("Данные отправлены в NATS")
	return nil
}
