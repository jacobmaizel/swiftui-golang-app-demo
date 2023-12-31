// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notificationpreferences"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
)

// NotificationPreferencesDelete is the builder for deleting a NotificationPreferences entity.
type NotificationPreferencesDelete struct {
	config
	hooks    []Hook
	mutation *NotificationPreferencesMutation
}

// Where appends a list predicates to the NotificationPreferencesDelete builder.
func (npd *NotificationPreferencesDelete) Where(ps ...predicate.NotificationPreferences) *NotificationPreferencesDelete {
	npd.mutation.Where(ps...)
	return npd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (npd *NotificationPreferencesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, npd.sqlExec, npd.mutation, npd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (npd *NotificationPreferencesDelete) ExecX(ctx context.Context) int {
	n, err := npd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (npd *NotificationPreferencesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(notificationpreferences.Table, sqlgraph.NewFieldSpec(notificationpreferences.FieldID, field.TypeUUID))
	if ps := npd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, npd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	npd.mutation.done = true
	return affected, err
}

// NotificationPreferencesDeleteOne is the builder for deleting a single NotificationPreferences entity.
type NotificationPreferencesDeleteOne struct {
	npd *NotificationPreferencesDelete
}

// Where appends a list predicates to the NotificationPreferencesDelete builder.
func (npdo *NotificationPreferencesDeleteOne) Where(ps ...predicate.NotificationPreferences) *NotificationPreferencesDeleteOne {
	npdo.npd.mutation.Where(ps...)
	return npdo
}

// Exec executes the deletion query.
func (npdo *NotificationPreferencesDeleteOne) Exec(ctx context.Context) error {
	n, err := npdo.npd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notificationpreferences.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (npdo *NotificationPreferencesDeleteOne) ExecX(ctx context.Context) {
	if err := npdo.Exec(ctx); err != nil {
		panic(err)
	}
}
