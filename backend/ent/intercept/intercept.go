// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/squad"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/user"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutroutedata"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The ActionFunc type is an adapter to allow the use of ordinary function as a Querier.
type ActionFunc func(context.Context, *ent.ActionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ActionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ActionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ActionQuery", q)
}

// The TraverseAction type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAction func(context.Context, *ent.ActionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAction) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAction) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ActionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ActionQuery", q)
}

// The AppConfigFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppConfigFunc func(context.Context, *ent.AppConfigQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AppConfigFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AppConfigQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AppConfigQuery", q)
}

// The TraverseAppConfig type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAppConfig func(context.Context, *ent.AppConfigQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAppConfig) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAppConfig) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppConfigQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AppConfigQuery", q)
}

// The CompetitionFunc type is an adapter to allow the use of ordinary function as a Querier.
type CompetitionFunc func(context.Context, *ent.CompetitionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f CompetitionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.CompetitionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.CompetitionQuery", q)
}

// The TraverseCompetition type is an adapter to allow the use of ordinary function as Traverser.
type TraverseCompetition func(context.Context, *ent.CompetitionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseCompetition) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseCompetition) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CompetitionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.CompetitionQuery", q)
}

// The CompetitionResultFunc type is an adapter to allow the use of ordinary function as a Querier.
type CompetitionResultFunc func(context.Context, *ent.CompetitionResultQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f CompetitionResultFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.CompetitionResultQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.CompetitionResultQuery", q)
}

// The TraverseCompetitionResult type is an adapter to allow the use of ordinary function as Traverser.
type TraverseCompetitionResult func(context.Context, *ent.CompetitionResultQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseCompetitionResult) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseCompetitionResult) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CompetitionResultQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.CompetitionResultQuery", q)
}

// The FcmTokenFunc type is an adapter to allow the use of ordinary function as a Querier.
type FcmTokenFunc func(context.Context, *ent.FcmTokenQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f FcmTokenFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.FcmTokenQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.FcmTokenQuery", q)
}

// The TraverseFcmToken type is an adapter to allow the use of ordinary function as Traverser.
type TraverseFcmToken func(context.Context, *ent.FcmTokenQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFcmToken) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFcmToken) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FcmTokenQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.FcmTokenQuery", q)
}

// The GoalFunc type is an adapter to allow the use of ordinary function as a Querier.
type GoalFunc func(context.Context, *ent.GoalQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f GoalFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.GoalQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.GoalQuery", q)
}

// The TraverseGoal type is an adapter to allow the use of ordinary function as Traverser.
type TraverseGoal func(context.Context, *ent.GoalQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseGoal) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseGoal) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoalQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.GoalQuery", q)
}

// The InviteFunc type is an adapter to allow the use of ordinary function as a Querier.
type InviteFunc func(context.Context, *ent.InviteQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f InviteFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.InviteQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.InviteQuery", q)
}

// The TraverseInvite type is an adapter to allow the use of ordinary function as Traverser.
type TraverseInvite func(context.Context, *ent.InviteQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseInvite) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseInvite) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InviteQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.InviteQuery", q)
}

// The NotificationFunc type is an adapter to allow the use of ordinary function as a Querier.
type NotificationFunc func(context.Context, *ent.NotificationQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f NotificationFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.NotificationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.NotificationQuery", q)
}

// The TraverseNotification type is an adapter to allow the use of ordinary function as Traverser.
type TraverseNotification func(context.Context, *ent.NotificationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseNotification) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseNotification) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotificationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.NotificationQuery", q)
}

// The NotificationPreferencesFunc type is an adapter to allow the use of ordinary function as a Querier.
type NotificationPreferencesFunc func(context.Context, *ent.NotificationPreferencesQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f NotificationPreferencesFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.NotificationPreferencesQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.NotificationPreferencesQuery", q)
}

// The TraverseNotificationPreferences type is an adapter to allow the use of ordinary function as Traverser.
type TraverseNotificationPreferences func(context.Context, *ent.NotificationPreferencesQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseNotificationPreferences) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseNotificationPreferences) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotificationPreferencesQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.NotificationPreferencesQuery", q)
}

// The PostFunc type is an adapter to allow the use of ordinary function as a Querier.
type PostFunc func(context.Context, *ent.PostQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PostFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PostQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PostQuery", q)
}

// The TraversePost type is an adapter to allow the use of ordinary function as Traverser.
type TraversePost func(context.Context, *ent.PostQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePost) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePost) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PostQuery", q)
}

// The ProfileFunc type is an adapter to allow the use of ordinary function as a Querier.
type ProfileFunc func(context.Context, *ent.ProfileQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ProfileFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ProfileQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ProfileQuery", q)
}

// The TraverseProfile type is an adapter to allow the use of ordinary function as Traverser.
type TraverseProfile func(context.Context, *ent.ProfileQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseProfile) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseProfile) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProfileQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ProfileQuery", q)
}

// The SquadFunc type is an adapter to allow the use of ordinary function as a Querier.
type SquadFunc func(context.Context, *ent.SquadQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f SquadFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.SquadQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.SquadQuery", q)
}

