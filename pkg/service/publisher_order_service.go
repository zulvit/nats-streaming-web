package service

import (
	"context"
	"nats-streaming-web/pkg/publisher"
)

type DataProcessor struct {
	Pub *publisher.Publisher
}

func NewDataProcessor(pub *publisher.Publisher) *DataProcessor {
	return &DataProcessor{Pub: pub}
}

func (dp *DataProcessor) ProcessAndPublish(ctx context.Context, data publisher.Serializable, subject string) error {
	//todo Здесь может быть дополнительная бизнес-логика перед публикацией
	return dp.Pub.Publish(ctx, data, subject)
}
