// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nats-streaming-web/ent/orderdata"
	"nats-streaming-web/ent/predicate"
	"nats-streaming-web/pkg/model"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeOrderData = "OrderData"
)

// OrderDataMutation represents an operation that mutates the OrderData nodes in the graph.
type OrderDataMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	order_uid          *string
	track_number       *string
	entry              *string
	delivery           **model.Delivery
	payment            **model.Payment
	items              *[]*model.Item
	appenditems        []*model.Item
	locale             *string
	internal_signature *string
	customer_id        *string
	delivery_service   *string
	shardkey           *string
	sm_id              *int
	addsm_id           *int
	date_created       *string
	oof_shard          *string
	clearedFields      map[string]struct{}
	done               bool
	oldValue           func(context.Context) (*OrderData, error)
	predicates         []predicate.OrderData
}

var _ ent.Mutation = (*OrderDataMutation)(nil)

// orderdataOption allows management of the mutation configuration using functional options.
type orderdataOption func(*OrderDataMutation)

// newOrderDataMutation creates new mutation for the OrderData entity.
func newOrderDataMutation(c config, op Op, opts ...orderdataOption) *OrderDataMutation {
	m := &OrderDataMutation{
		config:        c,
		op:            op,
		typ:           TypeOrderData,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOrderDataID sets the ID field of the mutation.
func withOrderDataID(id int) orderdataOption {
	return func(m *OrderDataMutation) {
		var (
			err   error
			once  sync.Once
			value *OrderData
		)
		m.oldValue = func(ctx context.Context) (*OrderData, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().OrderData.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOrderData sets the old OrderData of the mutation.
func withOrderData(node *OrderData) orderdataOption {
	return func(m *OrderDataMutation) {
		m.oldValue = func(context.Context) (*OrderData, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OrderDataMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OrderDataMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OrderDataMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *OrderDataMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().OrderData.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetOrderUID sets the "order_uid" field.
func (m *OrderDataMutation) SetOrderUID(s string) {
	m.order_uid = &s
}

// OrderUID returns the value of the "order_uid" field in the mutation.
func (m *OrderDataMutation) OrderUID() (r string, exists bool) {
	v := m.order_uid
	if v == nil {
		return
	}
	return *v, true
}

// OldOrderUID returns the old "order_uid" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldOrderUID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOrderUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOrderUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOrderUID: %w", err)
	}
	return oldValue.OrderUID, nil
}

// ResetOrderUID resets all changes to the "order_uid" field.
func (m *OrderDataMutation) ResetOrderUID() {
	m.order_uid = nil
}

// SetTrackNumber sets the "track_number" field.
func (m *OrderDataMutation) SetTrackNumber(s string) {
	m.track_number = &s
}

// TrackNumber returns the value of the "track_number" field in the mutation.
func (m *OrderDataMutation) TrackNumber() (r string, exists bool) {
	v := m.track_number
	if v == nil {
		return
	}
	return *v, true
}

// OldTrackNumber returns the old "track_number" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldTrackNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTrackNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTrackNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTrackNumber: %w", err)
	}
	return oldValue.TrackNumber, nil
}

// ResetTrackNumber resets all changes to the "track_number" field.
func (m *OrderDataMutation) ResetTrackNumber() {
	m.track_number = nil
}

// SetEntry sets the "entry" field.
func (m *OrderDataMutation) SetEntry(s string) {
	m.entry = &s
}

// Entry returns the value of the "entry" field in the mutation.
func (m *OrderDataMutation) Entry() (r string, exists bool) {
	v := m.entry
	if v == nil {
		return
	}
	return *v, true
}

// OldEntry returns the old "entry" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldEntry(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEntry is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEntry requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEntry: %w", err)
	}
	return oldValue.Entry, nil
}

// ResetEntry resets all changes to the "entry" field.
func (m *OrderDataMutation) ResetEntry() {
	m.entry = nil
}

// SetDelivery sets the "delivery" field.
func (m *OrderDataMutation) SetDelivery(value *model.Delivery) {
	m.delivery = &value
}

// Delivery returns the value of the "delivery" field in the mutation.
func (m *OrderDataMutation) Delivery() (r *model.Delivery, exists bool) {
	v := m.delivery
	if v == nil {
		return
	}
	return *v, true
}

// OldDelivery returns the old "delivery" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldDelivery(ctx context.Context) (v *model.Delivery, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDelivery is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDelivery requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDelivery: %w", err)
	}
	return oldValue.Delivery, nil
}