// The TraverseSquad type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSquad func(context.Context, *ent.SquadQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSquad) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSquad) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SquadQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.SquadQuery", q)
}

// The UserFunc type is an adapter to allow the use of ordinary function as a Querier.
type UserFunc func(context.Context, *ent.UserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f UserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// The TraverseUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseUser func(context.Context, *ent.UserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// The WorkoutFunc type is an adapter to allow the use of ordinary function as a Querier.
type WorkoutFunc func(context.Context, *ent.WorkoutQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f WorkoutFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.WorkoutQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.WorkoutQuery", q)
}

// The TraverseWorkout type is an adapter to allow the use of ordinary function as Traverser.
type TraverseWorkout func(context.Context, *ent.WorkoutQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseWorkout) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseWorkout) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WorkoutQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.WorkoutQuery", q)
}

// The WorkoutDataFunc type is an adapter to allow the use of ordinary function as a Querier.
type WorkoutDataFunc func(context.Context, *ent.WorkoutDataQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f WorkoutDataFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.WorkoutDataQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.WorkoutDataQuery", q)
}

// The TraverseWorkoutData type is an adapter to allow the use of ordinary function as Traverser.
type TraverseWorkoutData func(context.Context, *ent.WorkoutDataQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseWorkoutData) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseWorkoutData) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WorkoutDataQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.WorkoutDataQuery", q)
}

// The WorkoutRouteDataFunc type is an adapter to allow the use of ordinary function as a Querier.
type WorkoutRouteDataFunc func(context.Context, *ent.WorkoutRouteDataQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f WorkoutRouteDataFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.WorkoutRouteDataQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.WorkoutRouteDataQuery", q)
}

// The TraverseWorkoutRouteData type is an adapter to allow the use of ordinary function as Traverser.
type TraverseWorkoutRouteData func(context.Context, *ent.WorkoutRouteDataQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseWorkoutRouteData) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseWorkoutRouteData) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WorkoutRouteDataQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.WorkoutRouteDataQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.ActionQuery:
		return &query[*ent.ActionQuery, predicate.Action, action.OrderOption]{typ: ent.TypeAction, tq: q}, nil
	case *ent.AppConfigQuery:
		return &query[*ent.AppConfigQuery, predicate.AppConfig, appconfig.OrderOption]{typ: ent.TypeAppConfig, tq: q}, nil
	case *ent.CompetitionQuery:
		return &query[*ent.CompetitionQuery, predicate.Competition, competition.OrderOption]{typ: ent.TypeCompetition, tq: q}, nil
	case *ent.CompetitionResultQuery:
		return &query[*ent.CompetitionResultQuery, predicate.CompetitionResult, competitionresult.OrderOption]{typ: ent.TypeCompetitionResult, tq: q}, nil
	case *ent.FcmTokenQuery:
		return &query[*ent.FcmTokenQuery, predicate.FcmToken, fcmtoken.OrderOption]{typ: ent.TypeFcmToken, tq: q}, nil
	case *ent.GoalQuery:
		return &query[*ent.GoalQuery, predicate.Goal, goal.OrderOption]{typ: ent.TypeGoal, tq: q}, nil
	case *ent.InviteQuery:
		return &query[*ent.InviteQuery, predicate.Invite, invite.OrderOption]{typ: ent.TypeInvite, tq: q}, nil
	case *ent.NotificationQuery:
		return &query[*ent.NotificationQuery, predicate.Notification, notification.OrderOption]{typ: ent.TypeNotification, tq: q}, nil
	case *ent.NotificationPreferencesQuery:
		return &query[*ent.NotificationPreferencesQuery, predicate.NotificationPreferences, notificationpreferences.OrderOption]{typ: ent.TypeNotificationPreferences, tq: q}, nil
	case *ent.PostQuery:
		return &query[*ent.PostQuery, predicate.Post, post.OrderOption]{typ: ent.TypePost, tq: q}, nil
	case *ent.ProfileQuery:
		return &query[*ent.ProfileQuery, predicate.Profile, profile.OrderOption]{typ: ent.TypeProfile, tq: q}, nil
	case *ent.SquadQuery:
		return &query[*ent.SquadQuery, predicate.Squad, squad.OrderOption]{typ: ent.TypeSquad, tq: q}, nil
	case *ent.UserQuery:
		return &query[*ent.UserQuery, predicate.User, user.OrderOption]{typ: ent.TypeUser, tq: q}, nil
	case *ent.WorkoutQuery:
		return &query[*ent.WorkoutQuery, predicate.Workout, workout.OrderOption]{typ: ent.TypeWorkout, tq: q}, nil
	case *ent.WorkoutDataQuery:
		return &query[*ent.WorkoutDataQuery, predicate.WorkoutData, workoutdata.OrderOption]{typ: ent.TypeWorkoutData, tq: q}, nil
	case *ent.WorkoutRouteDataQuery:
		return &query[*ent.WorkoutRouteDataQuery, predicate.WorkoutRouteData, workoutroutedata.OrderOption]{typ: ent.TypeWorkoutRouteData, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
