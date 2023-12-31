// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
)

// Competition is the model entity for the Competition schema.
type Competition struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Public holds the value of the "public" field.
	Public bool `json:"public"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Start holds the value of the "start" field.
	Start time.Time `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End time.Time `json:"end,omitempty"`
	// Scheduled holds the value of the "scheduled" field.
	Scheduled bool `json:"scheduled"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// ParticipantTypes holds the value of the "participant_types" field.
	ParticipantTypes []string `json:"participant_types,omitempty"`
	// WorkoutTypes holds the value of the "workout_types" field.
	WorkoutTypes []string `json:"workout_types,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompetitionQuery when eager-loading is set.
	Edges             CompetitionEdges `json:"related"`
	competition_owner *uuid.UUID
	selectValues      sql.SelectValues
}

// CompetitionEdges holds the relations/edges for other nodes in the graph.
type CompetitionEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Profile `json:"owner,omitempty"`
	// Participants holds the value of the participants edge.
	Participants []*Profile `json:"participants,omitempty"`
	// Workouts holds the value of the workouts edge.
	Workouts []*Workout `json:"workouts,omitempty"`
	// WorkoutData holds the value of the workout_data edge.
	WorkoutData []*WorkoutData `json:"workout_data,omitempty"`
	// Results holds the value of the results edge.
	Results []*CompetitionResult `json:"results,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompetitionEdges) OwnerOrErr() (*Profile, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ParticipantsOrErr returns the Participants value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) ParticipantsOrErr() ([]*Profile, error) {
	if e.loadedTypes[1] {
		return e.Participants, nil
	}
	return nil, &NotLoadedError{edge: "participants"}
}

// WorkoutsOrErr returns the Workouts value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) WorkoutsOrErr() ([]*Workout, error) {
	if e.loadedTypes[2] {
		return e.Workouts, nil
	}
	return nil, &NotLoadedError{edge: "workouts"}
}

// WorkoutDataOrErr returns the WorkoutData value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) WorkoutDataOrErr() ([]*WorkoutData, error) {
	if e.loadedTypes[3] {
		return e.WorkoutData, nil
	}
	return nil, &NotLoadedError{edge: "workout_data"}
}

// ResultsOrErr returns the Results value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) ResultsOrErr() ([]*CompetitionResult, error) {
	if e.loadedTypes[4] {
		return e.Results, nil
	}
	return nil, &NotLoadedError{edge: "results"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Competition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case competition.FieldParticipantTypes, competition.FieldWorkoutTypes:
			values[i] = new([]byte)
		case competition.FieldPublic, competition.FieldScheduled:
			values[i] = new(sql.NullBool)
		case competition.FieldTitle, competition.FieldDescription, competition.FieldStatus, competition.FieldType:
			values[i] = new(sql.NullString)
		case competition.FieldCreatedAt, competition.FieldUpdatedAt, competition.FieldStart, competition.FieldEnd:
			values[i] = new(sql.NullTime)
		case competition.FieldID:
			values[i] = new(uuid.UUID)
		case competition.ForeignKeys[0]: // competition_owner
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Competition fields.
func (c *Competition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case competition.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case competition.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case competition.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case competition.FieldPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public", values[i])
			} else if value.Valid {
				c.Public = value.Bool
			}
		case competition.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case competition.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case competition.FieldStart:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				c.Start = value.Time
			}
		case competition.FieldEnd:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				c.End = value.Time
			}
		case competition.FieldScheduled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field scheduled", values[i])
			} else if value.Valid {
				c.Scheduled = value.Bool
			}
		case competition.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = value.String
			}
		case competition.FieldParticipantTypes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field participant_types", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.ParticipantTypes); err != nil {
					return fmt.Errorf("unmarshal field participant_types: %w", err)
				}
			}
		case competition.FieldWorkoutTypes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field workout_types", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.WorkoutTypes); err != nil {
					return fmt.Errorf("unmarshal field workout_types: %w", err)
				}
			}
		case competition.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case competition.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field competition_owner", values[i])
			} else if value.Valid {
				c.competition_owner = new(uuid.UUID)
				*c.competition_owner = *value.S.(*uuid.UUID)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Competition.
// This includes values selected through modifiers, order, etc.
func (c *Competition) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Competition entity.
func (c *Competition) QueryOwner() *ProfileQuery {
	return NewCompetitionClient(c.config).QueryOwner(c)
}

// QueryParticipants queries the "participants" edge of the Competition entity.
func (c *Competition) QueryParticipants() *ProfileQuery {
	return NewCompetitionClient(c.config).QueryParticipants(c)
}

// QueryWorkouts queries the "workouts" edge of the Competition entity.
func (c *Competition) QueryWorkouts() *WorkoutQuery {
	return NewCompetitionClient(c.config).QueryWorkouts(c)
}

// QueryWorkoutData queries the "workout_data" edge of the Competition entity.
func (c *Competition) QueryWorkoutData() *WorkoutDataQuery {
	return NewCompetitionClient(c.config).QueryWorkoutData(c)
}

// QueryResults queries the "results" edge of the Competition entity.
func (c *Competition) QueryResults() *CompetitionResultQuery {
	return NewCompetitionClient(c.config).QueryResults(c)
}

// Update returns a builder for updating this Competition.
// Note that you need to call Competition.Unwrap() before calling this method if this Competition
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Competition) Update() *CompetitionUpdateOne {
	return NewCompetitionClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Competition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Competition) Unwrap() *Competition {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Competition is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Competition) String() string {
	var builder strings.Builder
	builder.WriteString("Competition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("public=")
	builder.WriteString(fmt.Sprintf("%v", c.Public))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(c.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("start=")
	builder.WriteString(c.Start.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(c.End.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("scheduled=")
	builder.WriteString(fmt.Sprintf("%v", c.Scheduled))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(c.Status)
	builder.WriteString(", ")
	builder.WriteString("participant_types=")
	builder.WriteString(fmt.Sprintf("%v", c.ParticipantTypes))
	builder.WriteString(", ")
	builder.WriteString("workout_types=")
	builder.WriteString(fmt.Sprintf("%v", c.WorkoutTypes))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(c.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Competitions is a parsable slice of Competition.
type Competitions []*Competition
