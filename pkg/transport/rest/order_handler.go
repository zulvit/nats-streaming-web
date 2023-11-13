package rest

import (
	"encoding/json"
	"nats-streaming-web/dbclient"
	"net/http"
	"strconv"
)

func NewOrderHandler(db *dbclient.DBClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		orderID := request.URL.Query().Get("id")
		ctx := request.Context()
		if orderID == "" {
			http.Error(writer, "ID не найден", http.StatusBadRequest)
			return
		}

		idInt, err := strconv.Atoi(orderID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		order, err := db.GetOrderData(ctx, idInt)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(order)
	}
}
