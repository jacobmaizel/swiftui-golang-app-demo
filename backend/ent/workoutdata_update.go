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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutroutedata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

// WorkoutDataUpdate is the builder for updating WorkoutData entities.
type WorkoutDataUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkoutDataMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the WorkoutDataUpdate builder.
func (wdu *WorkoutDataUpdate) Where(ps ...predicate.WorkoutData) *WorkoutDataUpdate {
	wdu.mutation.Where(ps...)
	return wdu
}

// SetUpdatedAt sets the "updated_at" field.
func (wdu *WorkoutDataUpdate) SetUpdatedAt(t time.Time) *WorkoutDataUpdate {
	wdu.mutation.SetUpdatedAt(t)
	return wdu
}

// SetHealthkitWorkoutID sets the "healthkit_workout_id" field.
func (wdu *WorkoutDataUpdate) SetHealthkitWorkoutID(s string) *WorkoutDataUpdate {
	wdu.mutation.SetHealthkitWorkoutID(s)
	return wdu
}

// SetNillableHealthkitWorkoutID sets the "healthkit_workout_id" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableHealthkitWorkoutID(s *string) *WorkoutDataUpdate {
	if s != nil {
		wdu.SetHealthkitWorkoutID(*s)
	}
	return wdu
}

// ClearHealthkitWorkoutID clears the value of the "healthkit_workout_id" field.
func (wdu *WorkoutDataUpdate) ClearHealthkitWorkoutID() *WorkoutDataUpdate {
	wdu.mutation.ClearHealthkitWorkoutID()
	return wdu
}

// SetHealthkitWorkoutStartDate sets the "healthkit_workout_start_date" field.
func (wdu *WorkoutDataUpdate) SetHealthkitWorkoutStartDate(t time.Time) *WorkoutDataUpdate {
	wdu.mutation.SetHealthkitWorkoutStartDate(t)
	return wdu
}

// SetNillableHealthkitWorkoutStartDate sets the "healthkit_workout_start_date" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableHealthkitWorkoutStartDate(t *time.Time) *WorkoutDataUpdate {
	if t != nil {
		wdu.SetHealthkitWorkoutStartDate(*t)
	}
	return wdu
}

// ClearHealthkitWorkoutStartDate clears the value of the "healthkit_workout_start_date" field.
func (wdu *WorkoutDataUpdate) ClearHealthkitWorkoutStartDate() *WorkoutDataUpdate {
	wdu.mutation.ClearHealthkitWorkoutStartDate()
	return wdu
}

// SetHealthkitWorkoutEndDate sets the "healthkit_workout_end_date" field.
func (wdu *WorkoutDataUpdate) SetHealthkitWorkoutEndDate(t time.Time) *WorkoutDataUpdate {
	wdu.mutation.SetHealthkitWorkoutEndDate(t)
	return wdu
}

// SetNillableHealthkitWorkoutEndDate sets the "healthkit_workout_end_date" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableHealthkitWorkoutEndDate(t *time.Time) *WorkoutDataUpdate {
	if t != nil {
		wdu.SetHealthkitWorkoutEndDate(*t)
	}
	return wdu
}

// ClearHealthkitWorkoutEndDate clears the value of the "healthkit_workout_end_date" field.
func (wdu *WorkoutDataUpdate) ClearHealthkitWorkoutEndDate() *WorkoutDataUpdate {
	wdu.mutation.ClearHealthkitWorkoutEndDate()
	return wdu
}

// SetDistance sets the "distance" field.
func (wdu *WorkoutDataUpdate) SetDistance(f float64) *WorkoutDataUpdate {
	wdu.mutation.ResetDistance()
	wdu.mutation.SetDistance(f)
	return wdu
}

// SetNillableDistance sets the "distance" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableDistance(f *float64) *WorkoutDataUpdate {
	if f != nil {
		wdu.SetDistance(*f)
	}
	return wdu
}

// AddDistance adds f to the "distance" field.
func (wdu *WorkoutDataUpdate) AddDistance(f float64) *WorkoutDataUpdate {
	wdu.mutation.AddDistance(f)
	return wdu
}

// ClearDistance clears the value of the "distance" field.
func (wdu *WorkoutDataUpdate) ClearDistance() *WorkoutDataUpdate {
	wdu.mutation.ClearDistance()
	return wdu
}

