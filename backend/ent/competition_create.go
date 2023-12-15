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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competitionresult"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
)

// CompetitionCreate is the builder for creating a Competition entity.
type CompetitionCreate struct {
	config
	mutation *CompetitionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CompetitionCreate) SetCreatedAt(t time.Time) *CompetitionCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CompetitionCreate) SetNillableCreatedAt(t *time.Time) *CompetitionCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CompetitionCreate) SetUpdatedAt(t time.Time) *CompetitionCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CompetitionCreate) SetNillableUpdatedAt(t *time.Time) *CompetitionCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetPublic sets the "public" field.
func (cc *CompetitionCreate) SetPublic(b bool) *CompetitionCreate {
	cc.mutation.SetPublic(b)
	return cc
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (cc *CompetitionCreate) SetNillablePublic(b *bool) *CompetitionCreate {
	if b != nil {
		cc.SetPublic(*b)
	}
	return cc
}

// SetTitle sets the "title" field.
func (cc *CompetitionCreate) SetTitle(s string) *CompetitionCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CompetitionCreate) SetDescription(s string) *CompetitionCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetStart sets the "start" field.
func (cc *CompetitionCreate) SetStart(t time.Time) *CompetitionCreate {
	cc.mutation.SetStart(t)
	return cc
}

// SetEnd sets the "end" field.
func (cc *CompetitionCreate) SetEnd(t time.Time) *CompetitionCreate {
	cc.mutation.SetEnd(t)
	return cc
}

// SetScheduled sets the "scheduled" field.
func (cc *CompetitionCreate) SetScheduled(b bool) *CompetitionCreate {
	cc.mutation.SetScheduled(b)
	return cc
}

// SetNillableScheduled sets the "scheduled" field if the given value is not nil.
func (cc *CompetitionCreate) SetNillableScheduled(b *bool) *CompetitionCreate {
	if b != nil {
		cc.SetScheduled(*b)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CompetitionCreate) SetStatus(s string) *CompetitionCreate {
	cc.mutation.SetStatus(s)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CompetitionCreate) SetNillableStatus(s *string) *CompetitionCreate {
	if s != nil {
		cc.SetStatus(*s)
	}
	return cc
}

// SetParticipantTypes sets the "participant_types" field.
func (cc *CompetitionCreate) SetParticipantTypes(s []string) *CompetitionCreate {
	cc.mutation.SetParticipantTypes(s)
	return cc
}

// SetWorkoutTypes sets the "workout_types" field.
func (cc *CompetitionCreate) SetWorkoutTypes(s []string) *CompetitionCreate {
	cc.mutation.SetWorkoutTypes(s)
	return cc
}

// SetType sets the "type" field.
func (cc *CompetitionCreate) SetType(s string) *CompetitionCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cc *CompetitionCreate) SetNillableType(s *string) *CompetitionCreate {
	if s != nil {
		cc.SetType(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CompetitionCreate) SetID(u uuid.UUID) *CompetitionCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CompetitionCreate) SetNillableID(u *uuid.UUID) *CompetitionCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// SetOwnerID sets the "owner" edge to the Profile entity by ID.
func (cc *CompetitionCreate) SetOwnerID(id uuid.UUID) *CompetitionCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetOwner sets the "owner" edge to the Profile entity.
func (cc *CompetitionCreate) SetOwner(p *Profile) *CompetitionCreate {
	return cc.SetOwnerID(p.ID)
}

// AddParticipantIDs adds the "participants" edge to the Profile entity by IDs.
func (cc *CompetitionCreate) AddParticipantIDs(ids ...uuid.UUID) *CompetitionCreate {
	cc.mutation.AddParticipantIDs(ids...)
	return cc
}

// AddParticipants adds the "participants" edges to the Profile entity.
func (cc *CompetitionCreate) AddParticipants(p ...*Profile) *CompetitionCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddParticipantIDs(ids...)
}

// AddWorkoutIDs adds the "workouts" edge to the Workout entity by IDs.
func (cc *CompetitionCreate) AddWorkoutIDs(ids ...uuid.UUID) *CompetitionCreate {
	cc.mutation.AddWorkoutIDs(ids...)
	return cc
}

// AddWorkouts adds the "workouts" edges to the Workout entity.
func (cc *CompetitionCreate) AddWorkouts(w ...*Workout) *CompetitionCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cc.AddWorkoutIDs(ids...)
}

// AddWorkoutDatumIDs adds the "workout_data" edge to the WorkoutData entity by IDs.
func (cc *CompetitionCreate) AddWorkoutDatumIDs(ids ...uuid.UUID) *CompetitionCreate {
	cc.mutation.AddWorkoutDatumIDs(ids...)
	return cc
}

