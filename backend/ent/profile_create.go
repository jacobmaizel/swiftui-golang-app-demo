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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/action"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/appconfig"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competitionresult"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/fcmtoken"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/goal"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/invite"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notification"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notificationpreferences"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/squad"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/user"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProfileCreate) SetCreatedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreatedAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProfileCreate) SetUpdatedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdatedAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetPublic sets the "public" field.
func (pc *ProfileCreate) SetPublic(b bool) *ProfileCreate {
	pc.mutation.SetPublic(b)
	return pc
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (pc *ProfileCreate) SetNillablePublic(b *bool) *ProfileCreate {
	if b != nil {
		pc.SetPublic(*b)
	}
	return pc
}

// SetFirstName sets the "first_name" field.
func (pc *ProfileCreate) SetFirstName(s string) *ProfileCreate {
	pc.mutation.SetFirstName(s)
	return pc
}

// SetLastName sets the "last_name" field.
func (pc *ProfileCreate) SetLastName(s string) *ProfileCreate {
	pc.mutation.SetLastName(s)
	return pc
}

// SetPicture sets the "picture" field.
func (pc *ProfileCreate) SetPicture(s string) *ProfileCreate {
	pc.mutation.SetPicture(s)
	return pc
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (pc *ProfileCreate) SetNillablePicture(s *string) *ProfileCreate {
	if s != nil {
		pc.SetPicture(*s)
	}
	return pc
}

// SetBirthday sets the "birthday" field.
func (pc *ProfileCreate) SetBirthday(t time.Time) *ProfileCreate {
	pc.mutation.SetBirthday(t)
	return pc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableBirthday(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetBirthday(*t)
	}
	return pc
}

// SetOnboardingCompleted sets the "onboarding_completed" field.
func (pc *ProfileCreate) SetOnboardingCompleted(b bool) *ProfileCreate {
	pc.mutation.SetOnboardingCompleted(b)
	return pc
}

// SetNillableOnboardingCompleted sets the "onboarding_completed" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableOnboardingCompleted(b *bool) *ProfileCreate {
	if b != nil {
		pc.SetOnboardingCompleted(*b)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProfileCreate) SetID(u uuid.UUID) *ProfileCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableID(u *uuid.UUID) *ProfileCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *ProfileCreate) SetUserID(id uuid.UUID) *ProfileCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *ProfileCreate) SetUser(u *User) *ProfileCreate {
	return pc.SetUserID(u.ID)
}

// AddFcmTokenIDs adds the "fcm_tokens" edge to the FcmToken entity by IDs.
func (pc *ProfileCreate) AddFcmTokenIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddFcmTokenIDs(ids...)
	return pc
}

// AddFcmTokens adds the "fcm_tokens" edges to the FcmToken entity.
func (pc *ProfileCreate) AddFcmTokens(f ...*FcmToken) *ProfileCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pc.AddFcmTokenIDs(ids...)
}

// SetNotificationPreferencesID sets the "notification_preferences" edge to the NotificationPreferences entity by ID.
func (pc *ProfileCreate) SetNotificationPreferencesID(id uuid.UUID) *ProfileCreate {
	pc.mutation.SetNotificationPreferencesID(id)
	return pc
}

// SetNillableNotificationPreferencesID sets the "notification_preferences" edge to the NotificationPreferences entity by ID if the given value is not nil.
func (pc *ProfileCreate) SetNillableNotificationPreferencesID(id *uuid.UUID) *ProfileCreate {
	if id != nil {
		pc = pc.SetNotificationPreferencesID(*id)
	}
	return pc
}

// SetNotificationPreferences sets the "notification_preferences" edge to the NotificationPreferences entity.
func (pc *ProfileCreate) SetNotificationPreferences(n *NotificationPreferences) *ProfileCreate {
	return pc.SetNotificationPreferencesID(n.ID)
}

// SetAppConfigID sets the "app_config" edge to the AppConfig entity by ID.
func (pc *ProfileCreate) SetAppConfigID(id uuid.UUID) *ProfileCreate {
	pc.mutation.SetAppConfigID(id)
	return pc
}