// SetDuration sets the "duration" field.
func (wdu *WorkoutDataUpdate) SetDuration(s string) *WorkoutDataUpdate {
	wdu.mutation.SetDuration(s)
	return wdu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableDuration(s *string) *WorkoutDataUpdate {
	if s != nil {
		wdu.SetDuration(*s)
	}
	return wdu
}

// ClearDuration clears the value of the "duration" field.
func (wdu *WorkoutDataUpdate) ClearDuration() *WorkoutDataUpdate {
	wdu.mutation.ClearDuration()
	return wdu
}

// SetEnergyBurned sets the "energy_burned" field.
func (wdu *WorkoutDataUpdate) SetEnergyBurned(s string) *WorkoutDataUpdate {
	wdu.mutation.SetEnergyBurned(s)
	return wdu
}

// SetNillableEnergyBurned sets the "energy_burned" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableEnergyBurned(s *string) *WorkoutDataUpdate {
	if s != nil {
		wdu.SetEnergyBurned(*s)
	}
	return wdu
}

// ClearEnergyBurned clears the value of the "energy_burned" field.
func (wdu *WorkoutDataUpdate) ClearEnergyBurned() *WorkoutDataUpdate {
	wdu.mutation.ClearEnergyBurned()
	return wdu
}

// SetAverageHeartRate sets the "average_heart_rate" field.
func (wdu *WorkoutDataUpdate) SetAverageHeartRate(s string) *WorkoutDataUpdate {
	wdu.mutation.SetAverageHeartRate(s)
	return wdu
}

// SetNillableAverageHeartRate sets the "average_heart_rate" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableAverageHeartRate(s *string) *WorkoutDataUpdate {
	if s != nil {
		wdu.SetAverageHeartRate(*s)
	}
	return wdu
}

// ClearAverageHeartRate clears the value of the "average_heart_rate" field.
func (wdu *WorkoutDataUpdate) ClearAverageHeartRate() *WorkoutDataUpdate {
	wdu.mutation.ClearAverageHeartRate()
	return wdu
}

// SetSource sets the "source" field.
func (wdu *WorkoutDataUpdate) SetSource(s string) *WorkoutDataUpdate {
	wdu.mutation.SetSource(s)
	return wdu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableSource(s *string) *WorkoutDataUpdate {
	if s != nil {
		wdu.SetSource(*s)
	}
	return wdu
}

// SetLocationType sets the "location_type" field.
func (wdu *WorkoutDataUpdate) SetLocationType(s string) *WorkoutDataUpdate {
	wdu.mutation.SetLocationType(s)
	return wdu
}

// SetNillableLocationType sets the "location_type" field if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableLocationType(s *string) *WorkoutDataUpdate {
	if s != nil {
		wdu.SetLocationType(*s)
	}
	return wdu
}

// SetWeather sets the "weather" field.
func (wdu *WorkoutDataUpdate) SetWeather(sdw *shared.WorkoutDataWeather) *WorkoutDataUpdate {
	wdu.mutation.SetWeather(sdw)
	return wdu
}

// ClearWeather clears the value of the "weather" field.
func (wdu *WorkoutDataUpdate) ClearWeather() *WorkoutDataUpdate {
	wdu.mutation.ClearWeather()
	return wdu
}

// SetWorkoutID sets the "workout" edge to the Workout entity by ID.
func (wdu *WorkoutDataUpdate) SetWorkoutID(id uuid.UUID) *WorkoutDataUpdate {
	wdu.mutation.SetWorkoutID(id)
	return wdu
}

// SetNillableWorkoutID sets the "workout" edge to the Workout entity by ID if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableWorkoutID(id *uuid.UUID) *WorkoutDataUpdate {
	if id != nil {
		wdu = wdu.SetWorkoutID(*id)
	}
	return wdu
}

