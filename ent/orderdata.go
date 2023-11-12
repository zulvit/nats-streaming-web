// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"nats-streaming-web/ent/orderdata"
	"nats-streaming-web/pkg/model"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OrderData is the model entity for the OrderData schema.
type OrderData struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OrderUID holds the value of the "order_uid" field.
	OrderUID string `json:"order_uid,omitempty"`
	// TrackNumber holds the value of the "track_number" field.
	TrackNumber string `json:"track_number,omitempty"`
	// Entry holds the value of the "entry" field.
	Entry string `json:"entry,omitempty"`
	// Delivery holds the value of the "delivery" field.
	Delivery *model.Delivery `json:"delivery,omitempty"`
	// Payment holds the value of the "payment" field.
	Payment *model.Payment `json:"payment,omitempty"`
	// Items holds the value of the "items" field.
	Items []*model.Item `json:"items,omitempty"`
	// Locale holds the value of the "locale" field.
	Locale string `json:"locale,omitempty"`
	// InternalSignature holds the value of the "internal_signature" field.
	InternalSignature string `json:"internal_signature,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID string `json:"customer_id,omitempty"`
	// DeliveryService holds the value of the "delivery_service" field.
	DeliveryService string `json:"delivery_service,omitempty"`
	// Shardkey holds the value of the "shardkey" field.
	Shardkey string `json:"shardkey,omitempty"`
	// SmID holds the value of the "sm_id" field.
	SmID int `json:"sm_id,omitempty"`
	// DateCreated holds the value of the "date_created" field.
	DateCreated string `json:"date_created,omitempty"`
	// OofShard holds the value of the "oof_shard" field.
	OofShard     string `json:"oof_shard,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderdata.FieldDelivery, orderdata.FieldPayment, orderdata.FieldItems:
			values[i] = new([]byte)
		case orderdata.FieldID, orderdata.FieldSmID:
			values[i] = new(sql.NullInt64)
		case orderdata.FieldOrderUID, orderdata.FieldTrackNumber, orderdata.FieldEntry, orderdata.FieldLocale, orderdata.FieldInternalSignature, orderdata.FieldCustomerID, orderdata.FieldDeliveryService, orderdata.FieldShardkey, orderdata.FieldDateCreated, orderdata.FieldOofShard:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderData fields.
func (od *OrderData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderdata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			od.ID = int(value.Int64)
		case orderdata.FieldOrderUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_uid", values[i])
			} else if value.Valid {
				od.OrderUID = value.String
			}
		case orderdata.FieldTrackNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field track_number", values[i])
			} else if value.Valid {
				od.TrackNumber = value.String
			}
		case orderdata.FieldEntry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entry", values[i])
			} else if value.Valid {
				od.Entry = value.String
			}
		case orderdata.FieldDelivery:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field delivery", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &od.Delivery); err != nil {
					return fmt.Errorf("unmarshal field delivery: %w", err)
				}
			}
		case orderdata.FieldPayment:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field payment", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &od.Payment); err != nil {
					return fmt.Errorf("unmarshal field payment: %w", err)
				}
			}
		case orderdata.FieldItems:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field items", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &od.Items); err != nil {
					return fmt.Errorf("unmarshal field items: %w", err)
				}
			}
		case orderdata.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				od.Locale = value.String
			}
		case orderdata.FieldInternalSignature:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field internal_signature", values[i])
			} else if value.Valid {
				od.InternalSignature = value.String
			}
		case orderdata.FieldCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				od.CustomerID = value.String
			}
		case orderdata.FieldDeliveryService:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field delivery_service", values[i])
			} else if value.Valid {
				od.DeliveryService = value.String
			}
		case orderdata.FieldShardkey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shardkey", values[i])
			} else if value.Valid {
				od.Shardkey = value.String
			}
		case orderdata.FieldSmID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sm_id", values[i])
			} else if value.Valid {
				od.SmID = int(value.Int64)
			}
		case orderdata.FieldDateCreated:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field date_created", values[i])
			} else if value.Valid {
				od.DateCreated = value.String
			}
		case orderdata.FieldOofShard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field oof_shard", values[i])
			} else if value.Valid {
				od.OofShard = value.String
			}
		default:
			od.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderData.
// This includes values selected through modifiers, order, etc.
func (od *OrderData) Value(name string) (ent.Value, error) {
	return od.selectValues.Get(name)
}

// Update returns a builder for updating this OrderData.
// Note that you need to call OrderData.Unwrap() before calling this method if this OrderData
// was returned from a transaction, and the transaction was committed or rolled back.
func (od *OrderData) Update() *OrderDataUpdateOne {
	return NewOrderDataClient(od.config).UpdateOne(od)
}

// Unwrap unwraps the OrderData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (od *OrderData) Unwrap() *OrderData {
	_tx, ok := od.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderData is not a transactional entity")
	}
	od.config.driver = _tx.drv
	return od
}

// String implements the fmt.Stringer.
func (od *OrderData) String() string {
	var builder strings.Builder
	builder.WriteString("OrderData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", od.ID))
	builder.WriteString("order_uid=")
	builder.WriteString(od.OrderUID)
	builder.WriteString(", ")
	builder.WriteString("track_number=")
	builder.WriteString(od.TrackNumber)
	builder.WriteString(", ")
	builder.WriteString("entry=")
	builder.WriteString(od.Entry)
	builder.WriteString(", ")
	builder.WriteString("delivery=")
	builder.WriteString(fmt.Sprintf("%v", od.Delivery))
	builder.WriteString(", ")
	builder.WriteString("payment=")
	builder.WriteString(fmt.Sprintf("%v", od.Payment))
	builder.WriteString(", ")
	builder.WriteString("items=")
	builder.WriteString(fmt.Sprintf("%v", od.Items))
	builder.WriteString(", ")
	builder.WriteString("locale=")
	builder.WriteString(od.Locale)
	builder.WriteString(", ")
	builder.WriteString("internal_signature=")
	builder.WriteString(od.InternalSignature)
	builder.WriteString(", ")
	builder.WriteString("customer_id=")
	builder.WriteString(od.CustomerID)
	builder.WriteString(", ")
	builder.WriteString("delivery_service=")
	builder.WriteString(od.DeliveryService)
	builder.WriteString(", ")
	builder.WriteString("shardkey=")
	builder.WriteString(od.Shardkey)
	builder.WriteString(", ")
	builder.WriteString("sm_id=")
	builder.WriteString(fmt.Sprintf("%v", od.SmID))
	builder.WriteString(", ")
	builder.WriteString("date_created=")
	builder.WriteString(od.DateCreated)
	builder.WriteString(", ")
	builder.WriteString("oof_shard=")
	builder.WriteString(od.OofShard)
	builder.WriteByte(')')
	return builder.String()
}

// OrderDataSlice is a parsable slice of OrderData.
type OrderDataSlice []*OrderData
