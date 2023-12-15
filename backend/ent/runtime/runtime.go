// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/post"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/schema"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/squad"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/user"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutroutedata"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	actionMixin := schema.Action{}.Mixin()
	actionMixinFields0 := actionMixin[0].Fields()
	_ = actionMixinFields0
	actionMixinFields2 := actionMixin[2].Fields()
	_ = actionMixinFields2
	actionFields := schema.Action{}.Fields()
	_ = actionFields
	// actionDescCreatedAt is the schema descriptor for created_at field.
	actionDescCreatedAt := actionMixinFields2[0].Descriptor()
	// action.DefaultCreatedAt holds the default value on creation for the created_at field.
	action.DefaultCreatedAt = actionDescCreatedAt.Default.(func() time.Time)
	// actionDescUpdatedAt is the schema descriptor for updated_at field.
	actionDescUpdatedAt := actionMixinFields2[1].Descriptor()
	// action.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	action.DefaultUpdatedAt = actionDescUpdatedAt.Default.(func() time.Time)
	// action.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	action.UpdateDefaultUpdatedAt = actionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// actionDescTitle is the schema descriptor for title field.
	actionDescTitle := actionFields[0].Descriptor()
	// action.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	action.TitleValidator = actionDescTitle.Validators[0].(func(string) error)
	// actionDescID is the schema descriptor for id field.
	actionDescID := actionMixinFields0[0].Descriptor()
	// action.DefaultID holds the default value on creation for the id field.
	action.DefaultID = actionDescID.Default.(func() uuid.UUID)
	appconfigMixin := schema.AppConfig{}.Mixin()
	appconfigMixinFields0 := appconfigMixin[0].Fields()
	_ = appconfigMixinFields0
	appconfigMixinFields2 := appconfigMixin[2].Fields()
	_ = appconfigMixinFields2
	appconfigFields := schema.AppConfig{}.Fields()
	_ = appconfigFields
	// appconfigDescCreatedAt is the schema descriptor for created_at field.
	appconfigDescCreatedAt := appconfigMixinFields2[0].Descriptor()
	// appconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	appconfig.DefaultCreatedAt = appconfigDescCreatedAt.Default.(func() time.Time)
	// appconfigDescUpdatedAt is the schema descriptor for updated_at field.
	appconfigDescUpdatedAt := appconfigMixinFields2[1].Descriptor()
	// appconfig.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appconfig.DefaultUpdatedAt = appconfigDescUpdatedAt.Default.(func() time.Time)
	// appconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appconfig.UpdateDefaultUpdatedAt = appconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appconfigDescAutoSyncWorkouts is the schema descriptor for auto_sync_workouts field.
	appconfigDescAutoSyncWorkouts := appconfigFields[0].Descriptor()
	// appconfig.DefaultAutoSyncWorkouts holds the default value on creation for the auto_sync_workouts field.
	appconfig.DefaultAutoSyncWorkouts = appconfigDescAutoSyncWorkouts.Default.(bool)
	// appconfigDescID is the schema descriptor for id field.
	appconfigDescID := appconfigMixinFields0[0].Descriptor()
	// appconfig.DefaultID holds the default value on creation for the id field.
	appconfig.DefaultID = appconfigDescID.Default.(func() uuid.UUID)
	competitionMixin := schema.Competition{}.Mixin()
	competitionHooks := schema.Competition{}.Hooks()
	competition.Hooks[0] = competitionHooks[0]
	competitionMixinFields0 := competitionMixin[0].Fields()
	_ = competitionMixinFields0
	competitionMixinFields2 := competitionMixin[2].Fields()
	_ = competitionMixinFields2
	competitionMixinFields3 := competitionMixin[3].Fields()
	_ = competitionMixinFields3
	competitionFields := schema.Competition{}.Fields()
	_ = competitionFields
	// competitionDescCreatedAt is the schema descriptor for created_at field.
	competitionDescCreatedAt := competitionMixinFields2[0].Descriptor()
	// competition.DefaultCreatedAt holds the default value on creation for the created_at field.
	competition.DefaultCreatedAt = competitionDescCreatedAt.Default.(func() time.Time)
	// competitionDescUpdatedAt is the schema descriptor for updated_at field.
	competitionDescUpdatedAt := competitionMixinFields2[1].Descriptor()
	// competition.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	competition.DefaultUpdatedAt = competitionDescUpdatedAt.Default.(func() time.Time)
	// competition.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	competition.UpdateDefaultUpdatedAt = competitionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// competitionDescPublic is the schema descriptor for public field.
	competitionDescPublic := competitionMixinFields3[0].Descriptor()
	// competition.DefaultPublic holds the default value on creation for the public field.
	competition.DefaultPublic = competitionDescPublic.Default.(bool)
	// competitionDescScheduled is the schema descriptor for scheduled field.
	competitionDescScheduled := competitionFields[4].Descriptor()
	// competition.DefaultScheduled holds the default value on creation for the scheduled field.
	competition.DefaultScheduled = competitionDescScheduled.Default.(bool)
	// competitionDescStatus is the schema descriptor for status field.
	competitionDescStatus := competitionFields[5].Descriptor()
	// competition.DefaultStatus holds the default value on creation for the status field.
	competition.DefaultStatus = competitionDescStatus.Default.(string)
	// competitionDescType is the schema descriptor for type field.
	competitionDescType := competitionFields[8].Descriptor()
	// competition.DefaultType holds the default value on creation for the type field.
	competition.DefaultType = competitionDescType.Default.(string)
	// competition.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	competition.TypeValidator = competitionDescType.Validators[0].(func(string) error)
	// competitionDescID is the schema descriptor for id field.
	competitionDescID := competitionMixinFields0[0].Descriptor()
	// competition.DefaultID holds the default value on creation for the id field.
	competition.DefaultID = competitionDescID.Default.(func() uuid.UUID)
	competitionresultMixin := schema.CompetitionResult{}.Mixin()
	competitionresultMixinFields0 := competitionresultMixin[0].Fields()
	_ = competitionresultMixinFields0
	competitionresultMixinFields2 := competitionresultMixin[2].Fields()
	_ = competitionresultMixinFields2
	competitionresultFields := schema.CompetitionResult{}.Fields()
	_ = competitionresultFields
	// competitionresultDescCreatedAt is the schema descriptor for created_at field.
	competitionresultDescCreatedAt := competitionresultMixinFields2[0].Descriptor()
	// competitionresult.DefaultCreatedAt holds the default value on creation for the created_at field.
	competitionresult.DefaultCreatedAt = competitionresultDescCreatedAt.Default.(func() time.Time)
	// competitionresultDescUpdatedAt is the schema descriptor for updated_at field.
	competitionresultDescUpdatedAt := competitionresultMixinFields2[1].Descriptor()
	// competitionresult.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	competitionresult.DefaultUpdatedAt = competitionresultDescUpdatedAt.Default.(func() time.Time)
	// competitionresult.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	competitionresult.UpdateDefaultUpdatedAt = competitionresultDescUpdatedAt.UpdateDefault.(func() time.Time)
	// competitionresultDescID is the schema descriptor for id field.
	competitionresultDescID := competitionresultMixinFields0[0].Descriptor()
	// competitionresult.DefaultID holds the default value on creation for the id field.
	competitionresult.DefaultID = competitionresultDescID.Default.(func() uuid.UUID)
	fcmtokenMixin := schema.FcmToken{}.Mixin()
	fcmtokenMixinFields0 := fcmtokenMixin[0].Fields()
	_ = fcmtokenMixinFields0
	fcmtokenMixinFields2 := fcmtokenMixin[2].Fields()
	_ = fcmtokenMixinFields2
	fcmtokenFields := schema.FcmToken{}.Fields()
	_ = fcmtokenFields
	// fcmtokenDescCreatedAt is the schema descriptor for created_at field.
	fcmtokenDescCreatedAt := fcmtokenMixinFields2[0].Descriptor()
	// fcmtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	fcmtoken.DefaultCreatedAt = fcmtokenDescCreatedAt.Default.(func() time.Time)
	// fcmtokenDescUpdatedAt is the schema descriptor for updated_at field.
	fcmtokenDescUpdatedAt := fcmtokenMixinFields2[1].Descriptor()
	// fcmtoken.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	fcmtoken.DefaultUpdatedAt = fcmtokenDescUpdatedAt.Default.(func() time.Time)
	// fcmtoken.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	fcmtoken.UpdateDefaultUpdatedAt = fcmtokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fcmtokenDescID is the schema descriptor for id field.
	fcmtokenDescID := fcmtokenMixinFields0[0].Descriptor()
	// fcmtoken.DefaultID holds the default value on creation for the id field.
	fcmtoken.DefaultID = fcmtokenDescID.Default.(func() uuid.UUID)
	goalMixin := schema.Goal{}.Mixin()
	goalMixinFields0 := goalMixin[0].Fields()
	_ = goalMixinFields0
	goalMixinFields2 := goalMixin[2].Fields()
	_ = goalMixinFields2
	goalFields := schema.Goal{}.Fields()
	_ = goalFields
	// goalDescCreatedAt is the schema descriptor for created_at field.
	goalDescCreatedAt := goalMixinFields2[0].Descriptor()
	// goal.DefaultCreatedAt holds the default value on creation for the created_at field.
	goal.DefaultCreatedAt = goalDescCreatedAt.Default.(func() time.Time)
	// goalDescUpdatedAt is the schema descriptor for updated_at field.
	goalDescUpdatedAt := goalMixinFields2[1].Descriptor()
	// goal.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	goal.DefaultUpdatedAt = goalDescUpdatedAt.Default.(func() time.Time)
	// goal.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	goal.UpdateDefaultUpdatedAt = goalDescUpdatedAt.UpdateDefault.(func() time.Time)
	// goalDescStatus is the schema descriptor for status field.
	goalDescStatus := goalFields[9].Descriptor()
	// goal.DefaultStatus holds the default value on creation for the status field.
	goal.DefaultStatus = goalDescStatus.Default.(string)
	// goalDescID is the schema descriptor for id field.
	goalDescID := goalMixinFields0[0].Descriptor()
	// goal.DefaultID holds the default value on creation for the id field.
	goal.DefaultID = goalDescID.Default.(func() uuid.UUID)
	inviteMixin := schema.Invite{}.Mixin()
	inviteMixinFields0 := inviteMixin[0].Fields()
	_ = inviteMixinFields0
	inviteMixinFields2 := inviteMixin[2].Fields()
	_ = inviteMixinFields2
	inviteFields := schema.Invite{}.Fields()
	_ = inviteFields
	// inviteDescCreatedAt is the schema descriptor for created_at field.
	inviteDescCreatedAt := inviteMixinFields2[0].Descriptor()
	// invite.DefaultCreatedAt holds the default value on creation for the created_at field.
	invite.DefaultCreatedAt = inviteDescCreatedAt.Default.(func() time.Time)
	// inviteDescUpdatedAt is the schema descriptor for updated_at field.
	inviteDescUpdatedAt := inviteMixinFields2[1].Descriptor()
	// invite.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	invite.DefaultUpdatedAt = inviteDescUpdatedAt.Default.(func() time.Time)
	// invite.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	invite.UpdateDefaultUpdatedAt = inviteDescUpdatedAt.UpdateDefault.(func() time.Time)
	// inviteDescStatus is the schema descriptor for status field.
	inviteDescStatus := inviteFields[0].Descriptor()
	// invite.DefaultStatus holds the default value on creation for the status field.
	invite.DefaultStatus = inviteDescStatus.Default.(string)
	// inviteDescID is the schema descriptor for id field.
	inviteDescID := inviteMixinFields0[0].Descriptor()
	// invite.DefaultID holds the default value on creation for the id field.
	invite.DefaultID = inviteDescID.Default.(func() uuid.UUID)
	notificationMixin := schema.Notification{}.Mixin()
	notificationMixinFields0 := notificationMixin[0].Fields()
	_ = notificationMixinFields0
	notificationMixinFields2 := notificationMixin[2].Fields()
	_ = notificationMixinFields2
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationMixinFields2[0].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	// notificationDescUpdatedAt is the schema descriptor for updated_at field.
	notificationDescUpdatedAt := notificationMixinFields2[1].Descriptor()
	// notification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notification.DefaultUpdatedAt = notificationDescUpdatedAt.Default.(func() time.Time)
	// notification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notification.UpdateDefaultUpdatedAt = notificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationDescID is the schema descriptor for id field.
	notificationDescID := notificationMixinFields0[0].Descriptor()
	// notification.DefaultID holds the default value on creation for the id field.
	notification.DefaultID = notificationDescID.Default.(func() uuid.UUID)
	notificationpreferencesMixin := schema.NotificationPreferences{}.Mixin()
	notificationpreferencesMixinFields0 := notificationpreferencesMixin[0].Fields()
	_ = notificationpreferencesMixinFields0
	notificationpreferencesMixinFields2 := notificationpreferencesMixin[2].Fields()
	_ = notificationpreferencesMixinFields2
	notificationpreferencesFields := schema.NotificationPreferences{}.Fields()
	_ = notificationpreferencesFields
	// notificationpreferencesDescCreatedAt is the schema descriptor for created_at field.
	notificationpreferencesDescCreatedAt := notificationpreferencesMixinFields2[0].Descriptor()
	// notificationpreferences.DefaultCreatedAt holds the default value on creation for the created_at field.
	notificationpreferences.DefaultCreatedAt = notificationpreferencesDescCreatedAt.Default.(func() time.Time)
	// notificationpreferencesDescUpdatedAt is the schema descriptor for updated_at field.
	notificationpreferencesDescUpdatedAt := notificationpreferencesMixinFields2[1].Descriptor()
	// notificationpreferences.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notificationpreferences.DefaultUpdatedAt = notificationpreferencesDescUpdatedAt.Default.(func() time.Time)
	// notificationpreferences.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notificationpreferences.UpdateDefaultUpdatedAt = notificationpreferencesDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationpreferencesDescID is the schema descriptor for id field.
	notificationpreferencesDescID := notificationpreferencesMixinFields0[0].Descriptor()
	// notificationpreferences.DefaultID holds the default value on creation for the id field.
	notificationpreferences.DefaultID = notificationpreferencesDescID.Default.(func() uuid.UUID)
	postMixin := schema.Post{}.Mixin()
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postMixinFields2 := postMixin[2].Fields()
	_ = postMixinFields2
	postMixinFields3 := postMixin[3].Fields()
	_ = postMixinFields3
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postMixinFields2[0].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postMixinFields2[1].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	// postDescPublic is the schema descriptor for public field.
	postDescPublic := postMixinFields3[0].Descriptor()
	// post.DefaultPublic holds the default value on creation for the public field.
	post.DefaultPublic = postDescPublic.Default.(bool)
	// postDescID is the schema descriptor for id field.
	postDescID := postMixinFields0[0].Descriptor()
	// post.DefaultID holds the default value on creation for the id field.
	post.DefaultID = postDescID.Default.(func() uuid.UUID)
	profileMixin := schema.Profile{}.Mixin()
	profileHooks := schema.Profile{}.Hooks()
	profile.Hooks[0] = profileHooks[0]
	profileMixinFields0 := profileMixin[0].Fields()
	_ = profileMixinFields0
	profileMixinFields2 := profileMixin[2].Fields()
	_ = profileMixinFields2
	profileMixinFields3 := profileMixin[3].Fields()
	_ = profileMixinFields3
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescCreatedAt is the schema descriptor for created_at field.
	profileDescCreatedAt := profileMixinFields2[0].Descriptor()
	// profile.DefaultCreatedAt holds the default value on creation for the created_at field.
	profile.DefaultCreatedAt = profileDescCreatedAt.Default.(func() time.Time)
	// profileDescUpdatedAt is the schema descriptor for updated_at field.
	profileDescUpdatedAt := profileMixinFields2[1].Descriptor()
	// profile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	profile.DefaultUpdatedAt = profileDescUpdatedAt.Default.(func() time.Time)
	// profile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	profile.UpdateDefaultUpdatedAt = profileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// profileDescPublic is the schema descriptor for public field.
	profileDescPublic := profileMixinFields3[0].Descriptor()
	// profile.DefaultPublic holds the default value on creation for the public field.
	profile.DefaultPublic = profileDescPublic.Default.(bool)
	// profileDescOnboardingCompleted is the schema descriptor for onboarding_completed field.
	profileDescOnboardingCompleted := profileFields[4].Descriptor()
	// profile.DefaultOnboardingCompleted holds the default value on creation for the onboarding_completed field.
	profile.DefaultOnboardingCompleted = profileDescOnboardingCompleted.Default.(bool)
	// profileDescID is the schema descriptor for id field.
	profileDescID := profileMixinFields0[0].Descriptor()
	// profile.DefaultID holds the default value on creation for the id field.
	profile.DefaultID = profileDescID.Default.(func() uuid.UUID)
	squadMixin := schema.Squad{}.Mixin()
	squadHooks := schema.Squad{}.Hooks()
	squad.Hooks[0] = squadHooks[0]
	squadMixinFields0 := squadMixin[0].Fields()
	_ = squadMixinFields0
	squadMixinFields2 := squadMixin[2].Fields()
	_ = squadMixinFields2
	squadMixinFields3 := squadMixin[3].Fields()
	_ = squadMixinFields3
	squadFields := schema.Squad{}.Fields()
	_ = squadFields
	// squadDescCreatedAt is the schema descriptor for created_at field.
	squadDescCreatedAt := squadMixinFields2[0].Descriptor()
	// squad.DefaultCreatedAt holds the default value on creation for the created_at field.
	squad.DefaultCreatedAt = squadDescCreatedAt.Default.(func() time.Time)
	// squadDescUpdatedAt is the schema descriptor for updated_at field.
	squadDescUpdatedAt := squadMixinFields2[1].Descriptor()
	// squad.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	squad.DefaultUpdatedAt = squadDescUpdatedAt.Default.(func() time.Time)
	// squad.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	squad.UpdateDefaultUpdatedAt = squadDescUpdatedAt.UpdateDefault.(func() time.Time)
	// squadDescPublic is the schema descriptor for public field.
	squadDescPublic := squadMixinFields3[0].Descriptor()
	// squad.DefaultPublic holds the default value on creation for the public field.
	squad.DefaultPublic = squadDescPublic.Default.(bool)
	// squadDescID is the schema descriptor for id field.
	squadDescID := squadMixinFields0[0].Descriptor()
	// squad.DefaultID holds the default value on creation for the id field.
	squad.DefaultID = squadDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields2 := userMixin[2].Fields()
	_ = userMixinFields2
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields2[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields2[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	workoutMixin := schema.Workout{}.Mixin()
	workoutMixinFields0 := workoutMixin[0].Fields()
	_ = workoutMixinFields0
	workoutMixinFields2 := workoutMixin[2].Fields()
	_ = workoutMixinFields2
	workoutFields := schema.Workout{}.Fields()
	_ = workoutFields
	// workoutDescCreatedAt is the schema descriptor for created_at field.
	workoutDescCreatedAt := workoutMixinFields2[0].Descriptor()
	// workout.DefaultCreatedAt holds the default value on creation for the created_at field.
	workout.DefaultCreatedAt = workoutDescCreatedAt.Default.(func() time.Time)
	// workoutDescUpdatedAt is the schema descriptor for updated_at field.
	workoutDescUpdatedAt := workoutMixinFields2[1].Descriptor()
	// workout.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	workout.DefaultUpdatedAt = workoutDescUpdatedAt.Default.(func() time.Time)
	// workout.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	workout.UpdateDefaultUpdatedAt = workoutDescUpdatedAt.UpdateDefault.(func() time.Time)
	// workoutDescHealthkitWorkoutActivityType is the schema descriptor for healthkit_workout_activity_type field.
	workoutDescHealthkitWorkoutActivityType := workoutFields[0].Descriptor()
	// workout.HealthkitWorkoutActivityTypeValidator is a validator for the "healthkit_workout_activity_type" field. It is called by the builders before save.
	workout.HealthkitWorkoutActivityTypeValidator = workoutDescHealthkitWorkoutActivityType.Validators[0].(func(string) error)
	// workoutDescID is the schema descriptor for id field.
	workoutDescID := workoutMixinFields0[0].Descriptor()
	// workout.DefaultID holds the default value on creation for the id field.
	workout.DefaultID = workoutDescID.Default.(func() uuid.UUID)
	workoutdataMixin := schema.WorkoutData{}.Mixin()
	workoutdataMixinFields0 := workoutdataMixin[0].Fields()
	_ = workoutdataMixinFields0
	workoutdataMixinFields2 := workoutdataMixin[2].Fields()
	_ = workoutdataMixinFields2
	workoutdataFields := schema.WorkoutData{}.Fields()
	_ = workoutdataFields
	// workoutdataDescCreatedAt is the schema descriptor for created_at field.
	workoutdataDescCreatedAt := workoutdataMixinFields2[0].Descriptor()
	// workoutdata.DefaultCreatedAt holds the default value on creation for the created_at field.
	workoutdata.DefaultCreatedAt = workoutdataDescCreatedAt.Default.(func() time.Time)
	// workoutdataDescUpdatedAt is the schema descriptor for updated_at field.
	workoutdataDescUpdatedAt := workoutdataMixinFields2[1].Descriptor()
	// workoutdata.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	workoutdata.DefaultUpdatedAt = workoutdataDescUpdatedAt.Default.(func() time.Time)
	// workoutdata.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	workoutdata.UpdateDefaultUpdatedAt = workoutdataDescUpdatedAt.UpdateDefault.(func() time.Time)
	// workoutdataDescHealthkitWorkoutStartDate is the schema descriptor for healthkit_workout_start_date field.
	workoutdataDescHealthkitWorkoutStartDate := workoutdataFields[1].Descriptor()
	// workoutdata.DefaultHealthkitWorkoutStartDate holds the default value on creation for the healthkit_workout_start_date field.
	workoutdata.DefaultHealthkitWorkoutStartDate = workoutdataDescHealthkitWorkoutStartDate.Default.(func() time.Time)
	// workoutdataDescHealthkitWorkoutEndDate is the schema descriptor for healthkit_workout_end_date field.
	workoutdataDescHealthkitWorkoutEndDate := workoutdataFields[2].Descriptor()
	// workoutdata.DefaultHealthkitWorkoutEndDate holds the default value on creation for the healthkit_workout_end_date field.
	workoutdata.DefaultHealthkitWorkoutEndDate = workoutdataDescHealthkitWorkoutEndDate.Default.(func() time.Time)
	// workoutdataDescSource is the schema descriptor for source field.
	workoutdataDescSource := workoutdataFields[7].Descriptor()
	// workoutdata.DefaultSource holds the default value on creation for the source field.
	workoutdata.DefaultSource = workoutdataDescSource.Default.(string)
	// workoutdata.SourceValidator is a validator for the "source" field. It is called by the builders before save.
	workoutdata.SourceValidator = workoutdataDescSource.Validators[0].(func(string) error)
	// workoutdataDescLocationType is the schema descriptor for location_type field.
	workoutdataDescLocationType := workoutdataFields[8].Descriptor()
	// workoutdata.DefaultLocationType holds the default value on creation for the location_type field.
	workoutdata.DefaultLocationType = workoutdataDescLocationType.Default.(string)
	// workoutdata.LocationTypeValidator is a validator for the "location_type" field. It is called by the builders before save.
	workoutdata.LocationTypeValidator = workoutdataDescLocationType.Validators[0].(func(string) error)
	// workoutdataDescID is the schema descriptor for id field.
	workoutdataDescID := workoutdataMixinFields0[0].Descriptor()
	// workoutdata.DefaultID holds the default value on creation for the id field.
	workoutdata.DefaultID = workoutdataDescID.Default.(func() uuid.UUID)
	workoutroutedataMixin := schema.WorkoutRouteData{}.Mixin()
	workoutroutedataMixinFields0 := workoutroutedataMixin[0].Fields()
	_ = workoutroutedataMixinFields0
	workoutroutedataMixinFields2 := workoutroutedataMixin[2].Fields()
	_ = workoutroutedataMixinFields2
	workoutroutedataFields := schema.WorkoutRouteData{}.Fields()
	_ = workoutroutedataFields
	// workoutroutedataDescCreatedAt is the schema descriptor for created_at field.
	workoutroutedataDescCreatedAt := workoutroutedataMixinFields2[0].Descriptor()
	// workoutroutedata.DefaultCreatedAt holds the default value on creation for the created_at field.
	workoutroutedata.DefaultCreatedAt = workoutroutedataDescCreatedAt.Default.(func() time.Time)
	// workoutroutedataDescUpdatedAt is the schema descriptor for updated_at field.
	workoutroutedataDescUpdatedAt := workoutroutedataMixinFields2[1].Descriptor()
	// workoutroutedata.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	workoutroutedata.DefaultUpdatedAt = workoutroutedataDescUpdatedAt.Default.(func() time.Time)
	// workoutroutedata.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	workoutroutedata.UpdateDefaultUpdatedAt = workoutroutedataDescUpdatedAt.UpdateDefault.(func() time.Time)
	// workoutroutedataDescID is the schema descriptor for id field.
	workoutroutedataDescID := workoutroutedataMixinFields0[0].Descriptor()
	// workoutroutedata.DefaultID holds the default value on creation for the id field.
	workoutroutedata.DefaultID = workoutroutedataDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.12.4-0.20230726082433-91c7fcc68504"           // Version of ent codegen.
	Sum     = "h1:wUSznEj31LlsQdpc6OvMYNGqugF1s+tY/KpMZsSdonw=" // Sum of ent codegen.
)