// SetWorkout sets the "workout" edge to the Workout entity.
func (wdu *WorkoutDataUpdate) SetWorkout(w *Workout) *WorkoutDataUpdate {
	return wdu.SetWorkoutID(w.ID)
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (wdu *WorkoutDataUpdate) SetProfileID(id uuid.UUID) *WorkoutDataUpdate {
	wdu.mutation.SetProfileID(id)
	return wdu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (wdu *WorkoutDataUpdate) SetProfile(p *Profile) *WorkoutDataUpdate {
	return wdu.SetProfileID(p.ID)
}

// SetWorkoutRouteDataID sets the "workout_route_data" edge to the WorkoutRouteData entity by ID.
func (wdu *WorkoutDataUpdate) SetWorkoutRouteDataID(id uuid.UUID) *WorkoutDataUpdate {
	wdu.mutation.SetWorkoutRouteDataID(id)
	return wdu
}

// SetNillableWorkoutRouteDataID sets the "workout_route_data" edge to the WorkoutRouteData entity by ID if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableWorkoutRouteDataID(id *uuid.UUID) *WorkoutDataUpdate {
	if id != nil {
		wdu = wdu.SetWorkoutRouteDataID(*id)
	}
	return wdu
}

// SetWorkoutRouteData sets the "workout_route_data" edge to the WorkoutRouteData entity.
func (wdu *WorkoutDataUpdate) SetWorkoutRouteData(w *WorkoutRouteData) *WorkoutDataUpdate {
	return wdu.SetWorkoutRouteDataID(w.ID)
}

// SetCompetitionID sets the "competition" edge to the Competition entity by ID.
func (wdu *WorkoutDataUpdate) SetCompetitionID(id uuid.UUID) *WorkoutDataUpdate {
	wdu.mutation.SetCompetitionID(id)
	return wdu
}

// SetNillableCompetitionID sets the "competition" edge to the Competition entity by ID if the given value is not nil.
func (wdu *WorkoutDataUpdate) SetNillableCompetitionID(id *uuid.UUID) *WorkoutDataUpdate {
	if id != nil {
		wdu = wdu.SetCompetitionID(*id)
	}
	return wdu
}

// SetCompetition sets the "competition" edge to the Competition entity.
func (wdu *WorkoutDataUpdate) SetCompetition(c *Competition) *WorkoutDataUpdate {
	return wdu.SetCompetitionID(c.ID)
}

// Mutation returns the WorkoutDataMutation object of the builder.
func (wdu *WorkoutDataUpdate) Mutation() *WorkoutDataMutation {
	return wdu.mutation
}

// ClearWorkout clears the "workout" edge to the Workout entity.
func (wdu *WorkoutDataUpdate) ClearWorkout() *WorkoutDataUpdate {
	wdu.mutation.ClearWorkout()
	return wdu
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (wdu *WorkoutDataUpdate) ClearProfile() *WorkoutDataUpdate {
	wdu.mutation.ClearProfile()
	return wdu
}

// ClearWorkoutRouteData clears the "workout_route_data" edge to the WorkoutRouteData entity.
func (wdu *WorkoutDataUpdate) ClearWorkoutRouteData() *WorkoutDataUpdate {
	wdu.mutation.ClearWorkoutRouteData()
	return wdu
}

// ClearCompetition clears the "competition" edge to the Competition entity.
func (wdu *WorkoutDataUpdate) ClearCompetition() *WorkoutDataUpdate {
	wdu.mutation.ClearCompetition()
	return wdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wdu *WorkoutDataUpdate) Save(ctx context.Context) (int, error) {
	wdu.defaults()
	return withHooks(ctx, wdu.sqlSave, wdu.mutation, wdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wdu *WorkoutDataUpdate) SaveX(ctx context.Context) int {
	affected, err := wdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wdu *WorkoutDataUpdate) Exec(ctx context.Context) error {
	_, err := wdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdu *WorkoutDataUpdate) ExecX(ctx context.Context) {
	if err := wdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wdu *WorkoutDataUpdate) defaults() {
	if _, ok := wdu.mutation.UpdatedAt(); !ok {
		v := workoutdata.UpdateDefaultUpdatedAt()
		wdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wdu *WorkoutDataUpdate) check() error {
	if v, ok := wdu.mutation.Source(); ok {
		if err := workoutdata.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "WorkoutData.source": %w`, err)}
		}
	}
	if v, ok := wdu.mutation.LocationType(); ok {
		if err := workoutdata.LocationTypeValidator(v); err != nil {
			return &ValidationError{Name: "location_type", err: fmt.Errorf(`ent: validator failed for field "WorkoutData.location_type": %w`, err)}
		}
	}
	if _, ok := wdu.mutation.ProfileID(); wdu.mutation.ProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutData.profile"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wdu *WorkoutDataUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkoutDataUpdate {
	wdu.modifiers = append(wdu.modifiers, modifiers...)
	return wdu
}

func (wdu *WorkoutDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workoutdata.Table, workoutdata.Columns, sqlgraph.NewFieldSpec(workoutdata.FieldID, field.TypeUUID))
	if ps := wdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wdu.mutation.UpdatedAt(); ok {
		_spec.SetField(workoutdata.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wdu.mutation.HealthkitWorkoutID(); ok {
		_spec.SetField(workoutdata.FieldHealthkitWorkoutID, field.TypeString, value)
	}
	if wdu.mutation.HealthkitWorkoutIDCleared() {
		_spec.ClearField(workoutdata.FieldHealthkitWorkoutID, field.TypeString)
	}
	if value, ok := wdu.mutation.HealthkitWorkoutStartDate(); ok {
		_spec.SetField(workoutdata.FieldHealthkitWorkoutStartDate, field.TypeTime, value)
	}
	if wdu.mutation.HealthkitWorkoutStartDateCleared() {
		_spec.ClearField(workoutdata.FieldHealthkitWorkoutStartDate, field.TypeTime)
	}
	if value, ok := wdu.mutation.HealthkitWorkoutEndDate(); ok {
		_spec.SetField(workoutdata.FieldHealthkitWorkoutEndDate, field.TypeTime, value)
	}
	if wdu.mutation.HealthkitWorkoutEndDateCleared() {
		_spec.ClearField(workoutdata.FieldHealthkitWorkoutEndDate, field.TypeTime)
	}
	if value, ok := wdu.mutation.Distance(); ok {
		_spec.SetField(workoutdata.FieldDistance, field.TypeFloat64, value)
	}
	if value, ok := wdu.mutation.AddedDistance(); ok {
		_spec.AddField(workoutdata.FieldDistance, field.TypeFloat64, value)
	}
	if wdu.mutation.DistanceCleared() {
		_spec.ClearField(workoutdata.FieldDistance, field.TypeFloat64)
	}
	if value, ok := wdu.mutation.Duration(); ok {
		_spec.SetField(workoutdata.FieldDuration, field.TypeString, value)
	}
	if wdu.mutation.DurationCleared() {
		_spec.ClearField(workoutdata.FieldDuration, field.TypeString)
	}
	if value, ok := wdu.mutation.EnergyBurned(); ok {
		_spec.SetField(workoutdata.FieldEnergyBurned, field.TypeString, value)
	}
	if wdu.mutation.EnergyBurnedCleared() {
		_spec.ClearField(workoutdata.FieldEnergyBurned, field.TypeString)
	}
	if value, ok := wdu.mutation.AverageHeartRate(); ok {
		_spec.SetField(workoutdata.FieldAverageHeartRate, field.TypeString, value)
	}
	if wdu.mutation.AverageHeartRateCleared() {
		_spec.ClearField(workoutdata.FieldAverageHeartRate, field.TypeString)
	}
	if value, ok := wdu.mutation.Source(); ok {
		_spec.SetField(workoutdata.FieldSource, field.TypeString, value)
	}
	if value, ok := wdu.mutation.LocationType(); ok {
		_spec.SetField(workoutdata.FieldLocationType, field.TypeString, value)
	}
	if value, ok := wdu.mutation.Weather(); ok {
		_spec.SetField(workoutdata.FieldWeather, field.TypeJSON, value)
	}
	if wdu.mutation.WeatherCleared() {
		_spec.ClearField(workoutdata.FieldWeather, field.TypeJSON)
	}
	if wdu.mutation.WorkoutCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.WorkoutTable,
			Columns: []string{workoutdata.WorkoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wdu.mutation.WorkoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.WorkoutTable,
			Columns: []string{workoutdata.WorkoutColumn},
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
	if wdu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.ProfileTable,
			Columns: []string{workoutdata.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wdu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.ProfileTable,
			Columns: []string{workoutdata.ProfileColumn},
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
	if wdu.mutation.WorkoutRouteDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workoutdata.WorkoutRouteDataTable,
			Columns: []string{workoutdata.WorkoutRouteDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workoutroutedata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wdu.mutation.WorkoutRouteDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workoutdata.WorkoutRouteDataTable,
			Columns: []string{workoutdata.WorkoutRouteDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workoutroutedata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wdu.mutation.CompetitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.CompetitionTable,
			Columns: []string{workoutdata.CompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wdu.mutation.CompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.CompetitionTable,
			Columns: []string{workoutdata.CompetitionColumn},
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
	_spec.AddModifiers(wdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workoutdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wdu.mutation.done = true
	return n, nil
}

// WorkoutDataUpdateOne is the builder for updating a single WorkoutData entity.
type WorkoutDataUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkoutDataMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (wduo *WorkoutDataUpdateOne) SetUpdatedAt(t time.Time) *WorkoutDataUpdateOne {
	wduo.mutation.SetUpdatedAt(t)
	return wduo
}

// SetHealthkitWorkoutID sets the "healthkit_workout_id" field.
func (wduo *WorkoutDataUpdateOne) SetHealthkitWorkoutID(s string) *WorkoutDataUpdateOne {
	wduo.mutation.SetHealthkitWorkoutID(s)
	return wduo
}

// SetNillableHealthkitWorkoutID sets the "healthkit_workout_id" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableHealthkitWorkoutID(s *string) *WorkoutDataUpdateOne {
	if s != nil {
		wduo.SetHealthkitWorkoutID(*s)
	}
	return wduo
}

// ClearHealthkitWorkoutID clears the value of the "healthkit_workout_id" field.
func (wduo *WorkoutDataUpdateOne) ClearHealthkitWorkoutID() *WorkoutDataUpdateOne {
	wduo.mutation.ClearHealthkitWorkoutID()
	return wduo
}

// SetHealthkitWorkoutStartDate sets the "healthkit_workout_start_date" field.
func (wduo *WorkoutDataUpdateOne) SetHealthkitWorkoutStartDate(t time.Time) *WorkoutDataUpdateOne {
	wduo.mutation.SetHealthkitWorkoutStartDate(t)
	return wduo
}

// SetNillableHealthkitWorkoutStartDate sets the "healthkit_workout_start_date" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableHealthkitWorkoutStartDate(t *time.Time) *WorkoutDataUpdateOne {
	if t != nil {
		wduo.SetHealthkitWorkoutStartDate(*t)
	}
	return wduo
}

// ClearHealthkitWorkoutStartDate clears the value of the "healthkit_workout_start_date" field.
func (wduo *WorkoutDataUpdateOne) ClearHealthkitWorkoutStartDate() *WorkoutDataUpdateOne {
	wduo.mutation.ClearHealthkitWorkoutStartDate()
	return wduo
}

// SetHealthkitWorkoutEndDate sets the "healthkit_workout_end_date" field.
func (wduo *WorkoutDataUpdateOne) SetHealthkitWorkoutEndDate(t time.Time) *WorkoutDataUpdateOne {
	wduo.mutation.SetHealthkitWorkoutEndDate(t)
	return wduo
}

// SetNillableHealthkitWorkoutEndDate sets the "healthkit_workout_end_date" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableHealthkitWorkoutEndDate(t *time.Time) *WorkoutDataUpdateOne {
	if t != nil {
		wduo.SetHealthkitWorkoutEndDate(*t)
	}
	return wduo
}

// ClearHealthkitWorkoutEndDate clears the value of the "healthkit_workout_end_date" field.
func (wduo *WorkoutDataUpdateOne) ClearHealthkitWorkoutEndDate() *WorkoutDataUpdateOne {
	wduo.mutation.ClearHealthkitWorkoutEndDate()
	return wduo
}

// SetDistance sets the "distance" field.
func (wduo *WorkoutDataUpdateOne) SetDistance(f float64) *WorkoutDataUpdateOne {
	wduo.mutation.ResetDistance()
	wduo.mutation.SetDistance(f)
	return wduo
}

// SetNillableDistance sets the "distance" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableDistance(f *float64) *WorkoutDataUpdateOne {
	if f != nil {
		wduo.SetDistance(*f)
	}
	return wduo
}

// AddDistance adds f to the "distance" field.
func (wduo *WorkoutDataUpdateOne) AddDistance(f float64) *WorkoutDataUpdateOne {
	wduo.mutation.AddDistance(f)
	return wduo
}

// ClearDistance clears the value of the "distance" field.
func (wduo *WorkoutDataUpdateOne) ClearDistance() *WorkoutDataUpdateOne {
	wduo.mutation.ClearDistance()
	return wduo
}

// SetDuration sets the "duration" field.
func (wduo *WorkoutDataUpdateOne) SetDuration(s string) *WorkoutDataUpdateOne {
	wduo.mutation.SetDuration(s)
	return wduo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableDuration(s *string) *WorkoutDataUpdateOne {
	if s != nil {
		wduo.SetDuration(*s)
	}
	return wduo
}

// ClearDuration clears the value of the "duration" field.
func (wduo *WorkoutDataUpdateOne) ClearDuration() *WorkoutDataUpdateOne {
	wduo.mutation.ClearDuration()
	return wduo
}

// SetEnergyBurned sets the "energy_burned" field.
func (wduo *WorkoutDataUpdateOne) SetEnergyBurned(s string) *WorkoutDataUpdateOne {
	wduo.mutation.SetEnergyBurned(s)
	return wduo
}

// SetNillableEnergyBurned sets the "energy_burned" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableEnergyBurned(s *string) *WorkoutDataUpdateOne {
	if s != nil {
		wduo.SetEnergyBurned(*s)
	}
	return wduo
}

// ClearEnergyBurned clears the value of the "energy_burned" field.
func (wduo *WorkoutDataUpdateOne) ClearEnergyBurned() *WorkoutDataUpdateOne {
	wduo.mutation.ClearEnergyBurned()
	return wduo
}

// SetAverageHeartRate sets the "average_heart_rate" field.
func (wduo *WorkoutDataUpdateOne) SetAverageHeartRate(s string) *WorkoutDataUpdateOne {
	wduo.mutation.SetAverageHeartRate(s)
	return wduo
}

// SetNillableAverageHeartRate sets the "average_heart_rate" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableAverageHeartRate(s *string) *WorkoutDataUpdateOne {
	if s != nil {
		wduo.SetAverageHeartRate(*s)
	}
	return wduo
}

// ClearAverageHeartRate clears the value of the "average_heart_rate" field.
func (wduo *WorkoutDataUpdateOne) ClearAverageHeartRate() *WorkoutDataUpdateOne {
	wduo.mutation.ClearAverageHeartRate()
	return wduo
}

// SetSource sets the "source" field.
func (wduo *WorkoutDataUpdateOne) SetSource(s string) *WorkoutDataUpdateOne {
	wduo.mutation.SetSource(s)
	return wduo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableSource(s *string) *WorkoutDataUpdateOne {
	if s != nil {
		wduo.SetSource(*s)
	}
	return wduo
}

// SetLocationType sets the "location_type" field.
func (wduo *WorkoutDataUpdateOne) SetLocationType(s string) *WorkoutDataUpdateOne {
	wduo.mutation.SetLocationType(s)
	return wduo
}

// SetNillableLocationType sets the "location_type" field if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableLocationType(s *string) *WorkoutDataUpdateOne {
	if s != nil {
		wduo.SetLocationType(*s)
	}
	return wduo
}

// SetWeather sets the "weather" field.
func (wduo *WorkoutDataUpdateOne) SetWeather(sdw *shared.WorkoutDataWeather) *WorkoutDataUpdateOne {
	wduo.mutation.SetWeather(sdw)
	return wduo
}

// ClearWeather clears the value of the "weather" field.
func (wduo *WorkoutDataUpdateOne) ClearWeather() *WorkoutDataUpdateOne {
	wduo.mutation.ClearWeather()
	return wduo
}

// SetWorkoutID sets the "workout" edge to the Workout entity by ID.
func (wduo *WorkoutDataUpdateOne) SetWorkoutID(id uuid.UUID) *WorkoutDataUpdateOne {
	wduo.mutation.SetWorkoutID(id)
	return wduo
}

// SetNillableWorkoutID sets the "workout" edge to the Workout entity by ID if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableWorkoutID(id *uuid.UUID) *WorkoutDataUpdateOne {
	if id != nil {
		wduo = wduo.SetWorkoutID(*id)
	}
	return wduo
}

// SetWorkout sets the "workout" edge to the Workout entity.
func (wduo *WorkoutDataUpdateOne) SetWorkout(w *Workout) *WorkoutDataUpdateOne {
	return wduo.SetWorkoutID(w.ID)
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (wduo *WorkoutDataUpdateOne) SetProfileID(id uuid.UUID) *WorkoutDataUpdateOne {
	wduo.mutation.SetProfileID(id)
	return wduo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (wduo *WorkoutDataUpdateOne) SetProfile(p *Profile) *WorkoutDataUpdateOne {
	return wduo.SetProfileID(p.ID)
}

// SetWorkoutRouteDataID sets the "workout_route_data" edge to the WorkoutRouteData entity by ID.
func (wduo *WorkoutDataUpdateOne) SetWorkoutRouteDataID(id uuid.UUID) *WorkoutDataUpdateOne {
	wduo.mutation.SetWorkoutRouteDataID(id)
	return wduo
}

// SetNillableWorkoutRouteDataID sets the "workout_route_data" edge to the WorkoutRouteData entity by ID if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableWorkoutRouteDataID(id *uuid.UUID) *WorkoutDataUpdateOne {
	if id != nil {
		wduo = wduo.SetWorkoutRouteDataID(*id)
	}
	return wduo
}

// SetWorkoutRouteData sets the "workout_route_data" edge to the WorkoutRouteData entity.
func (wduo *WorkoutDataUpdateOne) SetWorkoutRouteData(w *WorkoutRouteData) *WorkoutDataUpdateOne {
	return wduo.SetWorkoutRouteDataID(w.ID)
}

// SetCompetitionID sets the "competition" edge to the Competition entity by ID.
func (wduo *WorkoutDataUpdateOne) SetCompetitionID(id uuid.UUID) *WorkoutDataUpdateOne {
	wduo.mutation.SetCompetitionID(id)
	return wduo
}

// SetNillableCompetitionID sets the "competition" edge to the Competition entity by ID if the given value is not nil.
func (wduo *WorkoutDataUpdateOne) SetNillableCompetitionID(id *uuid.UUID) *WorkoutDataUpdateOne {
	if id != nil {
		wduo = wduo.SetCompetitionID(*id)
	}
	return wduo
}

// SetCompetition sets the "competition" edge to the Competition entity.
func (wduo *WorkoutDataUpdateOne) SetCompetition(c *Competition) *WorkoutDataUpdateOne {
	return wduo.SetCompetitionID(c.ID)
}

// Mutation returns the WorkoutDataMutation object of the builder.
func (wduo *WorkoutDataUpdateOne) Mutation() *WorkoutDataMutation {
	return wduo.mutation
}

// ClearWorkout clears the "workout" edge to the Workout entity.
func (wduo *WorkoutDataUpdateOne) ClearWorkout() *WorkoutDataUpdateOne {
	wduo.mutation.ClearWorkout()
	return wduo
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (wduo *WorkoutDataUpdateOne) ClearProfile() *WorkoutDataUpdateOne {
	wduo.mutation.ClearProfile()
	return wduo
}

// ClearWorkoutRouteData clears the "workout_route_data" edge to the WorkoutRouteData entity.
func (wduo *WorkoutDataUpdateOne) ClearWorkoutRouteData() *WorkoutDataUpdateOne {
	wduo.mutation.ClearWorkoutRouteData()
	return wduo
}

// ClearCompetition clears the "competition" edge to the Competition entity.
func (wduo *WorkoutDataUpdateOne) ClearCompetition() *WorkoutDataUpdateOne {
	wduo.mutation.ClearCompetition()
	return wduo
}

// Where appends a list predicates to the WorkoutDataUpdate builder.
func (wduo *WorkoutDataUpdateOne) Where(ps ...predicate.WorkoutData) *WorkoutDataUpdateOne {
	wduo.mutation.Where(ps...)
	return wduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wduo *WorkoutDataUpdateOne) Select(field string, fields ...string) *WorkoutDataUpdateOne {
	wduo.fields = append([]string{field}, fields...)
	return wduo
}

// Save executes the query and returns the updated WorkoutData entity.
func (wduo *WorkoutDataUpdateOne) Save(ctx context.Context) (*WorkoutData, error) {
	wduo.defaults()
	return withHooks(ctx, wduo.sqlSave, wduo.mutation, wduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wduo *WorkoutDataUpdateOne) SaveX(ctx context.Context) *WorkoutData {
	node, err := wduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wduo *WorkoutDataUpdateOne) Exec(ctx context.Context) error {
	_, err := wduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wduo *WorkoutDataUpdateOne) ExecX(ctx context.Context) {
	if err := wduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wduo *WorkoutDataUpdateOne) defaults() {
	if _, ok := wduo.mutation.UpdatedAt(); !ok {
		v := workoutdata.UpdateDefaultUpdatedAt()
		wduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wduo *WorkoutDataUpdateOne) check() error {
	if v, ok := wduo.mutation.Source(); ok {
		if err := workoutdata.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "WorkoutData.source": %w`, err)}
		}
	}
	if v, ok := wduo.mutation.LocationType(); ok {
		if err := workoutdata.LocationTypeValidator(v); err != nil {
			return &ValidationError{Name: "location_type", err: fmt.Errorf(`ent: validator failed for field "WorkoutData.location_type": %w`, err)}
		}
	}
	if _, ok := wduo.mutation.ProfileID(); wduo.mutation.ProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutData.profile"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wduo *WorkoutDataUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkoutDataUpdateOne {
	wduo.modifiers = append(wduo.modifiers, modifiers...)
	return wduo
}

func (wduo *WorkoutDataUpdateOne) sqlSave(ctx context.Context) (_node *WorkoutData, err error) {
	if err := wduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workoutdata.Table, workoutdata.Columns, sqlgraph.NewFieldSpec(workoutdata.FieldID, field.TypeUUID))
	id, ok := wduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkoutData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workoutdata.FieldID)
		for _, f := range fields {
			if !workoutdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workoutdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wduo.mutation.UpdatedAt(); ok {
		_spec.SetField(workoutdata.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wduo.mutation.HealthkitWorkoutID(); ok {
		_spec.SetField(workoutdata.FieldHealthkitWorkoutID, field.TypeString, value)
	}
	if wduo.mutation.HealthkitWorkoutIDCleared() {
		_spec.ClearField(workoutdata.FieldHealthkitWorkoutID, field.TypeString)
	}
	if value, ok := wduo.mutation.HealthkitWorkoutStartDate(); ok {
		_spec.SetField(workoutdata.FieldHealthkitWorkoutStartDate, field.TypeTime, value)
	}
	if wduo.mutation.HealthkitWorkoutStartDateCleared() {
		_spec.ClearField(workoutdata.FieldHealthkitWorkoutStartDate, field.TypeTime)
	}
	if value, ok := wduo.mutation.HealthkitWorkoutEndDate(); ok {
		_spec.SetField(workoutdata.FieldHealthkitWorkoutEndDate, field.TypeTime, value)
	}
	if wduo.mutation.HealthkitWorkoutEndDateCleared() {
		_spec.ClearField(workoutdata.FieldHealthkitWorkoutEndDate, field.TypeTime)
	}
	if value, ok := wduo.mutation.Distance(); ok {
		_spec.SetField(workoutdata.FieldDistance, field.TypeFloat64, value)
	}
	if value, ok := wduo.mutation.AddedDistance(); ok {
		_spec.AddField(workoutdata.FieldDistance, field.TypeFloat64, value)
	}
	if wduo.mutation.DistanceCleared() {
		_spec.ClearField(workoutdata.FieldDistance, field.TypeFloat64)
	}
	if value, ok := wduo.mutation.Duration(); ok {
		_spec.SetField(workoutdata.FieldDuration, field.TypeString, value)
	}
	if wduo.mutation.DurationCleared() {
		_spec.ClearField(workoutdata.FieldDuration, field.TypeString)
	}
	if value, ok := wduo.mutation.EnergyBurned(); ok {
		_spec.SetField(workoutdata.FieldEnergyBurned, field.TypeString, value)
	}
	if wduo.mutation.EnergyBurnedCleared() {
		_spec.ClearField(workoutdata.FieldEnergyBurned, field.TypeString)
	}
	if value, ok := wduo.mutation.AverageHeartRate(); ok {
		_spec.SetField(workoutdata.FieldAverageHeartRate, field.TypeString, value)
	}
	if wduo.mutation.AverageHeartRateCleared() {
		_spec.ClearField(workoutdata.FieldAverageHeartRate, field.TypeString)
	}
	if value, ok := wduo.mutation.Source(); ok {
		_spec.SetField(workoutdata.FieldSource, field.TypeString, value)
	}
	if value, ok := wduo.mutation.LocationType(); ok {
		_spec.SetField(workoutdata.FieldLocationType, field.TypeString, value)
	}
	if value, ok := wduo.mutation.Weather(); ok {
		_spec.SetField(workoutdata.FieldWeather, field.TypeJSON, value)
	}
	if wduo.mutation.WeatherCleared() {
		_spec.ClearField(workoutdata.FieldWeather, field.TypeJSON)
	}
	if wduo.mutation.WorkoutCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.WorkoutTable,
			Columns: []string{workoutdata.WorkoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wduo.mutation.WorkoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.WorkoutTable,
			Columns: []string{workoutdata.WorkoutColumn},
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
	if wduo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.ProfileTable,
			Columns: []string{workoutdata.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wduo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.ProfileTable,
			Columns: []string{workoutdata.ProfileColumn},
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
	if wduo.mutation.WorkoutRouteDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workoutdata.WorkoutRouteDataTable,
			Columns: []string{workoutdata.WorkoutRouteDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workoutroutedata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wduo.mutation.WorkoutRouteDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   workoutdata.WorkoutRouteDataTable,
			Columns: []string{workoutdata.WorkoutRouteDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workoutroutedata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wduo.mutation.CompetitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.CompetitionTable,
			Columns: []string{workoutdata.CompetitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wduo.mutation.CompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutdata.CompetitionTable,
			Columns: []string{workoutdata.CompetitionColumn},
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
	_spec.AddModifiers(wduo.modifiers...)
	_node = &WorkoutData{config: wduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workoutdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wduo.mutation.done = true
	return _node, nil
}
