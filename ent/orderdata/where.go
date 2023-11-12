// Code generated by ent, DO NOT EDIT.

package orderdata

import (
	"nats-streaming-web/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldID, id))
}

// OrderUID applies equality check predicate on the "order_uid" field. It's identical to OrderUIDEQ.
func OrderUID(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldOrderUID, v))
}

// TrackNumber applies equality check predicate on the "track_number" field. It's identical to TrackNumberEQ.
func TrackNumber(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldTrackNumber, v))
}

// Entry applies equality check predicate on the "entry" field. It's identical to EntryEQ.
func Entry(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldEntry, v))
}

// Locale applies equality check predicate on the "locale" field. It's identical to LocaleEQ.
func Locale(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldLocale, v))
}

// InternalSignature applies equality check predicate on the "internal_signature" field. It's identical to InternalSignatureEQ.
func InternalSignature(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldInternalSignature, v))
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldCustomerID, v))
}

// DeliveryService applies equality check predicate on the "delivery_service" field. It's identical to DeliveryServiceEQ.
func DeliveryService(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldDeliveryService, v))
}

// Shardkey applies equality check predicate on the "shardkey" field. It's identical to ShardkeyEQ.
func Shardkey(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldShardkey, v))
}

// SmID applies equality check predicate on the "sm_id" field. It's identical to SmIDEQ.
func SmID(v int) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldSmID, v))
}

// DateCreated applies equality check predicate on the "date_created" field. It's identical to DateCreatedEQ.
func DateCreated(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldDateCreated, v))
}

// OofShard applies equality check predicate on the "oof_shard" field. It's identical to OofShardEQ.
func OofShard(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldOofShard, v))
}

// OrderUIDEQ applies the EQ predicate on the "order_uid" field.
func OrderUIDEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldOrderUID, v))
}

// OrderUIDNEQ applies the NEQ predicate on the "order_uid" field.
func OrderUIDNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldOrderUID, v))
}

// OrderUIDIn applies the In predicate on the "order_uid" field.
func OrderUIDIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldOrderUID, vs...))
}

// OrderUIDNotIn applies the NotIn predicate on the "order_uid" field.
func OrderUIDNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldOrderUID, vs...))
}

// OrderUIDGT applies the GT predicate on the "order_uid" field.
func OrderUIDGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldOrderUID, v))
}

// OrderUIDGTE applies the GTE predicate on the "order_uid" field.
func OrderUIDGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldOrderUID, v))
}

// OrderUIDLT applies the LT predicate on the "order_uid" field.
func OrderUIDLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldOrderUID, v))
}

// OrderUIDLTE applies the LTE predicate on the "order_uid" field.
func OrderUIDLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldOrderUID, v))
}

// OrderUIDContains applies the Contains predicate on the "order_uid" field.
func OrderUIDContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldOrderUID, v))
}

// OrderUIDHasPrefix applies the HasPrefix predicate on the "order_uid" field.
func OrderUIDHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldOrderUID, v))
}

// OrderUIDHasSuffix applies the HasSuffix predicate on the "order_uid" field.
func OrderUIDHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldOrderUID, v))
}

// OrderUIDEqualFold applies the EqualFold predicate on the "order_uid" field.
func OrderUIDEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldOrderUID, v))
}

// OrderUIDContainsFold applies the ContainsFold predicate on the "order_uid" field.
func OrderUIDContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldOrderUID, v))
}

// TrackNumberEQ applies the EQ predicate on the "track_number" field.
func TrackNumberEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldTrackNumber, v))
}

// TrackNumberNEQ applies the NEQ predicate on the "track_number" field.
func TrackNumberNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldTrackNumber, v))
}

// TrackNumberIn applies the In predicate on the "track_number" field.
func TrackNumberIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldTrackNumber, vs...))
}

// TrackNumberNotIn applies the NotIn predicate on the "track_number" field.
func TrackNumberNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldTrackNumber, vs...))
}

// TrackNumberGT applies the GT predicate on the "track_number" field.
func TrackNumberGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldTrackNumber, v))
}

// TrackNumberGTE applies the GTE predicate on the "track_number" field.
func TrackNumberGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldTrackNumber, v))
}

// TrackNumberLT applies the LT predicate on the "track_number" field.
func TrackNumberLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldTrackNumber, v))
}

// TrackNumberLTE applies the LTE predicate on the "track_number" field.
func TrackNumberLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldTrackNumber, v))
}

// TrackNumberContains applies the Contains predicate on the "track_number" field.
func TrackNumberContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldTrackNumber, v))
}

// TrackNumberHasPrefix applies the HasPrefix predicate on the "track_number" field.
func TrackNumberHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldTrackNumber, v))
}

// TrackNumberHasSuffix applies the HasSuffix predicate on the "track_number" field.
func TrackNumberHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldTrackNumber, v))
}

