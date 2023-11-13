package rest

import (
	"context"
	"encoding/json"
	"errors"
	"nats-streaming-web/ent"
	"nats-streaming-web/pkg/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDBClient struct {
	OrderData *ent.OrderData
	// Флаг для имитации ошибки при вызове GetOrderData
	ShouldReturnError bool
}

// GetOrderData - мок-метод, соответствующий методу в DBClient.
func (m *MockDBClient) GetOrderData(ctx context.Context, id int) (*ent.OrderData, error) {
	if m.ShouldReturnError {
		return nil, errors.New("error fetching order data")
	}
	return m.OrderData, nil
}

func TestNewOrderHandler_Success(t *testing.T) {
	mockData := &model.OrderData{
		OrderUID:          "test123",
		TrackNumber:       "TK123456789",
		Entry:             "EntryExample",
		Locale:            "en",
		InternalSignature: "InternalSignatureExample",
		CustomerID:        "Customer123",
		DeliveryService:   "DeliveryServiceExample",
		Shardkey:          "ShardkeyExample",
		SmID:              123,
		DateCreated:       "2021-01-01T00:00:00Z",
		OofShard:          "OofShardExample",
	}

	mockDB := &MockDBClient{
		OrderData: &ent.OrderData{
			OrderUID:          mockData.OrderUID,
			TrackNumber:       mockData.TrackNumber,
			Entry:             mockData.Entry,
			Locale:            mockData.Locale,
			InternalSignature: mockData.InternalSignature,
			CustomerID:        mockData.CustomerID,
			DeliveryService:   mockData.DeliveryService,
			Shardkey:          mockData.Shardkey,
			SmID:              mockData.SmID,
			DateCreated:       mockData.DateCreated,
			OofShard:          mockData.OofShard,
		},
	}

	req, err := http.NewRequest("GET", "/?id=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewOrderHandler(mockDB)

	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Проверка тела ответа
	var order model.OrderData
	err = json.NewDecoder(rr.Body).Decode(&order)
	if err != nil {
		t.Fatal(err)
	}

	// Проверка, что данные в order соответствуют mockData
	if order.OrderUID != mockData.OrderUID ||
		order.TrackNumber != mockData.TrackNumber ||
		order.Entry != mockData.Entry ||
		order.Locale != mockData.Locale ||
		order.InternalSignature != mockData.InternalSignature ||
		order.CustomerID != mockData.CustomerID ||
		order.DeliveryService != mockData.DeliveryService ||
		order.Shardkey != mockData.Shardkey ||
		order.SmID != mockData.SmID ||
		order.DateCreated != mockData.DateCreated ||
		order.OofShard != mockData.OofShard {
		t.Errorf("handler returned incorrect order data")
	}

	if order.Delivery.Name != mockData.Delivery.Name ||
		order.Delivery.Phone != mockData.Delivery.Phone ||
		order.Delivery.Zip != mockData.Delivery.Zip ||
		order.Delivery.City != mockData.Delivery.City ||
		order.Delivery.Address != mockData.Delivery.Address ||
		order.Delivery.Region != mockData.Delivery.Region ||
		order.Delivery.Email != mockData.Delivery.Email {
		t.Errorf("handler returned incorrect delivery data")
	}

	if order.Payment.Transaction != mockData.Payment.Transaction ||
		order.Payment.RequestID != mockData.Payment.RequestID ||
		order.Payment.Currency != mockData.Payment.Currency ||
		order.Payment.Provider != mockData.Payment.Provider ||
		order.Payment.Amount != mockData.Payment.Amount ||
		order.Payment.PaymentDt != mockData.Payment.PaymentDt ||
		order.Payment.Bank != mockData.Payment.Bank ||
		order.Payment.DeliveryCost != mockData.Payment.DeliveryCost ||
		order.Payment.GoodsTotal != mockData.Payment.GoodsTotal ||
		order.Payment.CustomFee != mockData.Payment.CustomFee {
		t.Errorf("handler returned incorrect payment data")
	}

	// Проверка элементов заказа, если они есть
	for i, item := range order.Items {
		mockItem := mockData.Items[i]
		if item.ChrtID != mockItem.ChrtID ||
			item.TrackNumber != mockItem.TrackNumber ||
			item.Price != mockItem.Price ||
			item.Rid != mockItem.Rid ||
			item.Name != mockItem.Name ||
			item.Sale != mockItem.Sale ||
			item.Size != mockItem.Size ||
			item.TotalPrice != mockItem.TotalPrice ||
			item.NmID != mockItem.NmID ||
			item.Brand != mockItem.Brand ||
			item.Status != mockItem.Status {
			t.Errorf("handler returned incorrect item data at index %d", i)
		}
	}
}
func TestNewOrderHandler_NoID(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewOrderHandler(&MockDBClient{})

	handler.ServeHTTP(rr, req)

	// Проверка статуса ответа
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
func TestNewOrderHandler_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewOrderHandler(&MockDBClient{})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
func TestNewOrderHandler_DBError(t *testing.T) {
	mockDB := &MockDBClient{ShouldReturnError: true}

	req, err := http.NewRequest("GET", "/?id=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewOrderHandler(mockDB)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
