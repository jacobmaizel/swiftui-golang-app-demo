// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/squad"
)

// Squad is the model entity for the Squad schema.
type Squad struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SquadQuery when eager-loading is set.
	Edges        SquadEdges `json:"related"`
	squad_owner  *uuid.UUID
	selectValues sql.SelectValues
}

// SquadEdges holds the relations/edges for other nodes in the graph.
type SquadEdges struct {
	// Members holds the value of the members edge.
	Members []*Profile `json:"members,omitempty"`
	// Invites holds the value of the invites edge.
	Invites []*Invite `json:"invites,omitempty"`
	// CompetitionResults holds the value of the competition_results edge.
	CompetitionResults []*CompetitionResult `json:"competition_results,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *Profile `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e SquadEdges) MembersOrErr() ([]*Profile, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// InvitesOrErr returns the Invites value or an error if the edge
// was not loaded in eager-loading.
func (e SquadEdges) InvitesOrErr() ([]*Invite, error) {
	if e.loadedTypes[1] {
		return e.Invites, nil
	}
	return nil, &NotLoadedError{edge: "invites"}
}

// CompetitionResultsOrErr returns the CompetitionResults value or an error if the edge
// was not loaded in eager-loading.
func (e SquadEdges) CompetitionResultsOrErr() ([]*CompetitionResult, error) {
	if e.loadedTypes[2] {
		return e.CompetitionResults, nil
	}
	return nil, &NotLoadedError{edge: "competition_results"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SquadEdges) OwnerOrErr() (*Profile, error) {
	if e.loadedTypes[3] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Squad) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case squad.FieldPublic:
			values[i] = new(sql.NullBool)
		case squad.FieldTitle:
			values[i] = new(sql.NullString)
		case squad.FieldCreatedAt, squad.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case squad.FieldID:
			values[i] = new(uuid.UUID)
		case squad.ForeignKeys[0]: // squad_owner
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Squad fields.
func (s *Squad) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case squad.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case squad.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case squad.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case squad.FieldPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public", values[i])
			} else if value.Valid {
				s.Public = value.Bool
			}
		case squad.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case squad.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field squad_owner", values[i])
			} else if value.Valid {
				s.squad_owner = new(uuid.UUID)
				*s.squad_owner = *value.S.(*uuid.UUID)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Squad.
// This includes values selected through modifiers, order, etc.
func (s *Squad) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryMembers queries the "members" edge of the Squad entity.
func (s *Squad) QueryMembers() *ProfileQuery {
	return NewSquadClient(s.config).QueryMembers(s)
}

// QueryInvites queries the "invites" edge of the Squad entity.
func (s *Squad) QueryInvites() *InviteQuery {
	return NewSquadClient(s.config).QueryInvites(s)
}

// QueryCompetitionResults queries the "competition_results" edge of the Squad entity.
func (s *Squad) QueryCompetitionResults() *CompetitionResultQuery {
	return NewSquadClient(s.config).QueryCompetitionResults(s)
}

// QueryOwner queries the "owner" edge of the Squad entity.
func (s *Squad) QueryOwner() *ProfileQuery {
	return NewSquadClient(s.config).QueryOwner(s)
}

// Update returns a builder for updating this Squad.
// Note that you need to call Squad.Unwrap() before calling this method if this Squad
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Squad) Update() *SquadUpdateOne {
	return NewSquadClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Squad entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Squad) Unwrap() *Squad {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Squad is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Squad) String() string {
	var builder strings.Builder
	builder.WriteString("Squad(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("public=")
	builder.WriteString(fmt.Sprintf("%v", s.Public))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(s.Title)
	builder.WriteByte(')')
	return builder.String()
}

// Squads is a parsable slice of Squad.
type Squads []*Squad
