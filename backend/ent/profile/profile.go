// Code generated by ent, DO NOT EDIT.

package profile

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the profile type in the database.
	Label = "profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldPicture holds the string denoting the picture field in the database.
	FieldPicture = "picture"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldOnboardingCompleted holds the string denoting the onboarding_completed field in the database.
	FieldOnboardingCompleted = "onboarding_completed"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeFcmTokens holds the string denoting the fcm_tokens edge name in mutations.
	EdgeFcmTokens = "fcm_tokens"
	// EdgeNotificationPreferences holds the string denoting the notification_preferences edge name in mutations.
	EdgeNotificationPreferences = "notification_preferences"
	// EdgeAppConfig holds the string denoting the app_config edge name in mutations.
	EdgeAppConfig = "app_config"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// EdgeCompetitions holds the string denoting the competitions edge name in mutations.
	EdgeCompetitions = "competitions"
	// EdgeActions holds the string denoting the actions edge name in mutations.
	EdgeActions = "actions"
	// EdgeSquad holds the string denoting the squad edge name in mutations.
	EdgeSquad = "squad"
	// EdgeSquadsOwned holds the string denoting the squads_owned edge name in mutations.
	EdgeSquadsOwned = "squads_owned"
	// EdgeInvites holds the string denoting the invites edge name in mutations.
	EdgeInvites = "invites"
	// EdgeInvitesSent holds the string denoting the invites_sent edge name in mutations.
	EdgeInvitesSent = "invites_sent"
	// EdgeWorkouts holds the string denoting the workouts edge name in mutations.
	EdgeWorkouts = "workouts"
	// EdgeGoals holds the string denoting the goals edge name in mutations.
	EdgeGoals = "goals"
	// EdgeWorkoutData holds the string denoting the workout_data edge name in mutations.
	EdgeWorkoutData = "workout_data"
	// EdgeCompetitionResults holds the string denoting the competition_results edge name in mutations.
	EdgeCompetitionResults = "competition_results"
	// Table holds the table name of the profile in the database.
	Table = "profiles"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "profiles"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_profile"
	// FcmTokensTable is the table that holds the fcm_tokens relation/edge.
	FcmTokensTable = "fcm_tokens"
	// FcmTokensInverseTable is the table name for the FcmToken entity.
	// It exists in this package in order to avoid circular dependency with the "fcmtoken" package.
	FcmTokensInverseTable = "fcm_tokens"
	// FcmTokensColumn is the table column denoting the fcm_tokens relation/edge.
	FcmTokensColumn = "profile_fcm_tokens"
	// NotificationPreferencesTable is the table that holds the notification_preferences relation/edge.
	NotificationPreferencesTable = "notification_preferences"
	// NotificationPreferencesInverseTable is the table name for the NotificationPreferences entity.
	// It exists in this package in order to avoid circular dependency with the "notificationpreferences" package.
	NotificationPreferencesInverseTable = "notification_preferences"
	// NotificationPreferencesColumn is the table column denoting the notification_preferences relation/edge.
	NotificationPreferencesColumn = "profile_notification_preferences"
	// AppConfigTable is the table that holds the app_config relation/edge.
	AppConfigTable = "profiles"
	// AppConfigInverseTable is the table name for the AppConfig entity.
	// It exists in this package in order to avoid circular dependency with the "appconfig" package.
	AppConfigInverseTable = "app_configs"
	// AppConfigColumn is the table column denoting the app_config relation/edge.
	AppConfigColumn = "app_config_profile"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "profile_notifications"
	// CompetitionsTable is the table that holds the competitions relation/edge. The primary key declared below.
	CompetitionsTable = "competition_participants"
	// CompetitionsInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	CompetitionsInverseTable = "competitions"
	// ActionsTable is the table that holds the actions relation/edge.
	ActionsTable = "actions"
	// ActionsInverseTable is the table name for the Action entity.
	// It exists in this package in order to avoid circular dependency with the "action" package.
	ActionsInverseTable = "actions"
	// ActionsColumn is the table column denoting the actions relation/edge.
	ActionsColumn = "action_sender"
	// SquadTable is the table that holds the squad relation/edge. The primary key declared below.
	SquadTable = "squad_members"
	// SquadInverseTable is the table name for the Squad entity.
	// It exists in this package in order to avoid circular dependency with the "squad" package.
	SquadInverseTable = "squads"
	// SquadsOwnedTable is the table that holds the squads_owned relation/edge.
	SquadsOwnedTable = "squads"
	// SquadsOwnedInverseTable is the table name for the Squad entity.
	// It exists in this package in order to avoid circular dependency with the "squad" package.
	SquadsOwnedInverseTable = "squads"
	// SquadsOwnedColumn is the table column denoting the squads_owned relation/edge.
	SquadsOwnedColumn = "squad_owner"
	// InvitesTable is the table that holds the invites relation/edge.
	InvitesTable = "invites"
	// InvitesInverseTable is the table name for the Invite entity.
	// It exists in this package in order to avoid circular dependency with the "invite" package.
	InvitesInverseTable = "invites"
	// InvitesColumn is the table column denoting the invites relation/edge.
	InvitesColumn = "invite_receiver"
	// InvitesSentTable is the table that holds the invites_sent relation/edge.
	InvitesSentTable = "invites"
	// InvitesSentInverseTable is the table name for the Invite entity.
	// It exists in this package in order to avoid circular dependency with the "invite" package.
	InvitesSentInverseTable = "invites"
	// InvitesSentColumn is the table column denoting the invites_sent relation/edge.
	InvitesSentColumn = "invite_sender"
	// WorkoutsTable is the table that holds the workouts relation/edge.
	WorkoutsTable = "workouts"
	// WorkoutsInverseTable is the table name for the Workout entity.
	// It exists in this package in order to avoid circular dependency with the "workout" package.
	WorkoutsInverseTable = "workouts"
	// WorkoutsColumn is the table column denoting the workouts relation/edge.
	WorkoutsColumn = "workout_leader"
	// GoalsTable is the table that holds the goals relation/edge.
	GoalsTable = "goals"
	// GoalsInverseTable is the table name for the Goal entity.
	// It exists in this package in order to avoid circular dependency with the "goal" package.
	GoalsInverseTable = "goals"
	// GoalsColumn is the table column denoting the goals relation/edge.
	GoalsColumn = "goal_profile"
	// WorkoutDataTable is the table that holds the workout_data relation/edge.
	WorkoutDataTable = "workout_data"
	// WorkoutDataInverseTable is the table name for the WorkoutData entity.
	// It exists in this package in order to avoid circular dependency with the "workoutdata" package.
	WorkoutDataInverseTable = "workout_data"
	// WorkoutDataColumn is the table column denoting the workout_data relation/edge.
	WorkoutDataColumn = "workout_data_profile"
	// CompetitionResultsTable is the table that holds the competition_results relation/edge.
	CompetitionResultsTable = "competition_results"
	// CompetitionResultsInverseTable is the table name for the CompetitionResult entity.
	// It exists in this package in order to avoid circular dependency with the "competitionresult" package.
	CompetitionResultsInverseTable = "competition_results"
	// CompetitionResultsColumn is the table column denoting the competition_results relation/edge.
	CompetitionResultsColumn = "competition_result_profile"
)