// ResetDelivery resets all changes to the "delivery" field.
func (m *OrderDataMutation) ResetDelivery() {
	m.delivery = nil
}

// SetPayment sets the "payment" field.
func (m *OrderDataMutation) SetPayment(value *model.Payment) {
	m.payment = &value
}

// Payment returns the value of the "payment" field in the mutation.
func (m *OrderDataMutation) Payment() (r *model.Payment, exists bool) {
	v := m.payment
	if v == nil {
		return
	}
	return *v, true
}

// OldPayment returns the old "payment" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldPayment(ctx context.Context) (v *model.Payment, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPayment is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPayment requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPayment: %w", err)
	}
	return oldValue.Payment, nil
}

// ResetPayment resets all changes to the "payment" field.
func (m *OrderDataMutation) ResetPayment() {
	m.payment = nil
}

// SetItems sets the "items" field.
func (m *OrderDataMutation) SetItems(value []*model.Item) {
	m.items = &value
	m.appenditems = nil
}

// Items returns the value of the "items" field in the mutation.
func (m *OrderDataMutation) Items() (r []*model.Item, exists bool) {
	v := m.items
	if v == nil {
		return
	}
	return *v, true
}

// OldItems returns the old "items" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldItems(ctx context.Context) (v []*model.Item, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldItems is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldItems requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldItems: %w", err)
	}
	return oldValue.Items, nil
}

// AppendItems adds value to the "items" field.
func (m *OrderDataMutation) AppendItems(value []*model.Item) {
	m.appenditems = append(m.appenditems, value...)
}

// AppendedItems returns the list of values that were appended to the "items" field in this mutation.
func (m *OrderDataMutation) AppendedItems() ([]*model.Item, bool) {
	if len(m.appenditems) == 0 {
		return nil, false
	}
	return m.appenditems, true
}

// ResetItems resets all changes to the "items" field.
func (m *OrderDataMutation) ResetItems() {
	m.items = nil
	m.appenditems = nil
}

// SetLocale sets the "locale" field.
func (m *OrderDataMutation) SetLocale(s string) {
	m.locale = &s
}

// Locale returns the value of the "locale" field in the mutation.
func (m *OrderDataMutation) Locale() (r string, exists bool) {
	v := m.locale
	if v == nil {
		return
	}
	return *v, true
}

// OldLocale returns the old "locale" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldLocale(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLocale is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLocale requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLocale: %w", err)
	}
	return oldValue.Locale, nil
}

// ResetLocale resets all changes to the "locale" field.
func (m *OrderDataMutation) ResetLocale() {
	m.locale = nil
}

// SetInternalSignature sets the "internal_signature" field.
func (m *OrderDataMutation) SetInternalSignature(s string) {
	m.internal_signature = &s
}

// InternalSignature returns the value of the "internal_signature" field in the mutation.
func (m *OrderDataMutation) InternalSignature() (r string, exists bool) {
	v := m.internal_signature
	if v == nil {
		return
	}
	return *v, true
}

// OldInternalSignature returns the old "internal_signature" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldInternalSignature(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInternalSignature is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInternalSignature requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInternalSignature: %w", err)
	}
	return oldValue.InternalSignature, nil
}

// ResetInternalSignature resets all changes to the "internal_signature" field.
func (m *OrderDataMutation) ResetInternalSignature() {
	m.internal_signature = nil
}

