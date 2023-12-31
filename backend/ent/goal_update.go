// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/goal"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/squad"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// GoalUpdate is the builder for updating Goal entities.
type GoalUpdate struct {
	config
	hooks     []Hook
	mutation  *GoalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GoalUpdate builder.
func (gu *GoalUpdate) Where(ps ...predicate.Goal) *GoalUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GoalUpdate) SetUpdatedAt(t time.Time) *GoalUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetType sets the "type" field.
func (gu *GoalUpdate) SetType(s string) *GoalUpdate {
	gu.mutation.SetType(s)
	return gu
}

// SetStart sets the "start" field.
func (gu *GoalUpdate) SetStart(t time.Time) *GoalUpdate {
	gu.mutation.SetStart(t)
	return gu
}

// SetEnd sets the "end" field.
func (gu *GoalUpdate) SetEnd(t time.Time) *GoalUpdate {
	gu.mutation.SetEnd(t)
	return gu
}

// SetHealthkitWorkoutActivityType sets the "healthkit_workout_activity_type" field.
func (gu *GoalUpdate) SetHealthkitWorkoutActivityType(s string) *GoalUpdate {
	gu.mutation.SetHealthkitWorkoutActivityType(s)
	return gu
}

// SetNillableHealthkitWorkoutActivityType sets the "healthkit_workout_activity_type" field if the given value is not nil.
func (gu *GoalUpdate) SetNillableHealthkitWorkoutActivityType(s *string) *GoalUpdate {
	if s != nil {
		gu.SetHealthkitWorkoutActivityType(*s)
	}
	return gu
}

// ClearHealthkitWorkoutActivityType clears the value of the "healthkit_workout_activity_type" field.
func (gu *GoalUpdate) ClearHealthkitWorkoutActivityType() *GoalUpdate {
	gu.mutation.ClearHealthkitWorkoutActivityType()
	return gu
}

// SetAction sets the "action" field.
func (gu *GoalUpdate) SetAction(s string) *GoalUpdate {
	gu.mutation.SetAction(s)
	return gu
}

// SetValue sets the "value" field.
func (gu *GoalUpdate) SetValue(s string) *GoalUpdate {
	gu.mutation.SetValue(s)
	return gu
}

// SetUnit sets the "unit" field.
func (gu *GoalUpdate) SetUnit(s string) *GoalUpdate {
	gu.mutation.SetUnit(s)
	return gu
}

// SetValueAggregationType sets the "value_aggregation_type" field.
func (gu *GoalUpdate) SetValueAggregationType(s string) *GoalUpdate {
	gu.mutation.SetValueAggregationType(s)
	return gu
}

// SetTimeInterval sets the "time_interval" field.
func (gu *GoalUpdate) SetTimeInterval(s string) *GoalUpdate {
	gu.mutation.SetTimeInterval(s)
	return gu
}

