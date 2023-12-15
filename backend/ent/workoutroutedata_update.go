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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutroutedata"
)

// WorkoutRouteDataUpdate is the builder for updating WorkoutRouteData entities.
type WorkoutRouteDataUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkoutRouteDataMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the WorkoutRouteDataUpdate builder.
func (wrdu *WorkoutRouteDataUpdate) Where(ps ...predicate.WorkoutRouteData) *WorkoutRouteDataUpdate {
	wrdu.mutation.Where(ps...)
	return wrdu
}

// SetUpdatedAt sets the "updated_at" field.
func (wrdu *WorkoutRouteDataUpdate) SetUpdatedAt(t time.Time) *WorkoutRouteDataUpdate {
	wrdu.mutation.SetUpdatedAt(t)
	return wrdu
}

// SetWorkoutID sets the "workout" edge to the Workout entity by ID.
func (wrdu *WorkoutRouteDataUpdate) SetWorkoutID(id uuid.UUID) *WorkoutRouteDataUpdate {
	wrdu.mutation.SetWorkoutID(id)
	return wrdu
}

// SetWorkout sets the "workout" edge to the Workout entity.
func (wrdu *WorkoutRouteDataUpdate) SetWorkout(w *Workout) *WorkoutRouteDataUpdate {
	return wrdu.SetWorkoutID(w.ID)
}

// SetWorkoutDataID sets the "workout_data" edge to the WorkoutData entity by ID.
func (wrdu *WorkoutRouteDataUpdate) SetWorkoutDataID(id uuid.UUID) *WorkoutRouteDataUpdate {
	wrdu.mutation.SetWorkoutDataID(id)
	return wrdu
}

// SetNillableWorkoutDataID sets the "workout_data" edge to the WorkoutData entity by ID if the given value is not nil.
func (wrdu *WorkoutRouteDataUpdate) SetNillableWorkoutDataID(id *uuid.UUID) *WorkoutRouteDataUpdate {
	if id != nil {
		wrdu = wrdu.SetWorkoutDataID(*id)
	}
	return wrdu
}

// SetWorkoutData sets the "workout_data" edge to the WorkoutData entity.
func (wrdu *WorkoutRouteDataUpdate) SetWorkoutData(w *WorkoutData) *WorkoutRouteDataUpdate {
	return wrdu.SetWorkoutDataID(w.ID)
}

// Mutation returns the WorkoutRouteDataMutation object of the builder.
func (wrdu *WorkoutRouteDataUpdate) Mutation() *WorkoutRouteDataMutation {
	return wrdu.mutation
}

// ClearWorkout clears the "workout" edge to the Workout entity.
func (wrdu *WorkoutRouteDataUpdate) ClearWorkout() *WorkoutRouteDataUpdate {
	wrdu.mutation.ClearWorkout()
	return wrdu
}