// TrackNumberEqualFold applies the EqualFold predicate on the "track_number" field.
func TrackNumberEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldTrackNumber, v))
}

// TrackNumberContainsFold applies the ContainsFold predicate on the "track_number" field.
func TrackNumberContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldTrackNumber, v))
}

// EntryEQ applies the EQ predicate on the "entry" field.
func EntryEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldEntry, v))
}

// EntryNEQ applies the NEQ predicate on the "entry" field.
func EntryNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldEntry, v))
}

// EntryIn applies the In predicate on the "entry" field.
func EntryIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldEntry, vs...))
}

// EntryNotIn applies the NotIn predicate on the "entry" field.
func EntryNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldEntry, vs...))
}

// EntryGT applies the GT predicate on the "entry" field.
func EntryGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldEntry, v))
}

// EntryGTE applies the GTE predicate on the "entry" field.
func EntryGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldEntry, v))
}

// EntryLT applies the LT predicate on the "entry" field.
func EntryLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldEntry, v))
}

// EntryLTE applies the LTE predicate on the "entry" field.
func EntryLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldEntry, v))
}

// EntryContains applies the Contains predicate on the "entry" field.
func EntryContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldEntry, v))
}

// EntryHasPrefix applies the HasPrefix predicate on the "entry" field.
func EntryHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldEntry, v))
}

// EntryHasSuffix applies the HasSuffix predicate on the "entry" field.
func EntryHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldEntry, v))
}

// EntryEqualFold applies the EqualFold predicate on the "entry" field.
func EntryEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldEntry, v))
}

// EntryContainsFold applies the ContainsFold predicate on the "entry" field.
func EntryContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldEntry, v))
}

// LocaleEQ applies the EQ predicate on the "locale" field.
func LocaleEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldLocale, v))
}

// LocaleNEQ applies the NEQ predicate on the "locale" field.
func LocaleNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldLocale, v))
}

// LocaleIn applies the In predicate on the "locale" field.
func LocaleIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldLocale, vs...))
}

// LocaleNotIn applies the NotIn predicate on the "locale" field.
func LocaleNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldLocale, vs...))
}

// LocaleGT applies the GT predicate on the "locale" field.
func LocaleGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldLocale, v))
}

// LocaleGTE applies the GTE predicate on the "locale" field.
func LocaleGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldLocale, v))
}

// LocaleLT applies the LT predicate on the "locale" field.
func LocaleLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldLocale, v))
}

// LocaleLTE applies the LTE predicate on the "locale" field.
func LocaleLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldLocale, v))
}

// LocaleContains applies the Contains predicate on the "locale" field.
func LocaleContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldLocale, v))
}

// LocaleHasPrefix applies the HasPrefix predicate on the "locale" field.
func LocaleHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldLocale, v))
}

// LocaleHasSuffix applies the HasSuffix predicate on the "locale" field.
func LocaleHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldLocale, v))
}

// LocaleEqualFold applies the EqualFold predicate on the "locale" field.
func LocaleEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldLocale, v))
}

// LocaleContainsFold applies the ContainsFold predicate on the "locale" field.
func LocaleContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldLocale, v))
}

// InternalSignatureEQ applies the EQ predicate on the "internal_signature" field.
func InternalSignatureEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldInternalSignature, v))
}

// InternalSignatureNEQ applies the NEQ predicate on the "internal_signature" field.
func InternalSignatureNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldInternalSignature, v))
}

// InternalSignatureIn applies the In predicate on the "internal_signature" field.
func InternalSignatureIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldInternalSignature, vs...))
}

// InternalSignatureNotIn applies the NotIn predicate on the "internal_signature" field.
func InternalSignatureNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldInternalSignature, vs...))
}

// InternalSignatureGT applies the GT predicate on the "internal_signature" field.
func InternalSignatureGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldInternalSignature, v))
}

// InternalSignatureGTE applies the GTE predicate on the "internal_signature" field.
func InternalSignatureGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldInternalSignature, v))
}

// InternalSignatureLT applies the LT predicate on the "internal_signature" field.
func InternalSignatureLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldInternalSignature, v))
}

// InternalSignatureLTE applies the LTE predicate on the "internal_signature" field.
func InternalSignatureLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldInternalSignature, v))
}

// InternalSignatureContains applies the Contains predicate on the "internal_signature" field.
func InternalSignatureContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldInternalSignature, v))
}

// InternalSignatureHasPrefix applies the HasPrefix predicate on the "internal_signature" field.
func InternalSignatureHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldInternalSignature, v))
}

// InternalSignatureHasSuffix applies the HasSuffix predicate on the "internal_signature" field.
func InternalSignatureHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldInternalSignature, v))
}

