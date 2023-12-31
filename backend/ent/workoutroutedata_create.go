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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutroutedata"
)

// WorkoutRouteDataCreate is the builder for creating a WorkoutRouteData entity.
type WorkoutRouteDataCreate struct {
	config
	mutation *WorkoutRouteDataMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (wrdc *WorkoutRouteDataCreate) SetCreatedAt(t time.Time) *WorkoutRouteDataCreate {
	wrdc.mutation.SetCreatedAt(t)
	return wrdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wrdc *WorkoutRouteDataCreate) SetNillableCreatedAt(t *time.Time) *WorkoutRouteDataCreate {
	if t != nil {
		wrdc.SetCreatedAt(*t)
	}
	return wrdc
}

// SetUpdatedAt sets the "updated_at" field.
func (wrdc *WorkoutRouteDataCreate) SetUpdatedAt(t time.Time) *WorkoutRouteDataCreate {
	wrdc.mutation.SetUpdatedAt(t)
	return wrdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wrdc *WorkoutRouteDataCreate) SetNillableUpdatedAt(t *time.Time) *WorkoutRouteDataCreate {
	if t != nil {
		wrdc.SetUpdatedAt(*t)
	}
	return wrdc
}

// SetID sets the "id" field.
func (wrdc *WorkoutRouteDataCreate) SetID(u uuid.UUID) *WorkoutRouteDataCreate {
	wrdc.mutation.SetID(u)
	return wrdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wrdc *WorkoutRouteDataCreate) SetNillableID(u *uuid.UUID) *WorkoutRouteDataCreate {
	if u != nil {
		wrdc.SetID(*u)
	}
	return wrdc
}

// SetWorkoutID sets the "workout" edge to the Workout entity by ID.
func (wrdc *WorkoutRouteDataCreate) SetWorkoutID(id uuid.UUID) *WorkoutRouteDataCreate {
	wrdc.mutation.SetWorkoutID(id)
	return wrdc
}

// SetWorkout sets the "workout" edge to the Workout entity.
func (wrdc *WorkoutRouteDataCreate) SetWorkout(w *Workout) *WorkoutRouteDataCreate {
	return wrdc.SetWorkoutID(w.ID)
}

// SetWorkoutDataID sets the "workout_data" edge to the WorkoutData entity by ID.
func (wrdc *WorkoutRouteDataCreate) SetWorkoutDataID(id uuid.UUID) *WorkoutRouteDataCreate {
	wrdc.mutation.SetWorkoutDataID(id)
	return wrdc
}

// SetNillableWorkoutDataID sets the "workout_data" edge to the WorkoutData entity by ID if the given value is not nil.
func (wrdc *WorkoutRouteDataCreate) SetNillableWorkoutDataID(id *uuid.UUID) *WorkoutRouteDataCreate {
	if id != nil {
		wrdc = wrdc.SetWorkoutDataID(*id)
	}
	return wrdc
}

// SetWorkoutData sets the "workout_data" edge to the WorkoutData entity.
func (wrdc *WorkoutRouteDataCreate) SetWorkoutData(w *WorkoutData) *WorkoutRouteDataCreate {
	return wrdc.SetWorkoutDataID(w.ID)
}

// Mutation returns the WorkoutRouteDataMutation object of the builder.
func (wrdc *WorkoutRouteDataCreate) Mutation() *WorkoutRouteDataMutation {
	return wrdc.mutation
}

// Save creates the WorkoutRouteData in the database.
func (wrdc *WorkoutRouteDataCreate) Save(ctx context.Context) (*WorkoutRouteData, error) {
	wrdc.defaults()
	return withHooks(ctx, wrdc.sqlSave, wrdc.mutation, wrdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wrdc *WorkoutRouteDataCreate) SaveX(ctx context.Context) *WorkoutRouteData {
	v, err := wrdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wrdc *WorkoutRouteDataCreate) Exec(ctx context.Context) error {
	_, err := wrdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrdc *WorkoutRouteDataCreate) ExecX(ctx context.Context) {
	if err := wrdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wrdc *WorkoutRouteDataCreate) defaults() {
	if _, ok := wrdc.mutation.CreatedAt(); !ok {
		v := workoutroutedata.DefaultCreatedAt()
		wrdc.mutation.SetCreatedAt(v)
	}
	if _, ok := wrdc.mutation.UpdatedAt(); !ok {
		v := workoutroutedata.DefaultUpdatedAt()
		wrdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wrdc.mutation.ID(); !ok {
		v := workoutroutedata.DefaultID()
		wrdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wrdc *WorkoutRouteDataCreate) check() error {
	if _, ok := wrdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WorkoutRouteData.created_at"`)}
	}
	if _, ok := wrdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WorkoutRouteData.updated_at"`)}
	}
	if _, ok := wrdc.mutation.WorkoutID(); !ok {
		return &ValidationError{Name: "workout", err: errors.New(`ent: missing required edge "WorkoutRouteData.workout"`)}
	}
	return nil
}

func (wrdc *WorkoutRouteDataCreate) sqlSave(ctx context.Context) (*WorkoutRouteData, error) {
	if err := wrdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wrdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wrdc.driver, _spec); err != nil {
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
	wrdc.mutation.id = &_node.ID
	wrdc.mutation.done = true
	return _node, nil
}

func (wrdc *WorkoutRouteDataCreate) createSpec() (*WorkoutRouteData, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkoutRouteData{config: wrdc.config}
		_spec = sqlgraph.NewCreateSpec(workoutroutedata.Table, sqlgraph.NewFieldSpec(workoutroutedata.FieldID, field.TypeUUID))
	)
	if id, ok := wrdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wrdc.mutation.CreatedAt(); ok {
		_spec.SetField(workoutroutedata.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wrdc.mutation.UpdatedAt(); ok {
		_spec.SetField(workoutroutedata.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := wrdc.mutation.WorkoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutroutedata.WorkoutTable,
			Columns: []string{workoutroutedata.WorkoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workout_route_data_workout = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wrdc.mutation.WorkoutDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workoutroutedata.WorkoutDataTable,
			Columns: []string{workoutroutedata.WorkoutDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workoutdata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workout_data_workout_route_data = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkoutRouteDataCreateBulk is the builder for creating many WorkoutRouteData entities in bulk.
type WorkoutRouteDataCreateBulk struct {
	config
	builders []*WorkoutRouteDataCreate
}

// Save creates the WorkoutRouteData entities in the database.
func (wrdcb *WorkoutRouteDataCreateBulk) Save(ctx context.Context) ([]*WorkoutRouteData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wrdcb.builders))
	nodes := make([]*WorkoutRouteData, len(wrdcb.builders))
	mutators := make([]Mutator, len(wrdcb.builders))
	for i := range wrdcb.builders {
		func(i int, root context.Context) {
			builder := wrdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkoutRouteDataMutation)
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
					_, err = mutators[i+1].Mutate(root, wrdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wrdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wrdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wrdcb *WorkoutRouteDataCreateBulk) SaveX(ctx context.Context) []*WorkoutRouteData {
	v, err := wrdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wrdcb *WorkoutRouteDataCreateBulk) Exec(ctx context.Context) error {
	_, err := wrdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrdcb *WorkoutRouteDataCreateBulk) ExecX(ctx context.Context) {
	if err := wrdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