// ClearWorkoutData clears the "workout_data" edge to the WorkoutData entity.
func (wrdu *WorkoutRouteDataUpdate) ClearWorkoutData() *WorkoutRouteDataUpdate {
	wrdu.mutation.ClearWorkoutData()
	return wrdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wrdu *WorkoutRouteDataUpdate) Save(ctx context.Context) (int, error) {
	wrdu.defaults()
	return withHooks(ctx, wrdu.sqlSave, wrdu.mutation, wrdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wrdu *WorkoutRouteDataUpdate) SaveX(ctx context.Context) int {
	affected, err := wrdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wrdu *WorkoutRouteDataUpdate) Exec(ctx context.Context) error {
	_, err := wrdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrdu *WorkoutRouteDataUpdate) ExecX(ctx context.Context) {
	if err := wrdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wrdu *WorkoutRouteDataUpdate) defaults() {
	if _, ok := wrdu.mutation.UpdatedAt(); !ok {
		v := workoutroutedata.UpdateDefaultUpdatedAt()
		wrdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wrdu *WorkoutRouteDataUpdate) check() error {
	if _, ok := wrdu.mutation.WorkoutID(); wrdu.mutation.WorkoutCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutRouteData.workout"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wrdu *WorkoutRouteDataUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkoutRouteDataUpdate {
	wrdu.modifiers = append(wrdu.modifiers, modifiers...)
	return wrdu
}

func (wrdu *WorkoutRouteDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wrdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workoutroutedata.Table, workoutroutedata.Columns, sqlgraph.NewFieldSpec(workoutroutedata.FieldID, field.TypeUUID))
	if ps := wrdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wrdu.mutation.UpdatedAt(); ok {
		_spec.SetField(workoutroutedata.FieldUpdatedAt, field.TypeTime, value)
	}
	if wrdu.mutation.WorkoutCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wrdu.mutation.WorkoutIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wrdu.mutation.WorkoutDataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wrdu.mutation.WorkoutDataIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(wrdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wrdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workoutroutedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wrdu.mutation.done = true
	return n, nil
}

// WorkoutRouteDataUpdateOne is the builder for updating a single WorkoutRouteData entity.
type WorkoutRouteDataUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkoutRouteDataMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (wrduo *WorkoutRouteDataUpdateOne) SetUpdatedAt(t time.Time) *WorkoutRouteDataUpdateOne {
	wrduo.mutation.SetUpdatedAt(t)
	return wrduo
}

// SetWorkoutID sets the "workout" edge to the Workout entity by ID.
func (wrduo *WorkoutRouteDataUpdateOne) SetWorkoutID(id uuid.UUID) *WorkoutRouteDataUpdateOne {
	wrduo.mutation.SetWorkoutID(id)
	return wrduo
}

// SetWorkout sets the "workout" edge to the Workout entity.
func (wrduo *WorkoutRouteDataUpdateOne) SetWorkout(w *Workout) *WorkoutRouteDataUpdateOne {
	return wrduo.SetWorkoutID(w.ID)
}

// SetWorkoutDataID sets the "workout_data" edge to the WorkoutData entity by ID.
func (wrduo *WorkoutRouteDataUpdateOne) SetWorkoutDataID(id uuid.UUID) *WorkoutRouteDataUpdateOne {
	wrduo.mutation.SetWorkoutDataID(id)
	return wrduo
}

// SetNillableWorkoutDataID sets the "workout_data" edge to the WorkoutData entity by ID if the given value is not nil.
func (wrduo *WorkoutRouteDataUpdateOne) SetNillableWorkoutDataID(id *uuid.UUID) *WorkoutRouteDataUpdateOne {
	if id != nil {
		wrduo = wrduo.SetWorkoutDataID(*id)
	}
	return wrduo
}

// SetWorkoutData sets the "workout_data" edge to the WorkoutData entity.
func (wrduo *WorkoutRouteDataUpdateOne) SetWorkoutData(w *WorkoutData) *WorkoutRouteDataUpdateOne {
	return wrduo.SetWorkoutDataID(w.ID)
}

// Mutation returns the WorkoutRouteDataMutation object of the builder.
func (wrduo *WorkoutRouteDataUpdateOne) Mutation() *WorkoutRouteDataMutation {
	return wrduo.mutation
}

// ClearWorkout clears the "workout" edge to the Workout entity.
func (wrduo *WorkoutRouteDataUpdateOne) ClearWorkout() *WorkoutRouteDataUpdateOne {
	wrduo.mutation.ClearWorkout()
	return wrduo
}

// ClearWorkoutData clears the "workout_data" edge to the WorkoutData entity.
func (wrduo *WorkoutRouteDataUpdateOne) ClearWorkoutData() *WorkoutRouteDataUpdateOne {
	wrduo.mutation.ClearWorkoutData()
	return wrduo
}

// Where appends a list predicates to the WorkoutRouteDataUpdate builder.
func (wrduo *WorkoutRouteDataUpdateOne) Where(ps ...predicate.WorkoutRouteData) *WorkoutRouteDataUpdateOne {
	wrduo.mutation.Where(ps...)
	return wrduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wrduo *WorkoutRouteDataUpdateOne) Select(field string, fields ...string) *WorkoutRouteDataUpdateOne {
	wrduo.fields = append([]string{field}, fields...)
	return wrduo
}

// Save executes the query and returns the updated WorkoutRouteData entity.
func (wrduo *WorkoutRouteDataUpdateOne) Save(ctx context.Context) (*WorkoutRouteData, error) {
	wrduo.defaults()
	return withHooks(ctx, wrduo.sqlSave, wrduo.mutation, wrduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wrduo *WorkoutRouteDataUpdateOne) SaveX(ctx context.Context) *WorkoutRouteData {
	node, err := wrduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wrduo *WorkoutRouteDataUpdateOne) Exec(ctx context.Context) error {
	_, err := wrduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrduo *WorkoutRouteDataUpdateOne) ExecX(ctx context.Context) {
	if err := wrduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wrduo *WorkoutRouteDataUpdateOne) defaults() {
	if _, ok := wrduo.mutation.UpdatedAt(); !ok {
		v := workoutroutedata.UpdateDefaultUpdatedAt()
		wrduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wrduo *WorkoutRouteDataUpdateOne) check() error {
	if _, ok := wrduo.mutation.WorkoutID(); wrduo.mutation.WorkoutCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutRouteData.workout"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wrduo *WorkoutRouteDataUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkoutRouteDataUpdateOne {
	wrduo.modifiers = append(wrduo.modifiers, modifiers...)
	return wrduo
}

func (wrduo *WorkoutRouteDataUpdateOne) sqlSave(ctx context.Context) (_node *WorkoutRouteData, err error) {
	if err := wrduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workoutroutedata.Table, workoutroutedata.Columns, sqlgraph.NewFieldSpec(workoutroutedata.FieldID, field.TypeUUID))
	id, ok := wrduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkoutRouteData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wrduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workoutroutedata.FieldID)
		for _, f := range fields {
			if !workoutroutedata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workoutroutedata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wrduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wrduo.mutation.UpdatedAt(); ok {
		_spec.SetField(workoutroutedata.FieldUpdatedAt, field.TypeTime, value)
	}
	if wrduo.mutation.WorkoutCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wrduo.mutation.WorkoutIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wrduo.mutation.WorkoutDataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wrduo.mutation.WorkoutDataIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(wrduo.modifiers...)
	_node = &WorkoutRouteData{config: wrduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wrduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workoutroutedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wrduo.mutation.done = true
	return _node, nil
}