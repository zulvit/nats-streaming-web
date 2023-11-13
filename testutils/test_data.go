package testutils

import (
	"context"
	service2 "nats-streaming-web/internal/service"
	"nats-streaming-web/pkg/model"
)

// CreateAndPublishTestOrder создает тестовые данные и публикует их.
func CreateAndPublishTestOrder(ctx context.Context, dataProcessor *service2.DataProcessor, subject string) error {
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

	return dataProcessor.ProcessAndPublish(ctx, &testOrder, subject)
}
