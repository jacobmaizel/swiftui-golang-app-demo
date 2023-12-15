// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notificationpreferences"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// NotificationPreferencesUpdate is the builder for updating NotificationPreferences entities.
type NotificationPreferencesUpdate struct {
	config
	hooks     []Hook
	mutation  *NotificationPreferencesMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the NotificationPreferencesUpdate builder.
func (npu *NotificationPreferencesUpdate) Where(ps ...predicate.NotificationPreferences) *NotificationPreferencesUpdate {
	npu.mutation.Where(ps...)
	return npu
}

// SetUpdatedAt sets the "updated_at" field.
func (npu *NotificationPreferencesUpdate) SetUpdatedAt(t time.Time) *NotificationPreferencesUpdate {
	npu.mutation.SetUpdatedAt(t)
	return npu
}

// SetSettings sets the "settings" field.
func (npu *NotificationPreferencesUpdate) SetSettings(sps *shared.NotificationPreferenceSettings) *NotificationPreferencesUpdate {
	npu.mutation.SetSettings(sps)
	return npu
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (npu *NotificationPreferencesUpdate) SetProfileID(id uuid.UUID) *NotificationPreferencesUpdate {
	npu.mutation.SetProfileID(id)
	return npu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (npu *NotificationPreferencesUpdate) SetProfile(p *Profile) *NotificationPreferencesUpdate {
	return npu.SetProfileID(p.ID)
}

// Mutation returns the NotificationPreferencesMutation object of the builder.
func (npu *NotificationPreferencesUpdate) Mutation() *NotificationPreferencesMutation {
	return npu.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (npu *NotificationPreferencesUpdate) ClearProfile() *NotificationPreferencesUpdate {
	npu.mutation.ClearProfile()
	return npu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (npu *NotificationPreferencesUpdate) Save(ctx context.Context) (int, error) {
	npu.defaults()
	return withHooks(ctx, npu.sqlSave, npu.mutation, npu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (npu *NotificationPreferencesUpdate) SaveX(ctx context.Context) int {
	affected, err := npu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (npu *NotificationPreferencesUpdate) Exec(ctx context.Context) error {
	_, err := npu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (npu *NotificationPreferencesUpdate) ExecX(ctx context.Context) {
	if err := npu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (npu *NotificationPreferencesUpdate) defaults() {
	if _, ok := npu.mutation.UpdatedAt(); !ok {
		v := notificationpreferences.UpdateDefaultUpdatedAt()
		npu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (npu *NotificationPreferencesUpdate) check() error {
	if _, ok := npu.mutation.ProfileID(); npu.mutation.ProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NotificationPreferences.profile"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (npu *NotificationPreferencesUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotificationPreferencesUpdate {
	npu.modifiers = append(npu.modifiers, modifiers...)
	return npu
}

func (npu *NotificationPreferencesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := npu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notificationpreferences.Table, notificationpreferences.Columns, sqlgraph.NewFieldSpec(notificationpreferences.FieldID, field.TypeUUID))
	if ps := npu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := npu.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationpreferences.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := npu.mutation.Settings(); ok {
		_spec.SetField(notificationpreferences.FieldSettings, field.TypeJSON, value)
	}
	if npu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationpreferences.ProfileTable,
			Columns: []string{notificationpreferences.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := npu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationpreferences.ProfileTable,
			Columns: []string{notificationpreferences.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(npu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, npu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationpreferences.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	npu.mutation.done = true
	return n, nil
}

// NotificationPreferencesUpdateOne is the builder for updating a single NotificationPreferences entity.
type NotificationPreferencesUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *NotificationPreferencesMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (npuo *NotificationPreferencesUpdateOne) SetUpdatedAt(t time.Time) *NotificationPreferencesUpdateOne {
	npuo.mutation.SetUpdatedAt(t)
	return npuo
}

// SetSettings sets the "settings" field.
func (npuo *NotificationPreferencesUpdateOne) SetSettings(sps *shared.NotificationPreferenceSettings) *NotificationPreferencesUpdateOne {
	npuo.mutation.SetSettings(sps)
	return npuo
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (npuo *NotificationPreferencesUpdateOne) SetProfileID(id uuid.UUID) *NotificationPreferencesUpdateOne {
	npuo.mutation.SetProfileID(id)
	return npuo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (npuo *NotificationPreferencesUpdateOne) SetProfile(p *Profile) *NotificationPreferencesUpdateOne {
	return npuo.SetProfileID(p.ID)
}

// Mutation returns the NotificationPreferencesMutation object of the builder.
func (npuo *NotificationPreferencesUpdateOne) Mutation() *NotificationPreferencesMutation {
	return npuo.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (npuo *NotificationPreferencesUpdateOne) ClearProfile() *NotificationPreferencesUpdateOne {
	npuo.mutation.ClearProfile()
	return npuo
}

// Where appends a list predicates to the NotificationPreferencesUpdate builder.
func (npuo *NotificationPreferencesUpdateOne) Where(ps ...predicate.NotificationPreferences) *NotificationPreferencesUpdateOne {
	npuo.mutation.Where(ps...)
	return npuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (npuo *NotificationPreferencesUpdateOne) Select(field string, fields ...string) *NotificationPreferencesUpdateOne {
	npuo.fields = append([]string{field}, fields...)
	return npuo
}

// Save executes the query and returns the updated NotificationPreferences entity.
func (npuo *NotificationPreferencesUpdateOne) Save(ctx context.Context) (*NotificationPreferences, error) {
	npuo.defaults()
	return withHooks(ctx, npuo.sqlSave, npuo.mutation, npuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (npuo *NotificationPreferencesUpdateOne) SaveX(ctx context.Context) *NotificationPreferences {
	node, err := npuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (npuo *NotificationPreferencesUpdateOne) Exec(ctx context.Context) error {
	_, err := npuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (npuo *NotificationPreferencesUpdateOne) ExecX(ctx context.Context) {
	if err := npuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (npuo *NotificationPreferencesUpdateOne) defaults() {
	if _, ok := npuo.mutation.UpdatedAt(); !ok {
		v := notificationpreferences.UpdateDefaultUpdatedAt()
		npuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (npuo *NotificationPreferencesUpdateOne) check() error {
	if _, ok := npuo.mutation.ProfileID(); npuo.mutation.ProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NotificationPreferences.profile"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (npuo *NotificationPreferencesUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotificationPreferencesUpdateOne {
	npuo.modifiers = append(npuo.modifiers, modifiers...)
	return npuo
}

func (npuo *NotificationPreferencesUpdateOne) sqlSave(ctx context.Context) (_node *NotificationPreferences, err error) {
	if err := npuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notificationpreferences.Table, notificationpreferences.Columns, sqlgraph.NewFieldSpec(notificationpreferences.FieldID, field.TypeUUID))
	id, ok := npuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotificationPreferences.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := npuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationpreferences.FieldID)
		for _, f := range fields {
			if !notificationpreferences.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notificationpreferences.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := npuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := npuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationpreferences.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := npuo.mutation.Settings(); ok {
		_spec.SetField(notificationpreferences.FieldSettings, field.TypeJSON, value)
	}
	if npuo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationpreferences.ProfileTable,
			Columns: []string{notificationpreferences.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := npuo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   notificationpreferences.ProfileTable,
			Columns: []string{notificationpreferences.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(npuo.modifiers...)
	_node = &NotificationPreferences{config: npuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, npuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationpreferences.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	npuo.mutation.done = true
	return _node, nil
}