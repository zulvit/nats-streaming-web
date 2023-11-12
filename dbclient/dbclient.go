package dbclient

import (
	"context"
	_ "github.com/lib/pq"
	"nats-streaming-web/ent"
	"nats-streaming-web/internal/logger"
	"nats-streaming-web/pkg/model"
)

type DBClient struct {
	Client *ent.Client
}

func NewDBClient(dataSourceName string) (*DBClient, error) {
	client, err := ent.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		logger.ErrorLogger.Println("Ошибка применения миграции")
	}
	return &DBClient{Client: client}, nil
}

func (db *DBClient) SaveOrderData(ctx context.Context, orderData *model.OrderData) (*ent.OrderData, error) {
	// Преобразование []Item в []*model.Item
	var items []*model.Item
	for _, item := range orderData.Items {
		items = append(items, &model.Item{
			ChrtID:      item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			Rid:         item.Rid,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmID:        item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
		})
	}

	// Создание и сохранение объекта OrderData
	return db.Client.OrderData.Create().
		SetOrderUID(orderData.OrderUID).
		SetTrackNumber(orderData.TrackNumber).
		SetEntry(orderData.Entry).
		SetDelivery(&orderData.Delivery).
		SetPayment(&orderData.Payment).
		SetItems(items).
		SetLocale(orderData.Locale).
		SetInternalSignature(orderData.InternalSignature).
		SetCustomerID(orderData.CustomerID).
		SetDeliveryService(orderData.DeliveryService).
		SetShardkey(orderData.Shardkey).
		SetSmID(orderData.SmID).
		SetDateCreated(orderData.DateCreated).
		SetOofShard(orderData.OofShard).
		Save(ctx)
}