// SetCustomerID sets the "customer_id" field.
func (m *OrderDataMutation) SetCustomerID(s string) {
	m.customer_id = &s
}

// CustomerID returns the value of the "customer_id" field in the mutation.
func (m *OrderDataMutation) CustomerID() (r string, exists bool) {
	v := m.customer_id
	if v == nil {
		return
	}
	return *v, true
}

// OldCustomerID returns the old "customer_id" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldCustomerID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCustomerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCustomerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustomerID: %w", err)
	}
	return oldValue.CustomerID, nil
}

// ResetCustomerID resets all changes to the "customer_id" field.
func (m *OrderDataMutation) ResetCustomerID() {
	m.customer_id = nil
}

// SetDeliveryService sets the "delivery_service" field.
func (m *OrderDataMutation) SetDeliveryService(s string) {
	m.delivery_service = &s
}

// DeliveryService returns the value of the "delivery_service" field in the mutation.
func (m *OrderDataMutation) DeliveryService() (r string, exists bool) {
	v := m.delivery_service
	if v == nil {
		return
	}
	return *v, true
}

// OldDeliveryService returns the old "delivery_service" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldDeliveryService(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeliveryService is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeliveryService requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeliveryService: %w", err)
	}
	return oldValue.DeliveryService, nil
}

// ResetDeliveryService resets all changes to the "delivery_service" field.
func (m *OrderDataMutation) ResetDeliveryService() {
	m.delivery_service = nil
}

// SetShardkey sets the "shardkey" field.
func (m *OrderDataMutation) SetShardkey(s string) {
	m.shardkey = &s
}

// Shardkey returns the value of the "shardkey" field in the mutation.
func (m *OrderDataMutation) Shardkey() (r string, exists bool) {
	v := m.shardkey
	if v == nil {
		return
	}
	return *v, true
}

// OldShardkey returns the old "shardkey" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldShardkey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldShardkey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldShardkey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldShardkey: %w", err)
	}
	return oldValue.Shardkey, nil
}

// ResetShardkey resets all changes to the "shardkey" field.
func (m *OrderDataMutation) ResetShardkey() {
	m.shardkey = nil
}

// SetSmID sets the "sm_id" field.
func (m *OrderDataMutation) SetSmID(i int) {
	m.sm_id = &i
	m.addsm_id = nil
}

// SmID returns the value of the "sm_id" field in the mutation.
func (m *OrderDataMutation) SmID() (r int, exists bool) {
	v := m.sm_id
	if v == nil {
		return
	}
	return *v, true
}

// OldSmID returns the old "sm_id" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldSmID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSmID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSmID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSmID: %w", err)
	}
	return oldValue.SmID, nil
}

// AddSmID adds i to the "sm_id" field.
func (m *OrderDataMutation) AddSmID(i int) {
	if m.addsm_id != nil {
		*m.addsm_id += i
	} else {
		m.addsm_id = &i
	}
}

// AddedSmID returns the value that was added to the "sm_id" field in this mutation.
func (m *OrderDataMutation) AddedSmID() (r int, exists bool) {
	v := m.addsm_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetSmID resets all changes to the "sm_id" field.
func (m *OrderDataMutation) ResetSmID() {
	m.sm_id = nil
	m.addsm_id = nil
}

// SetDateCreated sets the "date_created" field.
func (m *OrderDataMutation) SetDateCreated(s string) {
	m.date_created = &s
}

// DateCreated returns the value of the "date_created" field in the mutation.
func (m *OrderDataMutation) DateCreated() (r string, exists bool) {
	v := m.date_created
	if v == nil {
		return
	}
	return *v, true
}

// OldDateCreated returns the old "date_created" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldDateCreated(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDateCreated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDateCreated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDateCreated: %w", err)
	}
	return oldValue.DateCreated, nil
}

// ResetDateCreated resets all changes to the "date_created" field.
func (m *OrderDataMutation) ResetDateCreated() {
	m.date_created = nil
}

