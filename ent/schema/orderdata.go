package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"nats-streaming-web/pkg/model"
)

type OrderData struct {
	ent.Schema
}

func (OrderData) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_uid").NotEmpty(),
		field.String("track_number"),
		field.String("entry"),
		field.JSON("delivery", &model.Delivery{}),
		field.JSON("payment", &model.Payment{}),
		field.JSON("items", []*model.Item{}),
		field.String("locale"),
		field.String("internal_signature"),
		field.String("customer_id"),
		field.String("delivery_service"),
		field.String("shardkey"),
		field.Int("sm_id"),
		field.String("date_created"),
		field.String("oof_shard"),
	}
}

func (OrderData) Edges() []ent.Edge {
	return nil
}
