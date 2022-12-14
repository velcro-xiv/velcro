// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/velcro-xiv/velcro/ent/logevent"
	"github.com/velcro-xiv/velcro/ent/predicate"
)

// LogEventDelete is the builder for deleting a LogEvent entity.
type LogEventDelete struct {
	config
	hooks    []Hook
	mutation *LogEventMutation
}

// Where appends a list predicates to the LogEventDelete builder.
func (led *LogEventDelete) Where(ps ...predicate.LogEvent) *LogEventDelete {
	led.mutation.Where(ps...)
	return led
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (led *LogEventDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(led.hooks) == 0 {
		affected, err = led.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LogEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			led.mutation = mutation
			affected, err = led.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(led.hooks) - 1; i >= 0; i-- {
			if led.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = led.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, led.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (led *LogEventDelete) ExecX(ctx context.Context) int {
	n, err := led.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (led *LogEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: logevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logevent.FieldID,
			},
		},
	}
	if ps := led.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, led.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// LogEventDeleteOne is the builder for deleting a single LogEvent entity.
type LogEventDeleteOne struct {
	led *LogEventDelete
}

// Exec executes the deletion query.
func (ledo *LogEventDeleteOne) Exec(ctx context.Context) error {
	n, err := ledo.led.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{logevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ledo *LogEventDeleteOne) ExecX(ctx context.Context) {
	ledo.led.ExecX(ctx)
}
