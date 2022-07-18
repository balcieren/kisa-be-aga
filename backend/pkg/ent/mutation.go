// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/pkg/ent/predicate"
	"backend/pkg/ent/url"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUrl = "Url"
)

// URLMutation represents an operation that mutates the Url nodes in the graph.
type URLMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	short_id      *string
	url           *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Url, error)
	predicates    []predicate.Url
}

var _ ent.Mutation = (*URLMutation)(nil)

// urlOption allows management of the mutation configuration using functional options.
type urlOption func(*URLMutation)

// newURLMutation creates new mutation for the Url entity.
func newURLMutation(c config, op Op, opts ...urlOption) *URLMutation {
	m := &URLMutation{
		config:        c,
		op:            op,
		typ:           TypeUrl,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUrlID sets the ID field of the mutation.
func withUrlID(id uuid.UUID) urlOption {
	return func(m *URLMutation) {
		var (
			err   error
			once  sync.Once
			value *Url
		)
		m.oldValue = func(ctx context.Context) (*Url, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Url.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUrl sets the old Url of the mutation.
func withUrl(node *Url) urlOption {
	return func(m *URLMutation) {
		m.oldValue = func(context.Context) (*Url, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m URLMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m URLMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Url entities.
func (m *URLMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *URLMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *URLMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Url.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetShortID sets the "short_id" field.
func (m *URLMutation) SetShortID(s string) {
	m.short_id = &s
}

// ShortID returns the value of the "short_id" field in the mutation.
func (m *URLMutation) ShortID() (r string, exists bool) {
	v := m.short_id
	if v == nil {
		return
	}
	return *v, true
}

// OldShortID returns the old "short_id" field's value of the Url entity.
// If the Url object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *URLMutation) OldShortID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldShortID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldShortID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldShortID: %w", err)
	}
	return oldValue.ShortID, nil
}

// ResetShortID resets all changes to the "short_id" field.
func (m *URLMutation) ResetShortID() {
	m.short_id = nil
}

// SetURL sets the "url" field.
func (m *URLMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *URLMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Url entity.
// If the Url object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *URLMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *URLMutation) ResetURL() {
	m.url = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *URLMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *URLMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Url entity.
// If the Url object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *URLMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *URLMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the URLMutation builder.
func (m *URLMutation) Where(ps ...predicate.Url) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *URLMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Url).
func (m *URLMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *URLMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.short_id != nil {
		fields = append(fields, url.FieldShortID)
	}
	if m.url != nil {
		fields = append(fields, url.FieldURL)
	}
	if m.created_at != nil {
		fields = append(fields, url.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *URLMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case url.FieldShortID:
		return m.ShortID()
	case url.FieldURL:
		return m.URL()
	case url.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *URLMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case url.FieldShortID:
		return m.OldShortID(ctx)
	case url.FieldURL:
		return m.OldURL(ctx)
	case url.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Url field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *URLMutation) SetField(name string, value ent.Value) error {
	switch name {
	case url.FieldShortID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetShortID(v)
		return nil
	case url.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case url.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Url field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *URLMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *URLMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *URLMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Url numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *URLMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *URLMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *URLMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Url nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *URLMutation) ResetField(name string) error {
	switch name {
	case url.FieldShortID:
		m.ResetShortID()
		return nil
	case url.FieldURL:
		m.ResetURL()
		return nil
	case url.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Url field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *URLMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *URLMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *URLMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *URLMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *URLMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *URLMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *URLMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Url unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *URLMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Url edge %s", name)
}