// SetNillableAppConfigID sets the "app_config" edge to the AppConfig entity by ID if the given value is not nil.
func (pc *ProfileCreate) SetNillableAppConfigID(id *uuid.UUID) *ProfileCreate {
	if id != nil {
		pc = pc.SetAppConfigID(*id)
	}
	return pc
}

// SetAppConfig sets the "app_config" edge to the AppConfig entity.
func (pc *ProfileCreate) SetAppConfig(a *AppConfig) *ProfileCreate {
	return pc.SetAppConfigID(a.ID)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (pc *ProfileCreate) AddNotificationIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddNotificationIDs(ids...)
	return pc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (pc *ProfileCreate) AddNotifications(n ...*Notification) *ProfileCreate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pc.AddNotificationIDs(ids...)
}

// AddCompetitionIDs adds the "competitions" edge to the Competition entity by IDs.
func (pc *ProfileCreate) AddCompetitionIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddCompetitionIDs(ids...)
	return pc
}

// AddCompetitions adds the "competitions" edges to the Competition entity.
func (pc *ProfileCreate) AddCompetitions(c ...*Competition) *ProfileCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCompetitionIDs(ids...)
}

// AddActionIDs adds the "actions" edge to the Action entity by IDs.
func (pc *ProfileCreate) AddActionIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddActionIDs(ids...)
	return pc
}

// AddActions adds the "actions" edges to the Action entity.
func (pc *ProfileCreate) AddActions(a ...*Action) *ProfileCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddActionIDs(ids...)
}

// AddSquadIDs adds the "squad" edge to the Squad entity by IDs.
func (pc *ProfileCreate) AddSquadIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddSquadIDs(ids...)
	return pc
}

// AddSquad adds the "squad" edges to the Squad entity.
func (pc *ProfileCreate) AddSquad(s ...*Squad) *ProfileCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSquadIDs(ids...)
}

// AddSquadsOwnedIDs adds the "squads_owned" edge to the Squad entity by IDs.
func (pc *ProfileCreate) AddSquadsOwnedIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddSquadsOwnedIDs(ids...)
	return pc
}

// AddSquadsOwned adds the "squads_owned" edges to the Squad entity.
func (pc *ProfileCreate) AddSquadsOwned(s ...*Squad) *ProfileCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSquadsOwnedIDs(ids...)
}

// AddInviteIDs adds the "invites" edge to the Invite entity by IDs.
func (pc *ProfileCreate) AddInviteIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddInviteIDs(ids...)
	return pc
}

// AddInvites adds the "invites" edges to the Invite entity.
func (pc *ProfileCreate) AddInvites(i ...*Invite) *ProfileCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pc.AddInviteIDs(ids...)
}

// AddInvitesSentIDs adds the "invites_sent" edge to the Invite entity by IDs.
func (pc *ProfileCreate) AddInvitesSentIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddInvitesSentIDs(ids...)
	return pc
}

// AddInvitesSent adds the "invites_sent" edges to the Invite entity.
func (pc *ProfileCreate) AddInvitesSent(i ...*Invite) *ProfileCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pc.AddInvitesSentIDs(ids...)
}

// AddWorkoutIDs adds the "workouts" edge to the Workout entity by IDs.
func (pc *ProfileCreate) AddWorkoutIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddWorkoutIDs(ids...)
	return pc
}

// AddWorkouts adds the "workouts" edges to the Workout entity.
func (pc *ProfileCreate) AddWorkouts(w ...*Workout) *ProfileCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pc.AddWorkoutIDs(ids...)
}

// AddGoalIDs adds the "goals" edge to the Goal entity by IDs.
func (pc *ProfileCreate) AddGoalIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddGoalIDs(ids...)
	return pc
}

// AddGoals adds the "goals" edges to the Goal entity.
func (pc *ProfileCreate) AddGoals(g ...*Goal) *ProfileCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return pc.AddGoalIDs(ids...)
}

// AddWorkoutDatumIDs adds the "workout_data" edge to the WorkoutData entity by IDs.
func (pc *ProfileCreate) AddWorkoutDatumIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddWorkoutDatumIDs(ids...)
	return pc
}