// SetStatus sets the "status" field.
func (gu *GoalUpdate) SetStatus(s string) *GoalUpdate {
	gu.mutation.SetStatus(s)
	return gu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (gu *GoalUpdate) SetNillableStatus(s *string) *GoalUpdate {
	if s != nil {
		gu.SetStatus(*s)
	}
	return gu
}

// SetCurrentTotalValue sets the "current_total_value" field.
func (gu *GoalUpdate) SetCurrentTotalValue(s string) *GoalUpdate {
	gu.mutation.SetCurrentTotalValue(s)
	return gu
}

// SetNillableCurrentTotalValue sets the "current_total_value" field if the given value is not nil.
func (gu *GoalUpdate) SetNillableCurrentTotalValue(s *string) *GoalUpdate {
	if s != nil {
		gu.SetCurrentTotalValue(*s)
	}
	return gu
}

// ClearCurrentTotalValue clears the value of the "current_total_value" field.
func (gu *GoalUpdate) ClearCurrentTotalValue() *GoalUpdate {
	gu.mutation.ClearCurrentTotalValue()
	return gu
}

// SetPerWorkoutData sets the "per_workout_data" field.
func (gu *GoalUpdate) SetPerWorkoutData(swgde []shared.PerWorkoutGoalDataEntry) *GoalUpdate {
	gu.mutation.SetPerWorkoutData(swgde)
	return gu
}

// AppendPerWorkoutData appends swgde to the "per_workout_data" field.
func (gu *GoalUpdate) AppendPerWorkoutData(swgde []shared.PerWorkoutGoalDataEntry) *GoalUpdate {
	gu.mutation.AppendPerWorkoutData(swgde)
	return gu
}

// ClearPerWorkoutData clears the value of the "per_workout_data" field.
func (gu *GoalUpdate) ClearPerWorkoutData() *GoalUpdate {
	gu.mutation.ClearPerWorkoutData()
	return gu
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (gu *GoalUpdate) SetProfileID(id uuid.UUID) *GoalUpdate {
	gu.mutation.SetProfileID(id)
	return gu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (gu *GoalUpdate) SetProfile(p *Profile) *GoalUpdate {
	return gu.SetProfileID(p.ID)
}

// SetCompetitionID sets the "competition" edge to the Competition entity by ID.
func (gu *GoalUpdate) SetCompetitionID(id uuid.UUID) *GoalUpdate {
	gu.mutation.SetCompetitionID(id)
	return gu
}

// SetNillableCompetitionID sets the "competition" edge to the Competition entity by ID if the given value is not nil.
func (gu *GoalUpdate) SetNillableCompetitionID(id *uuid.UUID) *GoalUpdate {
	if id != nil {
		gu = gu.SetCompetitionID(*id)
	}
	return gu
}

// SetCompetition sets the "competition" edge to the Competition entity.
func (gu *GoalUpdate) SetCompetition(c *Competition) *GoalUpdate {
	return gu.SetCompetitionID(c.ID)
}

// SetSquadID sets the "squad" edge to the Squad entity by ID.
func (gu *GoalUpdate) SetSquadID(id uuid.UUID) *GoalUpdate {
	gu.mutation.SetSquadID(id)
	return gu
}

// SetNillableSquadID sets the "squad" edge to the Squad entity by ID if the given value is not nil.
func (gu *GoalUpdate) SetNillableSquadID(id *uuid.UUID) *GoalUpdate {
	if id != nil {
		gu = gu.SetSquadID(*id)
	}
	return gu
}

// SetSquad sets the "squad" edge to the Squad entity.
func (gu *GoalUpdate) SetSquad(s *Squad) *GoalUpdate {
	return gu.SetSquadID(s.ID)
}

// Mutation returns the GoalMutation object of the builder.
func (gu *GoalUpdate) Mutation() *GoalMutation {
	return gu.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (gu *GoalUpdate) ClearProfile() *GoalUpdate {
	gu.mutation.ClearProfile()
	return gu
}

// ClearCompetition clears the "competition" edge to the Competition entity.
func (gu *GoalUpdate) ClearCompetition() *GoalUpdate {
	gu.mutation.ClearCompetition()
	return gu
}

// ClearSquad clears the "squad" edge to the Squad entity.
func (gu *GoalUpdate) ClearSquad() *GoalUpdate {
	gu.mutation.ClearSquad()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GoalUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GoalUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GoalUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GoalUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GoalUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := goal.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GoalUpdate) check() error {
	if _, ok := gu.mutation.ProfileID(); gu.mutation.ProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Goal.profile"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gu *GoalUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoalUpdate {
	gu.modifiers = append(gu.modifiers, modifiers...)
	return gu
}

func (gu *GoalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(goal.Table, goal.Columns, sqlgraph.NewFieldSpec(goal.FieldID, field.TypeUUID))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(goal.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.GetType(); ok {
		_spec.SetField(goal.FieldType, field.TypeString, value)
	}
	if value, ok := gu.mutation.Start(); ok {
		_spec.SetField(goal.FieldStart, field.TypeTime, value)
	}
	if value, ok := gu.mutation.End(); ok {
		_spec.SetField(goal.FieldEnd, field.TypeTime, value)
	}
	if value, ok := gu.mutation.HealthkitWorkoutActivityType(); ok {
		_spec.SetField(goal.FieldHealthkitWorkoutActivityType, field.TypeString, value)
	}
	if gu.mutation.HealthkitWorkoutActivityTypeCleared() {
		_spec.ClearField(goal.FieldHealthkitWorkoutActivityType, field.TypeString)
	}
	if value, ok := gu.mutation.Action(); ok {
		_spec.SetField(goal.FieldAction, field.TypeString, value)
	}
	if value, ok := gu.mutation.Value(); ok {
		_spec.SetField(goal.FieldValue, field.TypeString, value)
	}
	if value, ok := gu.mutation.Unit(); ok {
		_spec.SetField(goal.FieldUnit, field.TypeString, value)
	}
	if value, ok := gu.mutation.ValueAggregationType(); ok {
		_spec.SetField(goal.FieldValueAggregationType, field.TypeString, value)
	}
	if value, ok := gu.mutation.TimeInterval(); ok {
		_spec.SetField(goal.FieldTimeInterval, field.TypeString, value)
	}
	if value, ok := gu.mutation.Status(); ok {
		_spec.SetField(goal.FieldStatus, field.TypeString, value)
	}
	if value, ok := gu.mutation.CurrentTotalValue(); ok {
		_spec.SetField(goal.FieldCurrentTotalValue, field.TypeString, value)
	}
	if gu.mutation.CurrentTotalValueCleared() {
		_spec.ClearField(goal.FieldCurrentTotalValue, field.TypeString)
	}
	if value, ok := gu.mutation.PerWorkoutData(); ok {
		_spec.SetField(goal.FieldPerWorkoutData, field.TypeJSON, value)
	}
	if value, ok := gu.mutation.AppendedPerWorkoutData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, goal.FieldPerWorkoutData, value)
		})
	}
	if gu.mutation.PerWorkoutDataCleared() {
		_spec.ClearField(goal.FieldPerWorkoutData, field.TypeJSON)
	}
	if gu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.ProfileTable,
			Columns: []string{goal.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.ProfileTable,
			Columns: []string{goal.ProfileColumn},
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
	if gu.mutation.CompetitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.CompetitionTable,
			Columns: []string{goal.CompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.CompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.CompetitionTable,
			Columns: []string{goal.CompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.SquadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.SquadTable,
			Columns: []string{goal.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.SquadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.SquadTable,
			Columns: []string{goal.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(gu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GoalUpdateOne is the builder for updating a single Goal entity.
type GoalUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GoalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GoalUpdateOne) SetUpdatedAt(t time.Time) *GoalUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetType sets the "type" field.
func (guo *GoalUpdateOne) SetType(s string) *GoalUpdateOne {
	guo.mutation.SetType(s)
	return guo
}

// SetStart sets the "start" field.
func (guo *GoalUpdateOne) SetStart(t time.Time) *GoalUpdateOne {
	guo.mutation.SetStart(t)
	return guo
}

// SetEnd sets the "end" field.
func (guo *GoalUpdateOne) SetEnd(t time.Time) *GoalUpdateOne {
	guo.mutation.SetEnd(t)
	return guo
}

// SetHealthkitWorkoutActivityType sets the "healthkit_workout_activity_type" field.
func (guo *GoalUpdateOne) SetHealthkitWorkoutActivityType(s string) *GoalUpdateOne {
	guo.mutation.SetHealthkitWorkoutActivityType(s)
	return guo
}

// SetNillableHealthkitWorkoutActivityType sets the "healthkit_workout_activity_type" field if the given value is not nil.
func (guo *GoalUpdateOne) SetNillableHealthkitWorkoutActivityType(s *string) *GoalUpdateOne {
	if s != nil {
		guo.SetHealthkitWorkoutActivityType(*s)
	}
	return guo
}

// ClearHealthkitWorkoutActivityType clears the value of the "healthkit_workout_activity_type" field.
func (guo *GoalUpdateOne) ClearHealthkitWorkoutActivityType() *GoalUpdateOne {
	guo.mutation.ClearHealthkitWorkoutActivityType()
	return guo
}

// SetAction sets the "action" field.
func (guo *GoalUpdateOne) SetAction(s string) *GoalUpdateOne {
	guo.mutation.SetAction(s)
	return guo
}

// SetValue sets the "value" field.
func (guo *GoalUpdateOne) SetValue(s string) *GoalUpdateOne {
	guo.mutation.SetValue(s)
	return guo
}

// SetUnit sets the "unit" field.
func (guo *GoalUpdateOne) SetUnit(s string) *GoalUpdateOne {
	guo.mutation.SetUnit(s)
	return guo
}

// SetValueAggregationType sets the "value_aggregation_type" field.
func (guo *GoalUpdateOne) SetValueAggregationType(s string) *GoalUpdateOne {
	guo.mutation.SetValueAggregationType(s)
	return guo
}

// SetTimeInterval sets the "time_interval" field.
func (guo *GoalUpdateOne) SetTimeInterval(s string) *GoalUpdateOne {
	guo.mutation.SetTimeInterval(s)
	return guo
}

// SetStatus sets the "status" field.
func (guo *GoalUpdateOne) SetStatus(s string) *GoalUpdateOne {
	guo.mutation.SetStatus(s)
	return guo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (guo *GoalUpdateOne) SetNillableStatus(s *string) *GoalUpdateOne {
	if s != nil {
		guo.SetStatus(*s)
	}
	return guo
}

// SetCurrentTotalValue sets the "current_total_value" field.
func (guo *GoalUpdateOne) SetCurrentTotalValue(s string) *GoalUpdateOne {
	guo.mutation.SetCurrentTotalValue(s)
	return guo
}

// SetNillableCurrentTotalValue sets the "current_total_value" field if the given value is not nil.
func (guo *GoalUpdateOne) SetNillableCurrentTotalValue(s *string) *GoalUpdateOne {
	if s != nil {
		guo.SetCurrentTotalValue(*s)
	}
	return guo
}

// ClearCurrentTotalValue clears the value of the "current_total_value" field.
func (guo *GoalUpdateOne) ClearCurrentTotalValue() *GoalUpdateOne {
	guo.mutation.ClearCurrentTotalValue()
	return guo
}

// SetPerWorkoutData sets the "per_workout_data" field.
func (guo *GoalUpdateOne) SetPerWorkoutData(swgde []shared.PerWorkoutGoalDataEntry) *GoalUpdateOne {
	guo.mutation.SetPerWorkoutData(swgde)
	return guo
}

// AppendPerWorkoutData appends swgde to the "per_workout_data" field.
func (guo *GoalUpdateOne) AppendPerWorkoutData(swgde []shared.PerWorkoutGoalDataEntry) *GoalUpdateOne {
	guo.mutation.AppendPerWorkoutData(swgde)
	return guo
}

// ClearPerWorkoutData clears the value of the "per_workout_data" field.
func (guo *GoalUpdateOne) ClearPerWorkoutData() *GoalUpdateOne {
	guo.mutation.ClearPerWorkoutData()
	return guo
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (guo *GoalUpdateOne) SetProfileID(id uuid.UUID) *GoalUpdateOne {
	guo.mutation.SetProfileID(id)
	return guo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (guo *GoalUpdateOne) SetProfile(p *Profile) *GoalUpdateOne {
	return guo.SetProfileID(p.ID)
}

// SetCompetitionID sets the "competition" edge to the Competition entity by ID.
func (guo *GoalUpdateOne) SetCompetitionID(id uuid.UUID) *GoalUpdateOne {
	guo.mutation.SetCompetitionID(id)
	return guo
}

// SetNillableCompetitionID sets the "competition" edge to the Competition entity by ID if the given value is not nil.
func (guo *GoalUpdateOne) SetNillableCompetitionID(id *uuid.UUID) *GoalUpdateOne {
	if id != nil {
		guo = guo.SetCompetitionID(*id)
	}
	return guo
}

// SetCompetition sets the "competition" edge to the Competition entity.
func (guo *GoalUpdateOne) SetCompetition(c *Competition) *GoalUpdateOne {
	return guo.SetCompetitionID(c.ID)
}

// SetSquadID sets the "squad" edge to the Squad entity by ID.
func (guo *GoalUpdateOne) SetSquadID(id uuid.UUID) *GoalUpdateOne {
	guo.mutation.SetSquadID(id)
	return guo
}

// SetNillableSquadID sets the "squad" edge to the Squad entity by ID if the given value is not nil.
func (guo *GoalUpdateOne) SetNillableSquadID(id *uuid.UUID) *GoalUpdateOne {
	if id != nil {
		guo = guo.SetSquadID(*id)
	}
	return guo
}

// SetSquad sets the "squad" edge to the Squad entity.
func (guo *GoalUpdateOne) SetSquad(s *Squad) *GoalUpdateOne {
	return guo.SetSquadID(s.ID)
}

// Mutation returns the GoalMutation object of the builder.
func (guo *GoalUpdateOne) Mutation() *GoalMutation {
	return guo.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (guo *GoalUpdateOne) ClearProfile() *GoalUpdateOne {
	guo.mutation.ClearProfile()
	return guo
}

// ClearCompetition clears the "competition" edge to the Competition entity.
func (guo *GoalUpdateOne) ClearCompetition() *GoalUpdateOne {
	guo.mutation.ClearCompetition()
	return guo
}

// ClearSquad clears the "squad" edge to the Squad entity.
func (guo *GoalUpdateOne) ClearSquad() *GoalUpdateOne {
	guo.mutation.ClearSquad()
	return guo
}

// Where appends a list predicates to the GoalUpdate builder.
func (guo *GoalUpdateOne) Where(ps ...predicate.Goal) *GoalUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GoalUpdateOne) Select(field string, fields ...string) *GoalUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Goal entity.
func (guo *GoalUpdateOne) Save(ctx context.Context) (*Goal, error) {
	guo.defaults()
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GoalUpdateOne) SaveX(ctx context.Context) *Goal {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GoalUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GoalUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GoalUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := goal.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GoalUpdateOne) check() error {
	if _, ok := guo.mutation.ProfileID(); guo.mutation.ProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Goal.profile"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (guo *GoalUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoalUpdateOne {
	guo.modifiers = append(guo.modifiers, modifiers...)
	return guo
}

func (guo *GoalUpdateOne) sqlSave(ctx context.Context) (_node *Goal, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(goal.Table, goal.Columns, sqlgraph.NewFieldSpec(goal.FieldID, field.TypeUUID))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Goal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goal.FieldID)
		for _, f := range fields {
			if !goal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(goal.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.GetType(); ok {
		_spec.SetField(goal.FieldType, field.TypeString, value)
	}
	if value, ok := guo.mutation.Start(); ok {
		_spec.SetField(goal.FieldStart, field.TypeTime, value)
	}
	if value, ok := guo.mutation.End(); ok {
		_spec.SetField(goal.FieldEnd, field.TypeTime, value)
	}
	if value, ok := guo.mutation.HealthkitWorkoutActivityType(); ok {
		_spec.SetField(goal.FieldHealthkitWorkoutActivityType, field.TypeString, value)
	}
	if guo.mutation.HealthkitWorkoutActivityTypeCleared() {
		_spec.ClearField(goal.FieldHealthkitWorkoutActivityType, field.TypeString)
	}
	if value, ok := guo.mutation.Action(); ok {
		_spec.SetField(goal.FieldAction, field.TypeString, value)
	}
	if value, ok := guo.mutation.Value(); ok {
		_spec.SetField(goal.FieldValue, field.TypeString, value)
	}
	if value, ok := guo.mutation.Unit(); ok {
		_spec.SetField(goal.FieldUnit, field.TypeString, value)
	}
	if value, ok := guo.mutation.ValueAggregationType(); ok {
		_spec.SetField(goal.FieldValueAggregationType, field.TypeString, value)
	}
	if value, ok := guo.mutation.TimeInterval(); ok {
		_spec.SetField(goal.FieldTimeInterval, field.TypeString, value)
	}
	if value, ok := guo.mutation.Status(); ok {
		_spec.SetField(goal.FieldStatus, field.TypeString, value)
	}
	if value, ok := guo.mutation.CurrentTotalValue(); ok {
		_spec.SetField(goal.FieldCurrentTotalValue, field.TypeString, value)
	}
	if guo.mutation.CurrentTotalValueCleared() {
		_spec.ClearField(goal.FieldCurrentTotalValue, field.TypeString)
	}
	if value, ok := guo.mutation.PerWorkoutData(); ok {
		_spec.SetField(goal.FieldPerWorkoutData, field.TypeJSON, value)
	}
	if value, ok := guo.mutation.AppendedPerWorkoutData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, goal.FieldPerWorkoutData, value)
		})
	}
	if guo.mutation.PerWorkoutDataCleared() {
		_spec.ClearField(goal.FieldPerWorkoutData, field.TypeJSON)
	}
	if guo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.ProfileTable,
			Columns: []string{goal.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.ProfileTable,
			Columns: []string{goal.ProfileColumn},
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
	if guo.mutation.CompetitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.CompetitionTable,
			Columns: []string{goal.CompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.CompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.CompetitionTable,
			Columns: []string{goal.CompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.SquadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.SquadTable,
			Columns: []string{goal.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.SquadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   goal.SquadTable,
			Columns: []string{goal.SquadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(guo.modifiers...)
	_node = &Goal{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