// Columns holds all SQL columns for profile fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldPublic,
	FieldFirstName,
	FieldLastName,
	FieldPicture,
	FieldBirthday,
	FieldOnboardingCompleted,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "profiles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"app_config_profile",
	"user_profile",
}

var (
	// CompetitionsPrimaryKey and CompetitionsColumn2 are the table columns denoting the
	// primary key for the competitions relation (M2M).
	CompetitionsPrimaryKey = []string{"competition_id", "profile_id"}
	// SquadPrimaryKey and SquadColumn2 are the table columns denoting the
	// primary key for the squad relation (M2M).
	SquadPrimaryKey = []string{"squad_id", "profile_id"}
)

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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/jacobmaizel/swiftui-golang-app-demo/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultPublic holds the default value on creation for the "public" field.
	DefaultPublic bool
	// DefaultOnboardingCompleted holds the default value on creation for the "onboarding_completed" field.
	DefaultOnboardingCompleted bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Profile queries.
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

// ByPublic orders the results by the public field.
func ByPublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublic, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByPicture orders the results by the picture field.
func ByPicture(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPicture, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByOnboardingCompleted orders the results by the onboarding_completed field.
func ByOnboardingCompleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOnboardingCompleted, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByFcmTokensCount orders the results by fcm_tokens count.
func ByFcmTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFcmTokensStep(), opts...)
	}
}