// SetOofShard sets the "oof_shard" field.
func (m *OrderDataMutation) SetOofShard(s string) {
	m.oof_shard = &s
}

// OofShard returns the value of the "oof_shard" field in the mutation.
func (m *OrderDataMutation) OofShard() (r string, exists bool) {
	v := m.oof_shard
	if v == nil {
		return
	}
	return *v, true
}

// OldOofShard returns the old "oof_shard" field's value of the OrderData entity.
// If the OrderData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderDataMutation) OldOofShard(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOofShard is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOofShard requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOofShard: %w", err)
	}
	return oldValue.OofShard, nil
}

// ResetOofShard resets all changes to the "oof_shard" field.
func (m *OrderDataMutation) ResetOofShard() {
	m.oof_shard = nil
}

// Where appends a list predicates to the OrderDataMutation builder.
func (m *OrderDataMutation) Where(ps ...predicate.OrderData) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the OrderDataMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *OrderDataMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.OrderData, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *OrderDataMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *OrderDataMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (OrderData).
func (m *OrderDataMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OrderDataMutation) Fields() []string {
	fields := make([]string, 0, 14)
	if m.order_uid != nil {
		fields = append(fields, orderdata.FieldOrderUID)
	}
	if m.track_number != nil {
		fields = append(fields, orderdata.FieldTrackNumber)
	}
	if m.entry != nil {
		fields = append(fields, orderdata.FieldEntry)
	}
	if m.delivery != nil {
		fields = append(fields, orderdata.FieldDelivery)
	}
	if m.payment != nil {
		fields = append(fields, orderdata.FieldPayment)
	}
	if m.items != nil {
		fields = append(fields, orderdata.FieldItems)
	}
	if m.locale != nil {
		fields = append(fields, orderdata.FieldLocale)
	}
	if m.internal_signature != nil {
		fields = append(fields, orderdata.FieldInternalSignature)
	}
	if m.customer_id != nil {
		fields = append(fields, orderdata.FieldCustomerID)
	}
	if m.delivery_service != nil {
		fields = append(fields, orderdata.FieldDeliveryService)
	}
	if m.shardkey != nil {
		fields = append(fields, orderdata.FieldShardkey)
	}
	if m.sm_id != nil {
		fields = append(fields, orderdata.FieldSmID)
	}
	if m.date_created != nil {
		fields = append(fields, orderdata.FieldDateCreated)
	}
	if m.oof_shard != nil {
		fields = append(fields, orderdata.FieldOofShard)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OrderDataMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case orderdata.FieldOrderUID:
		return m.OrderUID()
	case orderdata.FieldTrackNumber:
		return m.TrackNumber()
	case orderdata.FieldEntry:
		return m.Entry()
	case orderdata.FieldDelivery:
		return m.Delivery()
	case orderdata.FieldPayment:
		return m.Payment()
	case orderdata.FieldItems:
		return m.Items()
	case orderdata.FieldLocale:
		return m.Locale()
	case orderdata.FieldInternalSignature:
		return m.InternalSignature()
	case orderdata.FieldCustomerID:
		return m.CustomerID()
	case orderdata.FieldDeliveryService:
		return m.DeliveryService()
	case orderdata.FieldShardkey:
		return m.Shardkey()
	case orderdata.FieldSmID:
		return m.SmID()
	case orderdata.FieldDateCreated:
		return m.DateCreated()
	case orderdata.FieldOofShard:
		return m.OofShard()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OrderDataMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case orderdata.FieldOrderUID:
		return m.OldOrderUID(ctx)
	case orderdata.FieldTrackNumber:
		return m.OldTrackNumber(ctx)
	case orderdata.FieldEntry:
		return m.OldEntry(ctx)
	case orderdata.FieldDelivery:
		return m.OldDelivery(ctx)
	case orderdata.FieldPayment:
		return m.OldPayment(ctx)
	case orderdata.FieldItems:
		return m.OldItems(ctx)
	case orderdata.FieldLocale:
		return m.OldLocale(ctx)
	case orderdata.FieldInternalSignature:
		return m.OldInternalSignature(ctx)
	case orderdata.FieldCustomerID:
		return m.OldCustomerID(ctx)
	case orderdata.FieldDeliveryService:
		return m.OldDeliveryService(ctx)
	case orderdata.FieldShardkey:
		return m.OldShardkey(ctx)
	case orderdata.FieldSmID:
		return m.OldSmID(ctx)
	case orderdata.FieldDateCreated:
		return m.OldDateCreated(ctx)
	case orderdata.FieldOofShard:
		return m.OldOofShard(ctx)
	}
	return nil, fmt.Errorf("unknown OrderData field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderDataMutation) SetField(name string, value ent.Value) error {
	switch name {
	case orderdata.FieldOrderUID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOrderUID(v)
		return nil
	case orderdata.FieldTrackNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTrackNumber(v)
		return nil
	case orderdata.FieldEntry:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEntry(v)
		return nil
	case orderdata.FieldDelivery:
		v, ok := value.(*model.Delivery)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDelivery(v)
		return nil
	case orderdata.FieldPayment:
		v, ok := value.(*model.Payment)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPayment(v)
		return nil
	case orderdata.FieldItems:
		v, ok := value.([]*model.Item)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetItems(v)
		return nil
	case orderdata.FieldLocale:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLocale(v)
		return nil
	case orderdata.FieldInternalSignature:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInternalSignature(v)
		return nil
	case orderdata.FieldCustomerID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustomerID(v)
		return nil
	case orderdata.FieldDeliveryService:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeliveryService(v)
		return nil
	case orderdata.FieldShardkey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetShardkey(v)
		return nil
	case orderdata.FieldSmID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSmID(v)
		return nil
	case orderdata.FieldDateCreated:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDateCreated(v)
		return nil
	case orderdata.FieldOofShard:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOofShard(v)
		return nil
	}
	return fmt.Errorf("unknown OrderData field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OrderDataMutation) AddedFields() []string {
	var fields []string
	if m.addsm_id != nil {
		fields = append(fields, orderdata.FieldSmID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OrderDataMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case orderdata.FieldSmID:
		return m.AddedSmID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderDataMutation) AddField(name string, value ent.Value) error {
	switch name {
	case orderdata.FieldSmID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSmID(v)
		return nil
	}
	return fmt.Errorf("unknown OrderData numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OrderDataMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OrderDataMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OrderDataMutation) ClearField(name string) error {
	return fmt.Errorf("unknown OrderData nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OrderDataMutation) ResetField(name string) error {
	switch name {
	case orderdata.FieldOrderUID:
		m.ResetOrderUID()
		return nil
	case orderdata.FieldTrackNumber:
		m.ResetTrackNumber()
		return nil
	case orderdata.FieldEntry:
		m.ResetEntry()
		return nil
	case orderdata.FieldDelivery:
		m.ResetDelivery()
		return nil
	case orderdata.FieldPayment:
		m.ResetPayment()
		return nil
	case orderdata.FieldItems:
		m.ResetItems()
		return nil
	case orderdata.FieldLocale:
		m.ResetLocale()
		return nil
	case orderdata.FieldInternalSignature:
		m.ResetInternalSignature()
		return nil
	case orderdata.FieldCustomerID:
		m.ResetCustomerID()
		return nil
	case orderdata.FieldDeliveryService:
		m.ResetDeliveryService()
		return nil
	case orderdata.FieldShardkey:
		m.ResetShardkey()
		return nil
	case orderdata.FieldSmID:
		m.ResetSmID()
		return nil
	case orderdata.FieldDateCreated:
		m.ResetDateCreated()
		return nil
	case orderdata.FieldOofShard:
		m.ResetOofShard()
		return nil
	}
	return fmt.Errorf("unknown OrderData field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OrderDataMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OrderDataMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OrderDataMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OrderDataMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OrderDataMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OrderDataMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OrderDataMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown OrderData unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OrderDataMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown OrderData edge %s", name)
}
