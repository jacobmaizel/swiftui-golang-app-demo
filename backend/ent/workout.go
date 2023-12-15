// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
)

// Workout is the model entity for the Workout schema.
type Workout struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// HealthkitWorkoutActivityType holds the value of the "healthkit_workout_activity_type" field.
	HealthkitWorkoutActivityType string `json:"healthkit_workout_activity_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkoutQuery when eager-loading is set.
	Edges               WorkoutEdges `json:"related"`
	workout_leader      *uuid.UUID
	workout_competition *uuid.UUID
	selectValues        sql.SelectValues
}

// WorkoutEdges holds the relations/edges for other nodes in the graph.
type WorkoutEdges struct {
	// Invite holds the value of the invite edge.
	Invite []*Invite `json:"invite,omitempty"`
	// Leader holds the value of the leader edge.
	Leader *Profile `json:"leader,omitempty"`
	// Competition holds the value of the competition edge.
	Competition *Competition `json:"competition,omitempty"`
	// WorkoutData holds the value of the workout_data edge.
	WorkoutData []*WorkoutData `json:"workout_data,omitempty"`
	// WorkoutRouteData holds the value of the workout_route_data edge.
	WorkoutRouteData []*WorkoutRouteData `json:"workout_route_data,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// InviteOrErr returns the Invite value or an error if the edge
// was not loaded in eager-loading.
func (e WorkoutEdges) InviteOrErr() ([]*Invite, error) {
	if e.loadedTypes[0] {
		return e.Invite, nil
	}
	return nil, &NotLoadedError{edge: "invite"}
}

// LeaderOrErr returns the Leader value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkoutEdges) LeaderOrErr() (*Profile, error) {
	if e.loadedTypes[1] {
		if e.Leader == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Leader, nil
	}
	return nil, &NotLoadedError{edge: "leader"}
}

// CompetitionOrErr returns the Competition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkoutEdges) CompetitionOrErr() (*Competition, error) {
	if e.loadedTypes[2] {
		if e.Competition == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: competition.Label}
		}
		return e.Competition, nil
	}
	return nil, &NotLoadedError{edge: "competition"}
}

// WorkoutDataOrErr returns the WorkoutData value or an error if the edge
// was not loaded in eager-loading.
func (e WorkoutEdges) WorkoutDataOrErr() ([]*WorkoutData, error) {
	if e.loadedTypes[3] {
		return e.WorkoutData, nil
	}
	return nil, &NotLoadedError{edge: "workout_data"}
}

// WorkoutRouteDataOrErr returns the WorkoutRouteData value or an error if the edge
// was not loaded in eager-loading.
func (e WorkoutEdges) WorkoutRouteDataOrErr() ([]*WorkoutRouteData, error) {
	if e.loadedTypes[4] {
		return e.WorkoutRouteData, nil
	}
	return nil, &NotLoadedError{edge: "workout_route_data"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workout) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workout.FieldHealthkitWorkoutActivityType:
			values[i] = new(sql.NullString)
		case workout.FieldCreatedAt, workout.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case workout.FieldID:
			values[i] = new(uuid.UUID)
		case workout.ForeignKeys[0]: // workout_leader
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case workout.ForeignKeys[1]: // workout_competition
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workout fields.
func (w *Workout) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workout.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				w.ID = *value
			}
		case workout.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case workout.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		case workout.FieldHealthkitWorkoutActivityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field healthkit_workout_activity_type", values[i])
			} else if value.Valid {
				w.HealthkitWorkoutActivityType = value.String
			}
		case workout.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workout_leader", values[i])
			} else if value.Valid {
				w.workout_leader = new(uuid.UUID)
				*w.workout_leader = *value.S.(*uuid.UUID)
			}
		case workout.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workout_competition", values[i])
			} else if value.Valid {
				w.workout_competition = new(uuid.UUID)
				*w.workout_competition = *value.S.(*uuid.UUID)
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Workout.
// This includes values selected through modifiers, order, etc.
func (w *Workout) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryInvite queries the "invite" edge of the Workout entity.
func (w *Workout) QueryInvite() *InviteQuery {
	return NewWorkoutClient(w.config).QueryInvite(w)
}

// QueryLeader queries the "leader" edge of the Workout entity.
func (w *Workout) QueryLeader() *ProfileQuery {
	return NewWorkoutClient(w.config).QueryLeader(w)
}

// QueryCompetition queries the "competition" edge of the Workout entity.
func (w *Workout) QueryCompetition() *CompetitionQuery {
	return NewWorkoutClient(w.config).QueryCompetition(w)
}

// QueryWorkoutData queries the "workout_data" edge of the Workout entity.
func (w *Workout) QueryWorkoutData() *WorkoutDataQuery {
	return NewWorkoutClient(w.config).QueryWorkoutData(w)
}

// QueryWorkoutRouteData queries the "workout_route_data" edge of the Workout entity.
func (w *Workout) QueryWorkoutRouteData() *WorkoutRouteDataQuery {
	return NewWorkoutClient(w.config).QueryWorkoutRouteData(w)
}

// Update returns a builder for updating this Workout.
// Note that you need to call Workout.Unwrap() before calling this method if this Workout
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workout) Update() *WorkoutUpdateOne {
	return NewWorkoutClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Workout entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Workout) Unwrap() *Workout {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Workout is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workout) String() string {
	var builder strings.Builder
	builder.WriteString("Workout(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("healthkit_workout_activity_type=")
	builder.WriteString(w.HealthkitWorkoutActivityType)
	builder.WriteByte(')')
	return builder.String()
}

// Workouts is a parsable slice of Workout.
type Workouts []*Workout