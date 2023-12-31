// Code generated by ent, DO NOT EDIT.

package workout

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the workout type in the database.
	Label = "workout"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldHealthkitWorkoutActivityType holds the string denoting the healthkit_workout_activity_type field in the database.
	FieldHealthkitWorkoutActivityType = "healthkit_workout_activity_type"
	// EdgeInvite holds the string denoting the invite edge name in mutations.
	EdgeInvite = "invite"
	// EdgeLeader holds the string denoting the leader edge name in mutations.
	EdgeLeader = "leader"
	// EdgeCompetition holds the string denoting the competition edge name in mutations.
	EdgeCompetition = "competition"
	// EdgeWorkoutData holds the string denoting the workout_data edge name in mutations.
	EdgeWorkoutData = "workout_data"
	// EdgeWorkoutRouteData holds the string denoting the workout_route_data edge name in mutations.
	EdgeWorkoutRouteData = "workout_route_data"
	// Table holds the table name of the workout in the database.
	Table = "workouts"
	// InviteTable is the table that holds the invite relation/edge.
	InviteTable = "invites"
	// InviteInverseTable is the table name for the Invite entity.
	// It exists in this package in order to avoid circular dependency with the "invite" package.
	InviteInverseTable = "invites"
	// InviteColumn is the table column denoting the invite relation/edge.
	InviteColumn = "invite_workout"
	// LeaderTable is the table that holds the leader relation/edge.
	LeaderTable = "workouts"
	// LeaderInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	LeaderInverseTable = "profiles"
	// LeaderColumn is the table column denoting the leader relation/edge.
	LeaderColumn = "workout_leader"
	// CompetitionTable is the table that holds the competition relation/edge.
	CompetitionTable = "workouts"
	// CompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	CompetitionInverseTable = "competitions"
	// CompetitionColumn is the table column denoting the competition relation/edge.
	CompetitionColumn = "workout_competition"
	// WorkoutDataTable is the table that holds the workout_data relation/edge.
	WorkoutDataTable = "workout_data"
	// WorkoutDataInverseTable is the table name for the WorkoutData entity.
	// It exists in this package in order to avoid circular dependency with the "workoutdata" package.
	WorkoutDataInverseTable = "workout_data"
	// WorkoutDataColumn is the table column denoting the workout_data relation/edge.
	WorkoutDataColumn = "workout_data_workout"
	// WorkoutRouteDataTable is the table that holds the workout_route_data relation/edge.
	WorkoutRouteDataTable = "workout_route_data"
	// WorkoutRouteDataInverseTable is the table name for the WorkoutRouteData entity.
	// It exists in this package in order to avoid circular dependency with the "workoutroutedata" package.
	WorkoutRouteDataInverseTable = "workout_route_data"
	// WorkoutRouteDataColumn is the table column denoting the workout_route_data relation/edge.
	WorkoutRouteDataColumn = "workout_route_data_workout"
)

// Columns holds all SQL columns for workout fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldHealthkitWorkoutActivityType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "workouts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"workout_leader",
	"workout_competition",
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
	// HealthkitWorkoutActivityTypeValidator is a validator for the "healthkit_workout_activity_type" field. It is called by the builders before save.
	HealthkitWorkoutActivityTypeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Workout queries.
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

// ByHealthkitWorkoutActivityType orders the results by the healthkit_workout_activity_type field.
func ByHealthkitWorkoutActivityType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHealthkitWorkoutActivityType, opts...).ToFunc()
}

// ByInviteCount orders the results by invite count.
func ByInviteCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInviteStep(), opts...)
	}
}

// ByInvite orders the results by invite terms.
func ByInvite(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInviteStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLeaderField orders the results by leader field.
func ByLeaderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLeaderStep(), sql.OrderByField(field, opts...))
	}
}

// ByCompetitionField orders the results by competition field.
func ByCompetitionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompetitionStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkoutDataCount orders the results by workout_data count.
func ByWorkoutDataCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkoutDataStep(), opts...)
	}
}

// ByWorkoutData orders the results by workout_data terms.
func ByWorkoutData(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkoutDataStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkoutRouteDataCount orders the results by workout_route_data count.
func ByWorkoutRouteDataCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkoutRouteDataStep(), opts...)
	}
}

// ByWorkoutRouteData orders the results by workout_route_data terms.
func ByWorkoutRouteData(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkoutRouteDataStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newInviteStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InviteInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, InviteTable, InviteColumn),
	)
}
func newLeaderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LeaderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, LeaderTable, LeaderColumn),
	)
}
func newCompetitionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompetitionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CompetitionTable, CompetitionColumn),
	)
}
func newWorkoutDataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkoutDataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, WorkoutDataTable, WorkoutDataColumn),
	)
}
func newWorkoutRouteDataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkoutRouteDataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, WorkoutRouteDataTable, WorkoutRouteDataColumn),
	)
}
