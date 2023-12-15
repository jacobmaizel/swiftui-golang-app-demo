// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
)

// WorkoutDataDelete is the builder for deleting a WorkoutData entity.
type WorkoutDataDelete struct {
	config
	hooks    []Hook
	mutation *WorkoutDataMutation
}

// Where appends a list predicates to the WorkoutDataDelete builder.
func (wdd *WorkoutDataDelete) Where(ps ...predicate.WorkoutData) *WorkoutDataDelete {
	wdd.mutation.Where(ps...)
	return wdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wdd *WorkoutDataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wdd.sqlExec, wdd.mutation, wdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wdd *WorkoutDataDelete) ExecX(ctx context.Context) int {
	n, err := wdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wdd *WorkoutDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workoutdata.Table, sqlgraph.NewFieldSpec(workoutdata.FieldID, field.TypeUUID))
	if ps := wdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wdd.mutation.done = true
	return affected, err
}

// WorkoutDataDeleteOne is the builder for deleting a single WorkoutData entity.
type WorkoutDataDeleteOne struct {
	wdd *WorkoutDataDelete
}

// Where appends a list predicates to the WorkoutDataDelete builder.
func (wddo *WorkoutDataDeleteOne) Where(ps ...predicate.WorkoutData) *WorkoutDataDeleteOne {
	wddo.wdd.mutation.Where(ps...)
	return wddo
}

// Exec executes the deletion query.
func (wddo *WorkoutDataDeleteOne) Exec(ctx context.Context) error {
	n, err := wddo.wdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workoutdata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wddo *WorkoutDataDeleteOne) ExecX(ctx context.Context) {
	if err := wddo.Exec(ctx); err != nil {
		panic(err)
	}
}