// ByFcmTokens orders the results by fcm_tokens terms.
func ByFcmTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFcmTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotificationPreferencesField orders the results by notification_preferences field.
func ByNotificationPreferencesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationPreferencesStep(), sql.OrderByField(field, opts...))
	}
}

// ByAppConfigField orders the results by app_config field.
func ByAppConfigField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppConfigStep(), sql.OrderByField(field, opts...))
	}
}

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompetitionsCount orders the results by competitions count.
func ByCompetitionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCompetitionsStep(), opts...)
	}
}

// ByCompetitions orders the results by competitions terms.
func ByCompetitions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompetitionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByActionsCount orders the results by actions count.
func ByActionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newActionsStep(), opts...)
	}
}

// ByActions orders the results by actions terms.
func ByActions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySquadCount orders the results by squad count.
func BySquadCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSquadStep(), opts...)
	}
}

// BySquad orders the results by squad terms.
func BySquad(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSquadStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySquadsOwnedCount orders the results by squads_owned count.
func BySquadsOwnedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSquadsOwnedStep(), opts...)
	}
}

// BySquadsOwned orders the results by squads_owned terms.
func BySquadsOwned(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSquadsOwnedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInvitesCount orders the results by invites count.
func ByInvitesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInvitesStep(), opts...)
	}
}

// ByInvites orders the results by invites terms.
func ByInvites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvitesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInvitesSentCount orders the results by invites_sent count.
func ByInvitesSentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInvitesSentStep(), opts...)
	}
}

// ByInvitesSent orders the results by invites_sent terms.
func ByInvitesSent(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvitesSentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkoutsCount orders the results by workouts count.
func ByWorkoutsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkoutsStep(), opts...)
	}
}

// ByWorkouts orders the results by workouts terms.
func ByWorkouts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkoutsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGoalsCount orders the results by goals count.
func ByGoalsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGoalsStep(), opts...)
	}
}

// ByGoals orders the results by goals terms.
func ByGoals(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGoalsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByCompetitionResultsCount orders the results by competition_results count.
func ByCompetitionResultsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCompetitionResultsStep(), opts...)
	}
}

// ByCompetitionResults orders the results by competition_results terms.
func ByCompetitionResults(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompetitionResultsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
func newFcmTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FcmTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FcmTokensTable, FcmTokensColumn),
	)
}
func newNotificationPreferencesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationPreferencesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, NotificationPreferencesTable, NotificationPreferencesColumn),
	)
}
func newAppConfigStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppConfigInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, AppConfigTable, AppConfigColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
	)
}
func newCompetitionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompetitionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CompetitionsTable, CompetitionsPrimaryKey...),
	)
}
func newActionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ActionsTable, ActionsColumn),
	)
}
func newSquadStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SquadInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SquadTable, SquadPrimaryKey...),
	)
}
func newSquadsOwnedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SquadsOwnedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, SquadsOwnedTable, SquadsOwnedColumn),
	)
}
func newInvitesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvitesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, InvitesTable, InvitesColumn),
	)
}
func newInvitesSentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvitesSentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, InvitesSentTable, InvitesSentColumn),
	)
}
func newWorkoutsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkoutsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, WorkoutsTable, WorkoutsColumn),
	)
}
func newGoalsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GoalsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, GoalsTable, GoalsColumn),
	)
}
func newWorkoutDataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkoutDataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, WorkoutDataTable, WorkoutDataColumn),
	)
}
func newCompetitionResultsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompetitionResultsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CompetitionResultsTable, CompetitionResultsColumn),
	)
}
