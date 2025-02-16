// Code generated by ent, DO NOT EDIT.

package core

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"freefrom.space/nobot/core/botevents"
	"freefrom.space/nobot/core/internal"
	"freefrom.space/nobot/core/predicate"
)

// BotEventsDelete is the builder for deleting a BotEvents entity.
type BotEventsDelete struct {
	config
	hooks    []Hook
	mutation *BotEventsMutation
}

// Where appends a list predicates to the BotEventsDelete builder.
func (bed *BotEventsDelete) Where(ps ...predicate.BotEvents) *BotEventsDelete {
	bed.mutation.Where(ps...)
	return bed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bed *BotEventsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bed.sqlExec, bed.mutation, bed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bed *BotEventsDelete) ExecX(ctx context.Context) int {
	n, err := bed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bed *BotEventsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(botevents.Table, sqlgraph.NewFieldSpec(botevents.FieldID, field.TypeInt))
	_spec.Node.Schema = bed.schemaConfig.BotEvents
	ctx = internal.NewSchemaConfigContext(ctx, bed.schemaConfig)
	if ps := bed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bed.mutation.done = true
	return affected, err
}

// BotEventsDeleteOne is the builder for deleting a single BotEvents entity.
type BotEventsDeleteOne struct {
	bed *BotEventsDelete
}

// Where appends a list predicates to the BotEventsDelete builder.
func (bedo *BotEventsDeleteOne) Where(ps ...predicate.BotEvents) *BotEventsDeleteOne {
	bedo.bed.mutation.Where(ps...)
	return bedo
}

// Exec executes the deletion query.
func (bedo *BotEventsDeleteOne) Exec(ctx context.Context) error {
	n, err := bedo.bed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{botevents.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bedo *BotEventsDeleteOne) ExecX(ctx context.Context) {
	if err := bedo.Exec(ctx); err != nil {
		panic(err)
	}
}
