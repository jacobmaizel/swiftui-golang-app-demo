// Code generated by ent, DO NOT EDIT.

package invite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the invite type in the database.
	Label = "invite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeSender holds the string denoting the sender edge name in mutations.
	EdgeSender = "sender"
	// EdgeReceiver holds the string denoting the receiver edge name in mutations.
	EdgeReceiver = "receiver"
	// EdgeSquad holds the string denoting the squad edge name in mutations.
	EdgeSquad = "squad"
	// EdgeCompetition holds the string denoting the competition edge name in mutations.
	EdgeCompetition = "competition"
	// EdgeWorkout holds the string denoting the workout edge name in mutations.
	EdgeWorkout = "workout"
	// Table holds the table name of the invite in the database.
	Table = "invites"
	// SenderTable is the table that holds the sender relation/edge.
	SenderTable = "invites"
	// SenderInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	SenderInverseTable = "profiles"
	// SenderColumn is the table column denoting the sender relation/edge.
	SenderColumn = "invite_sender"
	// ReceiverTable is the table that holds the receiver relation/edge.
	ReceiverTable = "invites"
	// ReceiverInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ReceiverInverseTable = "profiles"
	// ReceiverColumn is the table column denoting the receiver relation/edge.
	ReceiverColumn = "invite_receiver"
	// SquadTable is the table that holds the squad relation/edge.
	SquadTable = "invites"
	// SquadInverseTable is the table name for the Squad entity.
	// It exists in this package in order to avoid circular dependency with the "squad" package.
	SquadInverseTable = "squads"
	// SquadColumn is the table column denoting the squad relation/edge.
	SquadColumn = "invite_squad"
	// CompetitionTable is the table that holds the competition relation/edge.
	CompetitionTable = "invites"
	// CompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	CompetitionInverseTable = "competitions"
	// CompetitionColumn is the table column denoting the competition relation/edge.
	CompetitionColumn = "invite_competition"
	// WorkoutTable is the table that holds the workout relation/edge.
	WorkoutTable = "invites"
	// WorkoutInverseTable is the table name for the Workout entity.
	// It exists in this package in order to avoid circular dependency with the "workout" package.
	WorkoutInverseTable = "workouts"
	// WorkoutColumn is the table column denoting the workout relation/edge.
	WorkoutColumn = "invite_workout"
)

// Columns holds all SQL columns for invite fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "invites"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"invite_sender",
	"invite_receiver",
	"invite_squad",
	"invite_competition",
	"invite_workout",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Invite queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySenderField orders the results by sender field.
func BySenderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenderStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceiverField orders the results by receiver field.
func ByReceiverField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiverStep(), sql.OrderByField(field, opts...))
	}
}

// BySquadField orders the results by squad field.
func BySquadField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSquadStep(), sql.OrderByField(field, opts...))
	}
}

// ByCompetitionField orders the results by competition field.
func ByCompetitionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompetitionStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkoutField orders the results by workout field.
func ByWorkoutField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkoutStep(), sql.OrderByField(field, opts...))
	}
}
func newSenderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SenderTable, SenderColumn),
	)
}
func newReceiverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ReceiverTable, ReceiverColumn),
	)
}
func newSquadStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SquadInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SquadTable, SquadColumn),
	)
}
func newCompetitionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompetitionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CompetitionTable, CompetitionColumn),
	)
}
func newWorkoutStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkoutInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, WorkoutTable, WorkoutColumn),
	)
}