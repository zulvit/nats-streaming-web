package publisher

import (
	"context"
	"github.com/nats-io/stan.go"
	"nats-streaming-web/internal/logger"
	"nats-streaming-web/pkg/model"
)

type Publisher struct {
	NatsConn stan.Conn
}

func NewPublisher(natsConn stan.Conn) *Publisher {
	return &Publisher{
		NatsConn: natsConn,
	}
}

func (p Publisher) PublishOrder(ctx context.Context, data *model.OrderData, subject string) error {
	serialize, err := data.Serialize()
	if err != nil {
		logger.ErrorLogger.Println("Ошибка при сериализации данных заказа:", err)
		return err
	}

	err = p.NatsConn.Publish(subject, serialize)
	if err != nil {
		logger.ErrorLogger.Println("Ошибка при публикации в NATS:", err)
		return err
	}

	logger.InfoLogger.Println("Заказ успешно отправлен в NATS:", data.OrderUID)
	return nil
}