// AddWorkoutData adds the "workout_data" edges to the WorkoutData entity.
func (cc *CompetitionCreate) AddWorkoutData(w ...*WorkoutData) *CompetitionCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cc.AddWorkoutDatumIDs(ids...)
}

// AddResultIDs adds the "results" edge to the CompetitionResult entity by IDs.
func (cc *CompetitionCreate) AddResultIDs(ids ...uuid.UUID) *CompetitionCreate {
	cc.mutation.AddResultIDs(ids...)
	return cc
}

// AddResults adds the "results" edges to the CompetitionResult entity.
func (cc *CompetitionCreate) AddResults(c ...*CompetitionResult) *CompetitionCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddResultIDs(ids...)
}

// Mutation returns the CompetitionMutation object of the builder.
func (cc *CompetitionCreate) Mutation() *CompetitionMutation {
	return cc.mutation
}

// Save creates the Competition in the database.
func (cc *CompetitionCreate) Save(ctx context.Context) (*Competition, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompetitionCreate) SaveX(ctx context.Context) *Competition {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CompetitionCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CompetitionCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CompetitionCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if competition.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized competition.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := competition.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if competition.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized competition.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := competition.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Public(); !ok {
		v := competition.DefaultPublic
		cc.mutation.SetPublic(v)
	}
	if _, ok := cc.mutation.Scheduled(); !ok {
		v := competition.DefaultScheduled
		cc.mutation.SetScheduled(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := competition.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.GetType(); !ok {
		v := competition.DefaultType
		cc.mutation.SetType(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if competition.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized competition.DefaultID (forgotten import ent/runtime?)")
		}
		v := competition.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CompetitionCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Competition.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Competition.updated_at"`)}
	}
	if _, ok := cc.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "Competition.public"`)}
	}
	if _, ok := cc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Competition.title"`)}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Competition.description"`)}
	}
	if _, ok := cc.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "Competition.start"`)}
	}
	if _, ok := cc.mutation.End(); !ok {
		return &ValidationError{Name: "end", err: errors.New(`ent: missing required field "Competition.end"`)}
	}
	if _, ok := cc.mutation.Scheduled(); !ok {
		return &ValidationError{Name: "scheduled", err: errors.New(`ent: missing required field "Competition.scheduled"`)}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Competition.status"`)}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Competition.type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := competition.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Competition.type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Competition.owner"`)}
	}
	return nil
}

func (cc *CompetitionCreate) sqlSave(ctx context.Context) (*Competition, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CompetitionCreate) createSpec() (*Competition, *sqlgraph.CreateSpec) {
	var (
		_node = &Competition{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(competition.Table, sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(competition.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(competition.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Public(); ok {
		_spec.SetField(competition.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if value, ok := cc.mutation.Title(); ok {
		_spec.SetField(competition.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(competition.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.Start(); ok {
		_spec.SetField(competition.FieldStart, field.TypeTime, value)
		_node.Start = value
	}
	if value, ok := cc.mutation.End(); ok {
		_spec.SetField(competition.FieldEnd, field.TypeTime, value)
		_node.End = value
	}
	if value, ok := cc.mutation.Scheduled(); ok {
		_spec.SetField(competition.FieldScheduled, field.TypeBool, value)
		_node.Scheduled = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(competition.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.ParticipantTypes(); ok {
		_spec.SetField(competition.FieldParticipantTypes, field.TypeJSON, value)
		_node.ParticipantTypes = value
	}
	if value, ok := cc.mutation.WorkoutTypes(); ok {
		_spec.SetField(competition.FieldWorkoutTypes, field.TypeJSON, value)
		_node.WorkoutTypes = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(competition.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   competition.OwnerTable,
			Columns: []string{competition.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.competition_owner = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   competition.ParticipantsTable,
			Columns: competition.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.WorkoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   competition.WorkoutsTable,
			Columns: []string{competition.WorkoutsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.WorkoutDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   competition.WorkoutDataTable,
			Columns: []string{competition.WorkoutDataColumn},
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
	if nodes := cc.mutation.ResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   competition.ResultsTable,
			Columns: []string{competition.ResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competitionresult.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompetitionCreateBulk is the builder for creating many Competition entities in bulk.
type CompetitionCreateBulk struct {
	config
	builders []*CompetitionCreate
}

// Save creates the Competition entities in the database.
func (ccb *CompetitionCreateBulk) Save(ctx context.Context) ([]*Competition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Competition, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompetitionMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CompetitionCreateBulk) SaveX(ctx context.Context) []*Competition {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CompetitionCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CompetitionCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}