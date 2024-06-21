// Code generated by ent, DO NOT EDIT.

package core

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"freefrom.space/nobot/core/events"
	"freefrom.space/nobot/core/internal"
	"freefrom.space/nobot/core/predicate"
)

// EventsDelete is the builder for deleting a Events entity.
type EventsDelete struct {
	config
	hooks    []Hook
	mutation *EventsMutation
}

// Where appends a list predicates to the EventsDelete builder.
func (ed *EventsDelete) Where(ps ...predicate.Events) *EventsDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EventsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ed.sqlExec, ed.mutation, ed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EventsDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EventsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(events.Table, sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt64))
	_spec.Node.Schema = ed.schemaConfig.Events
	ctx = internal.NewSchemaConfigContext(ctx, ed.schemaConfig)
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ed.mutation.done = true
	return affected, err
}

// EventsDeleteOne is the builder for deleting a single Events entity.
type EventsDeleteOne struct {
	ed *EventsDelete
}

// Where appends a list predicates to the EventsDelete builder.
func (edo *EventsDeleteOne) Where(ps ...predicate.Events) *EventsDeleteOne {
	edo.ed.mutation.Where(ps...)
	return edo
}

// Exec executes the deletion query.
func (edo *EventsDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{events.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EventsDeleteOne) ExecX(ctx context.Context) {
	if err := edo.Exec(ctx); err != nil {
		panic(err)
	}
}
