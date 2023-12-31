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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notification"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// NotificationCreate is the builder for creating a Notification entity.
type NotificationCreate struct {
	config
	mutation *NotificationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (nc *NotificationCreate) SetCreatedAt(t time.Time) *NotificationCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableCreatedAt(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NotificationCreate) SetUpdatedAt(t time.Time) *NotificationCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableUpdatedAt(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetTitle sets the "title" field.
func (nc *NotificationCreate) SetTitle(s string) *NotificationCreate {
	nc.mutation.SetTitle(s)
	return nc
}

// SetBody sets the "body" field.
func (nc *NotificationCreate) SetBody(s string) *NotificationCreate {
	nc.mutation.SetBody(s)
	return nc
}

// SetSent sets the "sent" field.
func (nc *NotificationCreate) SetSent(t time.Time) *NotificationCreate {
	nc.mutation.SetSent(t)
	return nc
}

// SetOpened sets the "opened" field.
func (nc *NotificationCreate) SetOpened(t time.Time) *NotificationCreate {
	nc.mutation.SetOpened(t)
	return nc
}

// SetNillableOpened sets the "opened" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableOpened(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetOpened(*t)
	}
	return nc
}

// SetData sets the "data" field.
func (nc *NotificationCreate) SetData(sd *shared.NotificationData) *NotificationCreate {
	nc.mutation.SetData(sd)
	return nc
}

// SetID sets the "id" field.
func (nc *NotificationCreate) SetID(u uuid.UUID) *NotificationCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableID(u *uuid.UUID) *NotificationCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (nc *NotificationCreate) SetProfileID(id uuid.UUID) *NotificationCreate {
	nc.mutation.SetProfileID(id)
	return nc
}

// SetProfile sets the "profile" edge to the Profile entity.
func (nc *NotificationCreate) SetProfile(p *Profile) *NotificationCreate {
	return nc.SetProfileID(p.ID)
}

// Mutation returns the NotificationMutation object of the builder.
func (nc *NotificationCreate) Mutation() *NotificationMutation {
	return nc.mutation
}

// Save creates the Notification in the database.
func (nc *NotificationCreate) Save(ctx context.Context) (*Notification, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotificationCreate) SaveX(ctx context.Context) *Notification {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotificationCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotificationCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NotificationCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := notification.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := notification.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := notification.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotificationCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Notification.created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Notification.updated_at"`)}
	}
	if _, ok := nc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Notification.title"`)}
	}
	if _, ok := nc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Notification.body"`)}
	}
	if _, ok := nc.mutation.Sent(); !ok {
		return &ValidationError{Name: "sent", err: errors.New(`ent: missing required field "Notification.sent"`)}
	}
	if _, ok := nc.mutation.ProfileID(); !ok {
		return &ValidationError{Name: "profile", err: errors.New(`ent: missing required edge "Notification.profile"`)}
	}
	return nil
}

func (nc *NotificationCreate) sqlSave(ctx context.Context) (*Notification, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
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
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NotificationCreate) createSpec() (*Notification, *sqlgraph.CreateSpec) {
	var (
		_node = &Notification{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(notification.Table, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID))
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.Title(); ok {
		_spec.SetField(notification.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := nc.mutation.Body(); ok {
		_spec.SetField(notification.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := nc.mutation.Sent(); ok {
		_spec.SetField(notification.FieldSent, field.TypeTime, value)
		_node.Sent = value
	}
	if value, ok := nc.mutation.Opened(); ok {
		_spec.SetField(notification.FieldOpened, field.TypeTime, value)
		_node.Opened = &value
	}
	if value, ok := nc.mutation.Data(); ok {
		_spec.SetField(notification.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	if nodes := nc.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.ProfileTable,
			Columns: []string{notification.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.profile_notifications = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NotificationCreateBulk is the builder for creating many Notification entities in bulk.
type NotificationCreateBulk struct {
	config
	builders []*NotificationCreate
}

// Save creates the Notification entities in the database.
func (ncb *NotificationCreateBulk) Save(ctx context.Context) ([]*Notification, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notification, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotificationCreateBulk) SaveX(ctx context.Context) []*Notification {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotificationCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotificationCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
