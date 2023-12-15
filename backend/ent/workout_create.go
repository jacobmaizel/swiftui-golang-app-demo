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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/invite"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutroutedata"
)

// WorkoutCreate is the builder for creating a Workout entity.
type WorkoutCreate struct {
	config
	mutation *WorkoutMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkoutCreate) SetCreatedAt(t time.Time) *WorkoutCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableCreatedAt(t *time.Time) *WorkoutCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WorkoutCreate) SetUpdatedAt(t time.Time) *WorkoutCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableUpdatedAt(t *time.Time) *WorkoutCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetHealthkitWorkoutActivityType sets the "healthkit_workout_activity_type" field.
func (wc *WorkoutCreate) SetHealthkitWorkoutActivityType(s string) *WorkoutCreate {
	wc.mutation.SetHealthkitWorkoutActivityType(s)
	return wc
}

// SetID sets the "id" field.
func (wc *WorkoutCreate) SetID(u uuid.UUID) *WorkoutCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableID(u *uuid.UUID) *WorkoutCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// AddInviteIDs adds the "invite" edge to the Invite entity by IDs.
func (wc *WorkoutCreate) AddInviteIDs(ids ...uuid.UUID) *WorkoutCreate {
	wc.mutation.AddInviteIDs(ids...)
	return wc
}

// AddInvite adds the "invite" edges to the Invite entity.
func (wc *WorkoutCreate) AddInvite(i ...*Invite) *WorkoutCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wc.AddInviteIDs(ids...)
}

// SetLeaderID sets the "leader" edge to the Profile entity by ID.
func (wc *WorkoutCreate) SetLeaderID(id uuid.UUID) *WorkoutCreate {
	wc.mutation.SetLeaderID(id)
	return wc
}

// SetNillableLeaderID sets the "leader" edge to the Profile entity by ID if the given value is not nil.
func (wc *WorkoutCreate) SetNillableLeaderID(id *uuid.UUID) *WorkoutCreate {
	if id != nil {
		wc = wc.SetLeaderID(*id)
	}
	return wc
}

// SetLeader sets the "leader" edge to the Profile entity.
func (wc *WorkoutCreate) SetLeader(p *Profile) *WorkoutCreate {
	return wc.SetLeaderID(p.ID)
}

// SetCompetitionID sets the "competition" edge to the Competition entity by ID.
func (wc *WorkoutCreate) SetCompetitionID(id uuid.UUID) *WorkoutCreate {
	wc.mutation.SetCompetitionID(id)
	return wc
}

// SetNillableCompetitionID sets the "competition" edge to the Competition entity by ID if the given value is not nil.
func (wc *WorkoutCreate) SetNillableCompetitionID(id *uuid.UUID) *WorkoutCreate {
	if id != nil {
		wc = wc.SetCompetitionID(*id)
	}
	return wc
}

// SetCompetition sets the "competition" edge to the Competition entity.
func (wc *WorkoutCreate) SetCompetition(c *Competition) *WorkoutCreate {
	return wc.SetCompetitionID(c.ID)
}

// AddWorkoutDatumIDs adds the "workout_data" edge to the WorkoutData entity by IDs.
func (wc *WorkoutCreate) AddWorkoutDatumIDs(ids ...uuid.UUID) *WorkoutCreate {
	wc.mutation.AddWorkoutDatumIDs(ids...)
	return wc
}

// AddWorkoutData adds the "workout_data" edges to the WorkoutData entity.
func (wc *WorkoutCreate) AddWorkoutData(w ...*WorkoutData) *WorkoutCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddWorkoutDatumIDs(ids...)
}

// AddWorkoutRouteDatumIDs adds the "workout_route_data" edge to the WorkoutRouteData entity by IDs.
func (wc *WorkoutCreate) AddWorkoutRouteDatumIDs(ids ...uuid.UUID) *WorkoutCreate {
	wc.mutation.AddWorkoutRouteDatumIDs(ids...)
	return wc
}

// AddWorkoutRouteData adds the "workout_route_data" edges to the WorkoutRouteData entity.
func (wc *WorkoutCreate) AddWorkoutRouteData(w ...*WorkoutRouteData) *WorkoutCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddWorkoutRouteDatumIDs(ids...)
}

// Mutation returns the WorkoutMutation object of the builder.
func (wc *WorkoutCreate) Mutation() *WorkoutMutation {
	return wc.mutation
}

// Save creates the Workout in the database.
func (wc *WorkoutCreate) Save(ctx context.Context) (*Workout, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkoutCreate) SaveX(ctx context.Context) *Workout {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkoutCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkoutCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkoutCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := workout.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := workout.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := workout.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkoutCreate) check() error {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Workout.created_at"`)}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Workout.updated_at"`)}
	}
	if _, ok := wc.mutation.HealthkitWorkoutActivityType(); !ok {
		return &ValidationError{Name: "healthkit_workout_activity_type", err: errors.New(`ent: missing required field "Workout.healthkit_workout_activity_type"`)}
	}
	if v, ok := wc.mutation.HealthkitWorkoutActivityType(); ok {
		if err := workout.HealthkitWorkoutActivityTypeValidator(v); err != nil {
			return &ValidationError{Name: "healthkit_workout_activity_type", err: fmt.Errorf(`ent: validator failed for field "Workout.healthkit_workout_activity_type": %w`, err)}
		}
	}
	return nil
}

func (wc *WorkoutCreate) sqlSave(ctx context.Context) (*Workout, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
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
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkoutCreate) createSpec() (*Workout, *sqlgraph.CreateSpec) {
	var (
		_node = &Workout{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(workout.Table, sqlgraph.NewFieldSpec(workout.FieldID, field.TypeUUID))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(workout.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.SetField(workout.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wc.mutation.HealthkitWorkoutActivityType(); ok {
		_spec.SetField(workout.FieldHealthkitWorkoutActivityType, field.TypeString, value)
		_node.HealthkitWorkoutActivityType = value
	}
	if nodes := wc.mutation.InviteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workout.InviteTable,
			Columns: []string{workout.InviteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.LeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workout.LeaderTable,
			Columns: []string{workout.LeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workout_leader = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.CompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workout.CompetitionTable,
			Columns: []string{workout.CompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workout_competition = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.WorkoutDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workout.WorkoutDataTable,
			Columns: []string{workout.WorkoutDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workoutdata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.WorkoutRouteDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workout.WorkoutRouteDataTable,
			Columns: []string{workout.WorkoutRouteDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workoutroutedata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkoutCreateBulk is the builder for creating many Workout entities in bulk.
type WorkoutCreateBulk struct {
	config
	builders []*WorkoutCreate
}

// Save creates the Workout entities in the database.
func (wcb *WorkoutCreateBulk) Save(ctx context.Context) ([]*Workout, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workout, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkoutMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkoutCreateBulk) SaveX(ctx context.Context) []*Workout {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkoutCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkoutCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}