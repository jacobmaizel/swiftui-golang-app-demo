// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notificationpreferences"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// NotificationPreferencesCreate is the builder for creating a NotificationPreferences entity.
type NotificationPreferencesCreate struct {
	config
	mutation *NotificationPreferencesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (npc *NotificationPreferencesCreate) SetCreatedAt(t time.Time) *NotificationPreferencesCreate {
	npc.mutation.SetCreatedAt(t)
	return npc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (npc *NotificationPreferencesCreate) SetNillableCreatedAt(t *time.Time) *NotificationPreferencesCreate {
	if t != nil {
		npc.SetCreatedAt(*t)
	}
	return npc
}

// SetUpdatedAt sets the "updated_at" field.
func (npc *NotificationPreferencesCreate) SetUpdatedAt(t time.Time) *NotificationPreferencesCreate {
	npc.mutation.SetUpdatedAt(t)
	return npc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (npc *NotificationPreferencesCreate) SetNillableUpdatedAt(t *time.Time) *NotificationPreferencesCreate {
	if t != nil {
		npc.SetUpdatedAt(*t)
	}
	return npc
}

// SetSettings sets the "settings" field.
func (npc *NotificationPreferencesCreate) SetSettings(sps *shared.NotificationPreferenceSettings) *NotificationPreferencesCreate {
	npc.mutation.SetSettings(sps)
	return npc
}

// SetID sets the "id" field.
func (npc *NotificationPreferencesCreate) SetID(u uuid.UUID) *NotificationPreferencesCreate {
	npc.mutation.SetID(u)
	return npc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (npc *NotificationPreferencesCreate) SetNillableID(u *uuid.UUID) *NotificationPreferencesCreate {
	if u != nil {
		npc.SetID(*u)
	}
	return npc
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (npc *NotificationPreferencesCreate) SetProfileID(id uuid.UUID) *NotificationPreferencesCreate {
	npc.mutation.SetProfileID(id)
	return npc
}

// SetProfile sets the "profile" edge to the Profile entity.
func (npc *NotificationPreferencesCreate) SetProfile(p *Profile) *NotificationPreferencesCreate {
	return npc.SetProfileID(p.ID)
}

// Mutation returns the NotificationPreferencesMutation object of the builder.
func (npc *NotificationPreferencesCreate) Mutation() *NotificationPreferencesMutation {
	return npc.mutation
}

// Save creates the NotificationPreferences in the database.
func (npc *NotificationPreferencesCreate) Save(ctx context.Context) (*NotificationPreferences, error) {
	npc.defaults()
	return withHooks(ctx, npc.sqlSave, npc.mutation, npc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (npc *NotificationPreferencesCreate) SaveX(ctx context.Context) *NotificationPreferences {
	v, err := npc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (npc *NotificationPreferencesCreate) Exec(ctx context.Context) error {
	_, err := npc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (npc *NotificationPreferencesCreate) ExecX(ctx context.Context) {
	if err := npc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (npc *NotificationPreferencesCreate) defaults() {
	if _, ok := npc.mutation.CreatedAt(); !ok {
		v := notificationpreferences.DefaultCreatedAt()
		npc.mutation.SetCreatedAt(v)
	}
	if _, ok := npc.mutation.UpdatedAt(); !ok {
		v := notificationpreferences.DefaultUpdatedAt()
		npc.mutation.SetUpdatedAt(v)
	}
	if _, ok := npc.mutation.ID(); !ok {
		v := notificationpreferences.DefaultID()
		npc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (npc *NotificationPreferencesCreate) check() error {
	if _, ok := npc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NotificationPreferences.created_at"`)}
	}
	if _, ok := npc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "NotificationPreferences.updated_at"`)}
	}
	if _, ok := npc.mutation.Settings(); !ok {
		return &ValidationError{Name: "settings", err: errors.New(`ent: missing required field "NotificationPreferences.settings"`)}
	}
	if _, ok := npc.mutation.ProfileID(); !ok {
		return &ValidationError{Name: "profile", err: errors.New(`ent: missing required edge "NotificationPreferences.profile"`)}
	}
	return nil
}

func (npc *NotificationPreferencesCreate) sqlSave(ctx context.Context) (*NotificationPreferences, error) {
	if err := npc.check(); err != nil {
		return nil, err
	}
	_node, _spec := npc.createSpec()
	if err := sqlgraph.CreateNode(ctx, npc.driver, _spec); err != nil {
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
	npc.mutation.id = &_node.ID
	npc.mutation.done = true
	return _node, nil
}

func (npc *NotificationPreferencesCreate) createSpec() (*NotificationPreferences, *sqlgraph.CreateSpec) {
	var (
		_node = &NotificationPreferences{config: npc.config}
		_spec = sqlgraph.NewCreateSpec(notificationpreferences.Table, sqlgraph.NewFieldSpec(notificationpreferences.FieldID, field.TypeUUID))
	)
	if id, ok := npc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := npc.mutation.CreatedAt(); ok {
		_spec.SetField(notificationpreferences.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := npc.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationpreferences.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := npc.mutation.Settings(); ok {
		_spec.SetField(notificationpreferences.FieldSettings, field.TypeJSON, value)
		_node.Settings = value
	}
	if nodes := npc.mutation.ProfileIDs(); len(nodes) > 0 {
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
		_node.profile_notification_preferences = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NotificationPreferencesCreateBulk is the builder for creating many NotificationPreferences entities in bulk.
type NotificationPreferencesCreateBulk struct {
	config
	builders []*NotificationPreferencesCreate
}

// Save creates the NotificationPreferences entities in the database.
func (npcb *NotificationPreferencesCreateBulk) Save(ctx context.Context) ([]*NotificationPreferences, error) {
	specs := make([]*sqlgraph.CreateSpec, len(npcb.builders))
	nodes := make([]*NotificationPreferences, len(npcb.builders))
	mutators := make([]Mutator, len(npcb.builders))
	for i := range npcb.builders {
		func(i int, root context.Context) {
			builder := npcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationPreferencesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, npcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, npcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, npcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (npcb *NotificationPreferencesCreateBulk) SaveX(ctx context.Context) []*NotificationPreferences {
	v, err := npcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (npcb *NotificationPreferencesCreateBulk) Exec(ctx context.Context) error {
	_, err := npcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (npcb *NotificationPreferencesCreateBulk) ExecX(ctx context.Context) {
	if err := npcb.Exec(ctx); err != nil {
		panic(err)
	}
}
