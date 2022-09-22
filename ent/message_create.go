// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/karashiiro/velcro/ent/message"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
}

// SetTimestamp sets the "timestamp" field.
func (mc *MessageCreate) SetTimestamp(t time.Time) *MessageCreate {
	mc.mutation.SetTimestamp(t)
	return mc
}

// SetVersion sets the "version" field.
func (mc *MessageCreate) SetVersion(i int) *MessageCreate {
	mc.mutation.SetVersion(i)
	return mc
}

// SetSourceAddress sets the "source_address" field.
func (mc *MessageCreate) SetSourceAddress(s string) *MessageCreate {
	mc.mutation.SetSourceAddress(s)
	return mc
}

// SetSourcePort sets the "source_port" field.
func (mc *MessageCreate) SetSourcePort(i int) *MessageCreate {
	mc.mutation.SetSourcePort(i)
	return mc
}

// SetDestinationAddress sets the "destination_address" field.
func (mc *MessageCreate) SetDestinationAddress(s string) *MessageCreate {
	mc.mutation.SetDestinationAddress(s)
	return mc
}

// SetDestinationPort sets the "destination_port" field.
func (mc *MessageCreate) SetDestinationPort(i int) *MessageCreate {
	mc.mutation.SetDestinationPort(i)
	return mc
}

// SetData sets the "data" field.
func (mc *MessageCreate) SetData(b []byte) *MessageCreate {
	mc.mutation.SetData(b)
	return mc
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Message)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Message.timestamp"`)}
	}
	if _, ok := mc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Message.version"`)}
	}
	if _, ok := mc.mutation.SourceAddress(); !ok {
		return &ValidationError{Name: "source_address", err: errors.New(`ent: missing required field "Message.source_address"`)}
	}
	if _, ok := mc.mutation.SourcePort(); !ok {
		return &ValidationError{Name: "source_port", err: errors.New(`ent: missing required field "Message.source_port"`)}
	}
	if _, ok := mc.mutation.DestinationAddress(); !ok {
		return &ValidationError{Name: "destination_address", err: errors.New(`ent: missing required field "Message.destination_address"`)}
	}
	if _, ok := mc.mutation.DestinationPort(); !ok {
		return &ValidationError{Name: "destination_port", err: errors.New(`ent: missing required field "Message.destination_port"`)}
	}
	if _, ok := mc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Message.data"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	if value, ok := mc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldVersion,
		})
		_node.Version = value
	}
	if value, ok := mc.mutation.SourceAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldSourceAddress,
		})
		_node.SourceAddress = value
	}
	if value, ok := mc.mutation.SourcePort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldSourcePort,
		})
		_node.SourcePort = value
	}
	if value, ok := mc.mutation.DestinationAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldDestinationAddress,
		})
		_node.DestinationAddress = value
	}
	if value, ok := mc.mutation.DestinationPort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: message.FieldDestinationPort,
		})
		_node.DestinationPort = value
	}
	if value, ok := mc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: message.FieldData,
		})
		_node.Data = value
	}
	return _node, _spec
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