// InternalSignatureEqualFold applies the EqualFold predicate on the "internal_signature" field.
func InternalSignatureEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldInternalSignature, v))
}

// InternalSignatureContainsFold applies the ContainsFold predicate on the "internal_signature" field.
func InternalSignatureContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldInternalSignature, v))
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldCustomerID, v))
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldCustomerID, v))
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldCustomerID, vs...))
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldCustomerID, vs...))
}

// CustomerIDGT applies the GT predicate on the "customer_id" field.
func CustomerIDGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldCustomerID, v))
}

// CustomerIDGTE applies the GTE predicate on the "customer_id" field.
func CustomerIDGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldCustomerID, v))
}

// CustomerIDLT applies the LT predicate on the "customer_id" field.
func CustomerIDLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldCustomerID, v))
}

// CustomerIDLTE applies the LTE predicate on the "customer_id" field.
func CustomerIDLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldCustomerID, v))
}

// CustomerIDContains applies the Contains predicate on the "customer_id" field.
func CustomerIDContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldCustomerID, v))
}

// CustomerIDHasPrefix applies the HasPrefix predicate on the "customer_id" field.
func CustomerIDHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldCustomerID, v))
}

// CustomerIDHasSuffix applies the HasSuffix predicate on the "customer_id" field.
func CustomerIDHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldCustomerID, v))
}

// CustomerIDEqualFold applies the EqualFold predicate on the "customer_id" field.
func CustomerIDEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldCustomerID, v))
}

// CustomerIDContainsFold applies the ContainsFold predicate on the "customer_id" field.
func CustomerIDContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldCustomerID, v))
}

// DeliveryServiceEQ applies the EQ predicate on the "delivery_service" field.
func DeliveryServiceEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldDeliveryService, v))
}

// DeliveryServiceNEQ applies the NEQ predicate on the "delivery_service" field.
func DeliveryServiceNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldDeliveryService, v))
}

// DeliveryServiceIn applies the In predicate on the "delivery_service" field.
func DeliveryServiceIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldDeliveryService, vs...))
}

// DeliveryServiceNotIn applies the NotIn predicate on the "delivery_service" field.
func DeliveryServiceNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldDeliveryService, vs...))
}

// DeliveryServiceGT applies the GT predicate on the "delivery_service" field.
func DeliveryServiceGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldDeliveryService, v))
}

// DeliveryServiceGTE applies the GTE predicate on the "delivery_service" field.
func DeliveryServiceGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldDeliveryService, v))
}

// DeliveryServiceLT applies the LT predicate on the "delivery_service" field.
func DeliveryServiceLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldDeliveryService, v))
}

// DeliveryServiceLTE applies the LTE predicate on the "delivery_service" field.
func DeliveryServiceLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldDeliveryService, v))
}

// DeliveryServiceContains applies the Contains predicate on the "delivery_service" field.
func DeliveryServiceContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldDeliveryService, v))
}

// DeliveryServiceHasPrefix applies the HasPrefix predicate on the "delivery_service" field.
func DeliveryServiceHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldDeliveryService, v))
}

// DeliveryServiceHasSuffix applies the HasSuffix predicate on the "delivery_service" field.
func DeliveryServiceHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldDeliveryService, v))
}

// DeliveryServiceEqualFold applies the EqualFold predicate on the "delivery_service" field.
func DeliveryServiceEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldDeliveryService, v))
}

// DeliveryServiceContainsFold applies the ContainsFold predicate on the "delivery_service" field.
func DeliveryServiceContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldDeliveryService, v))
}

// ShardkeyEQ applies the EQ predicate on the "shardkey" field.
func ShardkeyEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldShardkey, v))
}

// ShardkeyNEQ applies the NEQ predicate on the "shardkey" field.
func ShardkeyNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldShardkey, v))
}

// ShardkeyIn applies the In predicate on the "shardkey" field.
func ShardkeyIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldShardkey, vs...))
}

// ShardkeyNotIn applies the NotIn predicate on the "shardkey" field.
func ShardkeyNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldShardkey, vs...))
}

// ShardkeyGT applies the GT predicate on the "shardkey" field.
func ShardkeyGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldShardkey, v))
}

// ShardkeyGTE applies the GTE predicate on the "shardkey" field.
func ShardkeyGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldShardkey, v))
}

// ShardkeyLT applies the LT predicate on the "shardkey" field.
func ShardkeyLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldShardkey, v))
}

// ShardkeyLTE applies the LTE predicate on the "shardkey" field.
func ShardkeyLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldShardkey, v))
}

// ShardkeyContains applies the Contains predicate on the "shardkey" field.
func ShardkeyContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldShardkey, v))
}

// ShardkeyHasPrefix applies the HasPrefix predicate on the "shardkey" field.
func ShardkeyHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldShardkey, v))
}