// AddWorkoutData adds the "workout_data" edges to the WorkoutData entity.
func (pc *ProfileCreate) AddWorkoutData(w ...*WorkoutData) *ProfileCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pc.AddWorkoutDatumIDs(ids...)
}

// AddCompetitionResultIDs adds the "competition_results" edge to the CompetitionResult entity by IDs.
func (pc *ProfileCreate) AddCompetitionResultIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddCompetitionResultIDs(ids...)
	return pc
}

// AddCompetitionResults adds the "competition_results" edges to the CompetitionResult entity.
func (pc *ProfileCreate) AddCompetitionResults(c ...*CompetitionResult) *ProfileCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCompetitionResultIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfileCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfileCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfileCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if profile.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized profile.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := profile.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if profile.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized profile.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := profile.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Public(); !ok {
		v := profile.DefaultPublic
		pc.mutation.SetPublic(v)
	}
	if _, ok := pc.mutation.OnboardingCompleted(); !ok {
		v := profile.DefaultOnboardingCompleted
		pc.mutation.SetOnboardingCompleted(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		if profile.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized profile.DefaultID (forgotten import ent/runtime?)")
		}
		v := profile.DefaultID()
		pc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfileCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Profile.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Profile.updated_at"`)}
	}
	if _, ok := pc.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "Profile.public"`)}
	}
	if _, ok := pc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Profile.first_name"`)}
	}
	if _, ok := pc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Profile.last_name"`)}
	}
	if _, ok := pc.mutation.OnboardingCompleted(); !ok {
		return &ValidationError{Name: "onboarding_completed", err: errors.New(`ent: missing required field "Profile.onboarding_completed"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Profile.user"`)}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		_node = &Profile{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(profile.Table, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(profile.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(profile.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Public(); ok {
		_spec.SetField(profile.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if value, ok := pc.mutation.FirstName(); ok {
		_spec.SetField(profile.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := pc.mutation.LastName(); ok {
		_spec.SetField(profile.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := pc.mutation.Picture(); ok {
		_spec.SetField(profile.FieldPicture, field.TypeString, value)
		_node.Picture = value
	}
	if value, ok := pc.mutation.Birthday(); ok {
		_spec.SetField(profile.FieldBirthday, field.TypeTime, value)
		_node.Birthday = &value
	}
	if value, ok := pc.mutation.OnboardingCompleted(); ok {
		_spec.SetField(profile.FieldOnboardingCompleted, field.TypeBool, value)
		_node.OnboardingCompleted = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.FcmTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.FcmTokensTable,
			Columns: []string{profile.FcmTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fcmtoken.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.NotificationPreferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profile.NotificationPreferencesTable,
			Columns: []string{profile.NotificationPreferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationpreferences.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.AppConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.AppConfigTable,
			Columns: []string{profile.AppConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appconfig.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.app_config_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.NotificationsTable,
			Columns: []string{profile.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CompetitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   profile.CompetitionsTable,
			Columns: profile.CompetitionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competition.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ActionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.ActionsTable,
			Columns: []string{profile.ActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(action.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SquadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   profile.SquadTable,
			Columns: profile.SquadPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SquadsOwnedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.SquadsOwnedTable,
			Columns: []string{profile.SquadsOwnedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.InvitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.InvitesTable,
			Columns: []string{profile.InvitesColumn},
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
	if nodes := pc.mutation.InvitesSentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.InvitesSentTable,
			Columns: []string{profile.InvitesSentColumn},
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
	if nodes := pc.mutation.WorkoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.WorkoutsTable,
			Columns: []string{profile.WorkoutsColumn},
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
	if nodes := pc.mutation.GoalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.GoalsTable,
			Columns: []string{profile.GoalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(goal.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.WorkoutDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.WorkoutDataTable,
			Columns: []string{profile.WorkoutDataColumn},
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
	if nodes := pc.mutation.CompetitionResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.CompetitionResultsTable,
			Columns: []string{profile.CompetitionResultsColumn},
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

// ProfileCreateBulk is the builder for creating many Profile entities in bulk.
type ProfileCreateBulk struct {
	config
	builders []*ProfileCreate
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfileCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
