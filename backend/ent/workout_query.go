// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/invite"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutroutedata"
)

// WorkoutQuery is the builder for querying Workout entities.
type WorkoutQuery struct {
	config
	ctx                  *QueryContext
	order                []workout.OrderOption
	inters               []Interceptor
	predicates           []predicate.Workout
	withInvite           *InviteQuery
	withLeader           *ProfileQuery
	withCompetition      *CompetitionQuery
	withWorkoutData      *WorkoutDataQuery
	withWorkoutRouteData *WorkoutRouteDataQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkoutQuery builder.
func (wq *WorkoutQuery) Where(ps ...predicate.Workout) *WorkoutQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WorkoutQuery) Limit(limit int) *WorkoutQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WorkoutQuery) Offset(offset int) *WorkoutQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WorkoutQuery) Unique(unique bool) *WorkoutQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WorkoutQuery) Order(o ...workout.OrderOption) *WorkoutQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryInvite chains the current query on the "invite" edge.
func (wq *WorkoutQuery) QueryInvite() *InviteQuery {
	query := (&InviteClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, selector),
			sqlgraph.To(invite.Table, invite.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, workout.InviteTable, workout.InviteColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLeader chains the current query on the "leader" edge.
func (wq *WorkoutQuery) QueryLeader() *ProfileQuery {
	query := (&ProfileClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workout.LeaderTable, workout.LeaderColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompetition chains the current query on the "competition" edge.
func (wq *WorkoutQuery) QueryCompetition() *CompetitionQuery {
	query := (&CompetitionClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, selector),
			sqlgraph.To(competition.Table, competition.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workout.CompetitionTable, workout.CompetitionColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkoutData chains the current query on the "workout_data" edge.
func (wq *WorkoutQuery) QueryWorkoutData() *WorkoutDataQuery {
	query := (&WorkoutDataClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, selector),
			sqlgraph.To(workoutdata.Table, workoutdata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, workout.WorkoutDataTable, workout.WorkoutDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkoutRouteData chains the current query on the "workout_route_data" edge.
func (wq *WorkoutQuery) QueryWorkoutRouteData() *WorkoutRouteDataQuery {
	query := (&WorkoutRouteDataClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, selector),
			sqlgraph.To(workoutroutedata.Table, workoutroutedata.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, workout.WorkoutRouteDataTable, workout.WorkoutRouteDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Workout entity from the query.
// Returns a *NotFoundError when no Workout was found.
func (wq *WorkoutQuery) First(ctx context.Context) (*Workout, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workout.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WorkoutQuery) FirstX(ctx context.Context) *Workout {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Workout ID from the query.
// Returns a *NotFoundError when no Workout ID was found.
func (wq *WorkoutQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workout.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WorkoutQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Workout entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Workout entity is found.
// Returns a *NotFoundError when no Workout entities are found.
func (wq *WorkoutQuery) Only(ctx context.Context) (*Workout, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workout.Label}
	default:
		return nil, &NotSingularError{workout.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WorkoutQuery) OnlyX(ctx context.Context) *Workout {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Workout ID in the query.
// Returns a *NotSingularError when more than one Workout ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WorkoutQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workout.Label}
	default:
		err = &NotSingularError{workout.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WorkoutQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Workouts.
func (wq *WorkoutQuery) All(ctx context.Context) ([]*Workout, error) {
	ctx = setContextOp(ctx, wq.ctx, "All")
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Workout, *WorkoutQuery]()
	return withInterceptors[[]*Workout](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WorkoutQuery) AllX(ctx context.Context) []*Workout {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Workout IDs.
func (wq *WorkoutQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, "IDs")
	if err = wq.Select(workout.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WorkoutQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WorkoutQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, "Count")
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WorkoutQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WorkoutQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WorkoutQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, "Exist")
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WorkoutQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkoutQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WorkoutQuery) Clone() *WorkoutQuery {
	if wq == nil {
		return nil
	}
	return &WorkoutQuery{
		config:               wq.config,
		ctx:                  wq.ctx.Clone(),
		order:                append([]workout.OrderOption{}, wq.order...),
		inters:               append([]Interceptor{}, wq.inters...),
		predicates:           append([]predicate.Workout{}, wq.predicates...),
		withInvite:           wq.withInvite.Clone(),
		withLeader:           wq.withLeader.Clone(),
		withCompetition:      wq.withCompetition.Clone(),
		withWorkoutData:      wq.withWorkoutData.Clone(),
		withWorkoutRouteData: wq.withWorkoutRouteData.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithInvite tells the query-builder to eager-load the nodes that are connected to
// the "invite" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithInvite(opts ...func(*InviteQuery)) *WorkoutQuery {
	query := (&InviteClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withInvite = query
	return wq
}

// WithLeader tells the query-builder to eager-load the nodes that are connected to
// the "leader" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithLeader(opts ...func(*ProfileQuery)) *WorkoutQuery {
	query := (&ProfileClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withLeader = query
	return wq
}

// WithCompetition tells the query-builder to eager-load the nodes that are connected to
// the "competition" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithCompetition(opts ...func(*CompetitionQuery)) *WorkoutQuery {
	query := (&CompetitionClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withCompetition = query
	return wq
}

// WithWorkoutData tells the query-builder to eager-load the nodes that are connected to
// the "workout_data" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithWorkoutData(opts ...func(*WorkoutDataQuery)) *WorkoutQuery {
	query := (&WorkoutDataClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWorkoutData = query
	return wq
}

// WithWorkoutRouteData tells the query-builder to eager-load the nodes that are connected to
// the "workout_route_data" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WorkoutQuery) WithWorkoutRouteData(opts ...func(*WorkoutRouteDataQuery)) *WorkoutQuery {
	query := (&WorkoutRouteDataClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWorkoutRouteData = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Workout.Query().
//		GroupBy(workout.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WorkoutQuery) GroupBy(field string, fields ...string) *WorkoutGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkoutGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = workout.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Workout.Query().
//		Select(workout.FieldCreatedAt).
//		Scan(ctx, &v)
func (wq *WorkoutQuery) Select(fields ...string) *WorkoutSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WorkoutSelect{WorkoutQuery: wq}
	sbuild.label = workout.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkoutSelect configured with the given aggregations.
func (wq *WorkoutQuery) Aggregate(fns ...AggregateFunc) *WorkoutSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WorkoutQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !workout.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WorkoutQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Workout, error) {
	var (
		nodes       = []*Workout{}
		withFKs     = wq.withFKs
		_spec       = wq.querySpec()
		loadedTypes = [5]bool{
			wq.withInvite != nil,
			wq.withLeader != nil,
			wq.withCompetition != nil,
			wq.withWorkoutData != nil,
			wq.withWorkoutRouteData != nil,
		}
	)
	if wq.withLeader != nil || wq.withCompetition != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workout.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Workout).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Workout{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withInvite; query != nil {
		if err := wq.loadInvite(ctx, query, nodes,
			func(n *Workout) { n.Edges.Invite = []*Invite{} },
			func(n *Workout, e *Invite) { n.Edges.Invite = append(n.Edges.Invite, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withLeader; query != nil {
		if err := wq.loadLeader(ctx, query, nodes, nil,
			func(n *Workout, e *Profile) { n.Edges.Leader = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withCompetition; query != nil {
		if err := wq.loadCompetition(ctx, query, nodes, nil,
			func(n *Workout, e *Competition) { n.Edges.Competition = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWorkoutData; query != nil {
		if err := wq.loadWorkoutData(ctx, query, nodes,
			func(n *Workout) { n.Edges.WorkoutData = []*WorkoutData{} },
			func(n *Workout, e *WorkoutData) { n.Edges.WorkoutData = append(n.Edges.WorkoutData, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWorkoutRouteData; query != nil {
		if err := wq.loadWorkoutRouteData(ctx, query, nodes,
			func(n *Workout) { n.Edges.WorkoutRouteData = []*WorkoutRouteData{} },
			func(n *Workout, e *WorkoutRouteData) { n.Edges.WorkoutRouteData = append(n.Edges.WorkoutRouteData, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WorkoutQuery) loadInvite(ctx context.Context, query *InviteQuery, nodes []*Workout, init func(*Workout), assign func(*Workout, *Invite)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Workout)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Invite(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workout.InviteColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.invite_workout
		if fk == nil {
			return fmt.Errorf(`foreign-key "invite_workout" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "invite_workout" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WorkoutQuery) loadLeader(ctx context.Context, query *ProfileQuery, nodes []*Workout, init func(*Workout), assign func(*Workout, *Profile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Workout)
	for i := range nodes {
		if nodes[i].workout_leader == nil {
			continue
		}
		fk := *nodes[i].workout_leader
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(profile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workout_leader" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkoutQuery) loadCompetition(ctx context.Context, query *CompetitionQuery, nodes []*Workout, init func(*Workout), assign func(*Workout, *Competition)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Workout)
	for i := range nodes {
		if nodes[i].workout_competition == nil {
			continue
		}
		fk := *nodes[i].workout_competition
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(competition.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workout_competition" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WorkoutQuery) loadWorkoutData(ctx context.Context, query *WorkoutDataQuery, nodes []*Workout, init func(*Workout), assign func(*Workout, *WorkoutData)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Workout)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.WorkoutData(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workout.WorkoutDataColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.workout_data_workout
		if fk == nil {
			return fmt.Errorf(`foreign-key "workout_data_workout" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workout_data_workout" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WorkoutQuery) loadWorkoutRouteData(ctx context.Context, query *WorkoutRouteDataQuery, nodes []*Workout, init func(*Workout), assign func(*Workout, *WorkoutRouteData)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Workout)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.WorkoutRouteData(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workout.WorkoutRouteDataColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.workout_route_data_workout
		if fk == nil {
			return fmt.Errorf(`foreign-key "workout_route_data_workout" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workout_route_data_workout" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wq *WorkoutQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WorkoutQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workout.Table, workout.Columns, sqlgraph.NewFieldSpec(workout.FieldID, field.TypeUUID))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workout.FieldID)
		for i := range fields {
			if fields[i] != workout.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WorkoutQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(workout.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = workout.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wq.modifiers {
		m(selector)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wq *WorkoutQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkoutSelect {
	wq.modifiers = append(wq.modifiers, modifiers...)
	return wq.Select()
}

// WorkoutGroupBy is the group-by builder for Workout entities.
type WorkoutGroupBy struct {
	selector
	build *WorkoutQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WorkoutGroupBy) Aggregate(fns ...AggregateFunc) *WorkoutGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WorkoutGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, "GroupBy")
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkoutQuery, *WorkoutGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WorkoutGroupBy) sqlScan(ctx context.Context, root *WorkoutQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkoutSelect is the builder for selecting fields of Workout entities.
type WorkoutSelect struct {
	*WorkoutQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WorkoutSelect) Aggregate(fns ...AggregateFunc) *WorkoutSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WorkoutSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, "Select")
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkoutQuery, *WorkoutSelect](ctx, ws.WorkoutQuery, ws, ws.inters, v)
}

func (ws *WorkoutSelect) sqlScan(ctx context.Context, root *WorkoutQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ws *WorkoutSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkoutSelect {
	ws.modifiers = append(ws.modifiers, modifiers...)
	return ws
}