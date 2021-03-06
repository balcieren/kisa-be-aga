// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/pkg/ent/url"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// URLCreate is the builder for creating a Url entity.
type URLCreate struct {
	config
	mutation *URLMutation
	hooks    []Hook
}

// SetShortID sets the "short_id" field.
func (uc *URLCreate) SetShortID(s string) *URLCreate {
	uc.mutation.SetShortID(s)
	return uc
}

// SetURL sets the "url" field.
func (uc *URLCreate) SetURL(s string) *URLCreate {
	uc.mutation.SetURL(s)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *URLCreate) SetCreatedAt(t time.Time) *URLCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *URLCreate) SetNillableCreatedAt(t *time.Time) *URLCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *URLCreate) SetID(u uuid.UUID) *URLCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *URLCreate) SetNillableID(u *uuid.UUID) *URLCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// Mutation returns the URLMutation object of the builder.
func (uc *URLCreate) Mutation() *URLMutation {
	return uc.mutation
}

// Save creates the Url in the database.
func (uc *URLCreate) Save(ctx context.Context) (*Url, error) {
	var (
		err  error
		node *Url
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*URLMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Url)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from URLMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *URLCreate) SaveX(ctx context.Context) *Url {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *URLCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *URLCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *URLCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := url.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := url.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *URLCreate) check() error {
	if _, ok := uc.mutation.ShortID(); !ok {
		return &ValidationError{Name: "short_id", err: errors.New(`ent: missing required field "Url.short_id"`)}
	}
	if _, ok := uc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Url.url"`)}
	}
	if v, ok := uc.mutation.URL(); ok {
		if err := url.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Url.url": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Url.created_at"`)}
	}
	return nil
}

func (uc *URLCreate) sqlSave(ctx context.Context) (*Url, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (uc *URLCreate) createSpec() (*Url, *sqlgraph.CreateSpec) {
	var (
		_node = &Url{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: url.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: url.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.ShortID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: url.FieldShortID,
		})
		_node.ShortID = value
	}
	if value, ok := uc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: url.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: url.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	return _node, _spec
}

// URLCreateBulk is the builder for creating many Url entities in bulk.
type URLCreateBulk struct {
	config
	builders []*URLCreate
}

// Save creates the Url entities in the database.
func (ucb *URLCreateBulk) Save(ctx context.Context) ([]*Url, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Url, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*URLMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *URLCreateBulk) SaveX(ctx context.Context) []*Url {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *URLCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *URLCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
