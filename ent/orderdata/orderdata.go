// Code generated by ent, DO NOT EDIT.

package orderdata

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the orderdata type in the database.
	Label = "order_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrderUID holds the string denoting the order_uid field in the database.
	FieldOrderUID = "order_uid"
	// FieldTrackNumber holds the string denoting the track_number field in the database.
	FieldTrackNumber = "track_number"
	// FieldEntry holds the string denoting the entry field in the database.
	FieldEntry = "entry"
	// FieldDelivery holds the string denoting the delivery field in the database.
	FieldDelivery = "delivery"
	// FieldPayment holds the string denoting the payment field in the database.
	FieldPayment = "payment"
	// FieldItems holds the string denoting the items field in the database.
	FieldItems = "items"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldInternalSignature holds the string denoting the internal_signature field in the database.
	FieldInternalSignature = "internal_signature"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldDeliveryService holds the string denoting the delivery_service field in the database.
	FieldDeliveryService = "delivery_service"
	// FieldShardkey holds the string denoting the shardkey field in the database.
	FieldShardkey = "shardkey"
	// FieldSmID holds the string denoting the sm_id field in the database.
	FieldSmID = "sm_id"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldOofShard holds the string denoting the oof_shard field in the database.
	FieldOofShard = "oof_shard"
	// Table holds the table name of the orderdata in the database.
	Table = "order_data"
)

// Columns holds all SQL columns for orderdata fields.
var Columns = []string{
	FieldID,
	FieldOrderUID,
	FieldTrackNumber,
	FieldEntry,
	FieldDelivery,
	FieldPayment,
	FieldItems,
	FieldLocale,
	FieldInternalSignature,
	FieldCustomerID,
	FieldDeliveryService,
	FieldShardkey,
	FieldSmID,
	FieldDateCreated,
	FieldOofShard,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// OrderUIDValidator is a validator for the "order_uid" field. It is called by the builders before save.
	OrderUIDValidator func(string) error
)

// OrderOption defines the ordering options for the OrderData queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrderUID orders the results by the order_uid field.
func ByOrderUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderUID, opts...).ToFunc()
}

// ByTrackNumber orders the results by the track_number field.
func ByTrackNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrackNumber, opts...).ToFunc()
}

// ByEntry orders the results by the entry field.
func ByEntry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntry, opts...).ToFunc()
}

// ByLocale orders the results by the locale field.
func ByLocale(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocale, opts...).ToFunc()
}

// ByInternalSignature orders the results by the internal_signature field.
func ByInternalSignature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInternalSignature, opts...).ToFunc()
}

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// ByDeliveryService orders the results by the delivery_service field.
func ByDeliveryService(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeliveryService, opts...).ToFunc()
}

// ByShardkey orders the results by the shardkey field.
func ByShardkey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShardkey, opts...).ToFunc()
}

// BySmID orders the results by the sm_id field.
func BySmID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSmID, opts...).ToFunc()
}

// ByDateCreated orders the results by the date_created field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByOofShard orders the results by the oof_shard field.
func ByOofShard(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOofShard, opts...).ToFunc()
}