// ShardkeyHasSuffix applies the HasSuffix predicate on the "shardkey" field.
func ShardkeyHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldShardkey, v))
}

// ShardkeyEqualFold applies the EqualFold predicate on the "shardkey" field.
func ShardkeyEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldShardkey, v))
}

// ShardkeyContainsFold applies the ContainsFold predicate on the "shardkey" field.
func ShardkeyContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldShardkey, v))
}

// SmIDEQ applies the EQ predicate on the "sm_id" field.
func SmIDEQ(v int) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldSmID, v))
}

// SmIDNEQ applies the NEQ predicate on the "sm_id" field.
func SmIDNEQ(v int) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldSmID, v))
}

// SmIDIn applies the In predicate on the "sm_id" field.
func SmIDIn(vs ...int) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldSmID, vs...))
}

// SmIDNotIn applies the NotIn predicate on the "sm_id" field.
func SmIDNotIn(vs ...int) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldSmID, vs...))
}

// SmIDGT applies the GT predicate on the "sm_id" field.
func SmIDGT(v int) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldSmID, v))
}

// SmIDGTE applies the GTE predicate on the "sm_id" field.
func SmIDGTE(v int) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldSmID, v))
}

// SmIDLT applies the LT predicate on the "sm_id" field.
func SmIDLT(v int) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldSmID, v))
}

// SmIDLTE applies the LTE predicate on the "sm_id" field.
func SmIDLTE(v int) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldSmID, v))
}

// DateCreatedEQ applies the EQ predicate on the "date_created" field.
func DateCreatedEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "date_created" field.
func DateCreatedNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "date_created" field.
func DateCreatedIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "date_created" field.
func DateCreatedNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "date_created" field.
func DateCreatedGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "date_created" field.
func DateCreatedGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "date_created" field.
func DateCreatedLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "date_created" field.
func DateCreatedLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldDateCreated, v))
}

// DateCreatedContains applies the Contains predicate on the "date_created" field.
func DateCreatedContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldDateCreated, v))
}

// DateCreatedHasPrefix applies the HasPrefix predicate on the "date_created" field.
func DateCreatedHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldDateCreated, v))
}

// DateCreatedHasSuffix applies the HasSuffix predicate on the "date_created" field.
func DateCreatedHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldDateCreated, v))
}

// DateCreatedEqualFold applies the EqualFold predicate on the "date_created" field.
func DateCreatedEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldDateCreated, v))
}

// DateCreatedContainsFold applies the ContainsFold predicate on the "date_created" field.
func DateCreatedContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldDateCreated, v))
}

// OofShardEQ applies the EQ predicate on the "oof_shard" field.
func OofShardEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEQ(FieldOofShard, v))
}

// OofShardNEQ applies the NEQ predicate on the "oof_shard" field.
func OofShardNEQ(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNEQ(FieldOofShard, v))
}

// OofShardIn applies the In predicate on the "oof_shard" field.
func OofShardIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldIn(FieldOofShard, vs...))
}

// OofShardNotIn applies the NotIn predicate on the "oof_shard" field.
func OofShardNotIn(vs ...string) predicate.OrderData {
	return predicate.OrderData(sql.FieldNotIn(FieldOofShard, vs...))
}

// OofShardGT applies the GT predicate on the "oof_shard" field.
func OofShardGT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGT(FieldOofShard, v))
}

// OofShardGTE applies the GTE predicate on the "oof_shard" field.
func OofShardGTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldGTE(FieldOofShard, v))
}

// OofShardLT applies the LT predicate on the "oof_shard" field.
func OofShardLT(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLT(FieldOofShard, v))
}

// OofShardLTE applies the LTE predicate on the "oof_shard" field.
func OofShardLTE(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldLTE(FieldOofShard, v))
}

// OofShardContains applies the Contains predicate on the "oof_shard" field.
func OofShardContains(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContains(FieldOofShard, v))
}

// OofShardHasPrefix applies the HasPrefix predicate on the "oof_shard" field.
func OofShardHasPrefix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasPrefix(FieldOofShard, v))
}

// OofShardHasSuffix applies the HasSuffix predicate on the "oof_shard" field.
func OofShardHasSuffix(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldHasSuffix(FieldOofShard, v))
}

// OofShardEqualFold applies the EqualFold predicate on the "oof_shard" field.
func OofShardEqualFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldEqualFold(FieldOofShard, v))
}

// OofShardContainsFold applies the ContainsFold predicate on the "oof_shard" field.
func OofShardContainsFold(v string) predicate.OrderData {
	return predicate.OrderData(sql.FieldContainsFold(FieldOofShard, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderData) predicate.OrderData {
	return predicate.OrderData(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderData) predicate.OrderData {
	return predicate.OrderData(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderData) predicate.OrderData {
	return predicate.OrderData(sql.NotPredicates(p))
}
