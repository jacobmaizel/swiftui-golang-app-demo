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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workout"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutdata"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/workoutroutedata"
)

// WorkoutDataQuery is the builder for querying WorkoutData entities.
type WorkoutDataQuery struct {
	config
	ctx                  *QueryContext
	order                []workoutdata.OrderOption
	inters               []Interceptor
	predicates           []predicate.WorkoutData
	withWorkout          *WorkoutQuery
	withProfile          *ProfileQuery
	withWorkoutRouteData *WorkoutRouteDataQuery
	withCompetition      *CompetitionQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkoutDataQuery builder.
func (wdq *WorkoutDataQuery) Where(ps ...predicate.WorkoutData) *WorkoutDataQuery {
	wdq.predicates = append(wdq.predicates, ps...)
	return wdq
}

// Limit the number of records to be returned by this query.
func (wdq *WorkoutDataQuery) Limit(limit int) *WorkoutDataQuery {
	wdq.ctx.Limit = &limit
	return wdq
}

// Offset to start from.
func (wdq *WorkoutDataQuery) Offset(offset int) *WorkoutDataQuery {
	wdq.ctx.Offset = &offset
	return wdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wdq *WorkoutDataQuery) Unique(unique bool) *WorkoutDataQuery {
	wdq.ctx.Unique = &unique
	return wdq
}

// Order specifies how the records should be ordered.
func (wdq *WorkoutDataQuery) Order(o ...workoutdata.OrderOption) *WorkoutDataQuery {
	wdq.order = append(wdq.order, o...)
	return wdq
}

// QueryWorkout chains the current query on the "workout" edge.
func (wdq *WorkoutDataQuery) QueryWorkout() *WorkoutQuery {
	query := (&WorkoutClient{config: wdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workoutdata.Table, workoutdata.FieldID, selector),
			sqlgraph.To(workout.Table, workout.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workoutdata.WorkoutTable, workoutdata.WorkoutColumn),
		)
		fromU = sqlgraph.SetNeighbors(wdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProfile chains the current query on the "profile" edge.
func (wdq *WorkoutDataQuery) QueryProfile() *ProfileQuery {
	query := (&ProfileClient{config: wdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workoutdata.Table, workoutdata.FieldID, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workoutdata.ProfileTable, workoutdata.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(wdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkoutRouteData chains the current query on the "workout_route_data" edge.
func (wdq *WorkoutDataQuery) QueryWorkoutRouteData() *WorkoutRouteDataQuery {
	query := (&WorkoutRouteDataClient{config: wdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workoutdata.Table, workoutdata.FieldID, selector),
			sqlgraph.To(workoutroutedata.Table, workoutroutedata.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, workoutdata.WorkoutRouteDataTable, workoutdata.WorkoutRouteDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(wdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompetition chains the current query on the "competition" edge.
func (wdq *WorkoutDataQuery) QueryCompetition() *CompetitionQuery {
	query := (&CompetitionClient{config: wdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workoutdata.Table, workoutdata.FieldID, selector),
			sqlgraph.To(competition.Table, competition.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workoutdata.CompetitionTable, workoutdata.CompetitionColumn),
		)
		fromU = sqlgraph.SetNeighbors(wdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkoutData entity from the query.
// Returns a *NotFoundError when no WorkoutData was found.
func (wdq *WorkoutDataQuery) First(ctx context.Context) (*WorkoutData, error) {
	nodes, err := wdq.Limit(1).All(setContextOp(ctx, wdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workoutdata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wdq *WorkoutDataQuery) FirstX(ctx context.Context) *WorkoutData {
	node, err := wdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkoutData ID from the query.
// Returns a *NotFoundError when no WorkoutData ID was found.
func (wdq *WorkoutDataQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wdq.Limit(1).IDs(setContextOp(ctx, wdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workoutdata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wdq *WorkoutDataQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkoutData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkoutData entity is found.
// Returns a *NotFoundError when no WorkoutData entities are found.
func (wdq *WorkoutDataQuery) Only(ctx context.Context) (*WorkoutData, error) {
	nodes, err := wdq.Limit(2).All(setContextOp(ctx, wdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workoutdata.Label}
	default:
		return nil, &NotSingularError{workoutdata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wdq *WorkoutDataQuery) OnlyX(ctx context.Context) *WorkoutData {
	node, err := wdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkoutData ID in the query.
// Returns a *NotSingularError when more than one WorkoutData ID is found.
// Returns a *NotFoundError when no entities are found.
func (wdq *WorkoutDataQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wdq.Limit(2).IDs(setContextOp(ctx, wdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workoutdata.Label}
	default:
		err = &NotSingularError{workoutdata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wdq *WorkoutDataQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkoutDataSlice.
func (wdq *WorkoutDataQuery) All(ctx context.Context) ([]*WorkoutData, error) {
	ctx = setContextOp(ctx, wdq.ctx, "All")
	if err := wdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WorkoutData, *WorkoutDataQuery]()
	return withInterceptors[[]*WorkoutData](ctx, wdq, qr, wdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wdq *WorkoutDataQuery) AllX(ctx context.Context) []*WorkoutData {
	nodes, err := wdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkoutData IDs.
func (wdq *WorkoutDataQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wdq.ctx.Unique == nil && wdq.path != nil {
		wdq.Unique(true)
	}
	ctx = setContextOp(ctx, wdq.ctx, "IDs")
	if err = wdq.Select(workoutdata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wdq *WorkoutDataQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wdq *WorkoutDataQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wdq.ctx, "Count")
	if err := wdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wdq, querierCount[*WorkoutDataQuery](), wdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wdq *WorkoutDataQuery) CountX(ctx context.Context) int {
	count, err := wdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wdq *WorkoutDataQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wdq.ctx, "Exist")
	switch _, err := wdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wdq *WorkoutDataQuery) ExistX(ctx context.Context) bool {
	exist, err := wdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkoutDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wdq *WorkoutDataQuery) Clone() *WorkoutDataQuery {
	if wdq == nil {
		return nil
	}
	return &WorkoutDataQuery{
		config:               wdq.config,
		ctx:                  wdq.ctx.Clone(),
		order:                append([]workoutdata.OrderOption{}, wdq.order...),
		inters:               append([]Interceptor{}, wdq.inters...),
		predicates:           append([]predicate.WorkoutData{}, wdq.predicates...),
		withWorkout:          wdq.withWorkout.Clone(),
		withProfile:          wdq.withProfile.Clone(),
		withWorkoutRouteData: wdq.withWorkoutRouteData.Clone(),
		withCompetition:      wdq.withCompetition.Clone(),
		// clone intermediate query.
		sql:  wdq.sql.Clone(),
		path: wdq.path,
	}
}

// WithWorkout tells the query-builder to eager-load the nodes that are connected to
// the "workout" edge. The optional arguments are used to configure the query builder of the edge.
func (wdq *WorkoutDataQuery) WithWorkout(opts ...func(*WorkoutQuery)) *WorkoutDataQuery {
	query := (&WorkoutClient{config: wdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wdq.withWorkout = query
	return wdq
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (wdq *WorkoutDataQuery) WithProfile(opts ...func(*ProfileQuery)) *WorkoutDataQuery {
	query := (&ProfileClient{config: wdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wdq.withProfile = query
	return wdq
}

// WithWorkoutRouteData tells the query-builder to eager-load the nodes that are connected to
// the "workout_route_data" edge. The optional arguments are used to configure the query builder of the edge.
func (wdq *WorkoutDataQuery) WithWorkoutRouteData(opts ...func(*WorkoutRouteDataQuery)) *WorkoutDataQuery {
	query := (&WorkoutRouteDataClient{config: wdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wdq.withWorkoutRouteData = query
	return wdq
}

// WithCompetition tells the query-builder to eager-load the nodes that are connected to
// the "competition" edge. The optional arguments are used to configure the query builder of the edge.
func (wdq *WorkoutDataQuery) WithCompetition(opts ...func(*CompetitionQuery)) *WorkoutDataQuery {
	query := (&CompetitionClient{config: wdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wdq.withCompetition = query
	return wdq
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
//	client.WorkoutData.Query().
//		GroupBy(workoutdata.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wdq *WorkoutDataQuery) GroupBy(field string, fields ...string) *WorkoutDataGroupBy {
	wdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkoutDataGroupBy{build: wdq}
	grbuild.flds = &wdq.ctx.Fields
	grbuild.label = workoutdata.Label
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
//	client.WorkoutData.Query().
//		Select(workoutdata.FieldCreatedAt).
//		Scan(ctx, &v)
func (wdq *WorkoutDataQuery) Select(fields ...string) *WorkoutDataSelect {
	wdq.ctx.Fields = append(wdq.ctx.Fields, fields...)
	sbuild := &WorkoutDataSelect{WorkoutDataQuery: wdq}
	sbuild.label = workoutdata.Label
	sbuild.flds, sbuild.scan = &wdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkoutDataSelect configured with the given aggregations.
func (wdq *WorkoutDataQuery) Aggregate(fns ...AggregateFunc) *WorkoutDataSelect {
	return wdq.Select().Aggregate(fns...)
}

func (wdq *WorkoutDataQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wdq); err != nil {
				return err
			}
		}
	}
	for _, f := range wdq.ctx.Fields {
		if !workoutdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wdq.path != nil {
		prev, err := wdq.path(ctx)
		if err != nil {
			return err
		}
		wdq.sql = prev
	}
	return nil
}

func (wdq *WorkoutDataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkoutData, error) {
	var (
		nodes       = []*WorkoutData{}
		withFKs     = wdq.withFKs
		_spec       = wdq.querySpec()
		loadedTypes = [4]bool{
			wdq.withWorkout != nil,
			wdq.withProfile != nil,
			wdq.withWorkoutRouteData != nil,
			wdq.withCompetition != nil,
		}
	)
	if wdq.withWorkout != nil || wdq.withProfile != nil || wdq.withCompetition != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workoutdata.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WorkoutData).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WorkoutData{config: wdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wdq.modifiers) > 0 {
		_spec.Modifiers = wdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wdq.withWorkout; query != nil {
		if err := wdq.loadWorkout(ctx, query, nodes, nil,
			func(n *WorkoutData, e *Workout) { n.Edges.Workout = e }); err != nil {
			return nil, err
		}
	}
	if query := wdq.withProfile; query != nil {
		if err := wdq.loadProfile(ctx, query, nodes, nil,
			func(n *WorkoutData, e *Profile) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	if query := wdq.withWorkoutRouteData; query != nil {
		if err := wdq.loadWorkoutRouteData(ctx, query, nodes, nil,
			func(n *WorkoutData, e *WorkoutRouteData) { n.Edges.WorkoutRouteData = e }); err != nil {
			return nil, err
		}
	}
	if query := wdq.withCompetition; query != nil {
		if err := wdq.loadCompetition(ctx, query, nodes, nil,
			func(n *WorkoutData, e *Competition) { n.Edges.Competition = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wdq *WorkoutDataQuery) loadWorkout(ctx context.Context, query *WorkoutQuery, nodes []*WorkoutData, init func(*WorkoutData), assign func(*WorkoutData, *Workout)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkoutData)
	for i := range nodes {
		if nodes[i].workout_data_workout == nil {
			continue
		}
		fk := *nodes[i].workout_data_workout
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workout.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workout_data_workout" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wdq *WorkoutDataQuery) loadProfile(ctx context.Context, query *ProfileQuery, nodes []*WorkoutData, init func(*WorkoutData), assign func(*WorkoutData, *Profile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkoutData)
	for i := range nodes {
		if nodes[i].workout_data_profile == nil {
			continue
		}
		fk := *nodes[i].workout_data_profile
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
			return fmt.Errorf(`unexpected foreign-key "workout_data_profile" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wdq *WorkoutDataQuery) loadWorkoutRouteData(ctx context.Context, query *WorkoutRouteDataQuery, nodes []*WorkoutData, init func(*WorkoutData), assign func(*WorkoutData, *WorkoutRouteData)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*WorkoutData)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.WorkoutRouteData(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workoutdata.WorkoutRouteDataColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.workout_data_workout_route_data
		if fk == nil {
			return fmt.Errorf(`foreign-key "workout_data_workout_route_data" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workout_data_workout_route_data" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wdq *WorkoutDataQuery) loadCompetition(ctx context.Context, query *CompetitionQuery, nodes []*WorkoutData, init func(*WorkoutData), assign func(*WorkoutData, *Competition)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkoutData)
	for i := range nodes {
		if nodes[i].workout_data_competition == nil {
			continue
		}
		fk := *nodes[i].workout_data_competition
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
			return fmt.Errorf(`unexpected foreign-key "workout_data_competition" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wdq *WorkoutDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wdq.querySpec()
	if len(wdq.modifiers) > 0 {
		_spec.Modifiers = wdq.modifiers
	}
	_spec.Node.Columns = wdq.ctx.Fields
	if len(wdq.ctx.Fields) > 0 {
		_spec.Unique = wdq.ctx.Unique != nil && *wdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wdq.driver, _spec)
}

func (wdq *WorkoutDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workoutdata.Table, workoutdata.Columns, sqlgraph.NewFieldSpec(workoutdata.FieldID, field.TypeUUID))
	_spec.From = wdq.sql
	if unique := wdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wdq.path != nil {
		_spec.Unique = true
	}
	if fields := wdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workoutdata.FieldID)
		for i := range fields {
			if fields[i] != workoutdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wdq *WorkoutDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wdq.driver.Dialect())
	t1 := builder.Table(workoutdata.Table)
	columns := wdq.ctx.Fields
	if len(columns) == 0 {
		columns = workoutdata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wdq.sql != nil {
		selector = wdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wdq.ctx.Unique != nil && *wdq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wdq.modifiers {
		m(selector)
	}
	for _, p := range wdq.predicates {
		p(selector)
	}
	for _, p := range wdq.order {
		p(selector)
	}
	if offset := wdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wdq *WorkoutDataQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkoutDataSelect {
	wdq.modifiers = append(wdq.modifiers, modifiers...)
	return wdq.Select()
}

// WorkoutDataGroupBy is the group-by builder for WorkoutData entities.
type WorkoutDataGroupBy struct {
	selector
	build *WorkoutDataQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wdgb *WorkoutDataGroupBy) Aggregate(fns ...AggregateFunc) *WorkoutDataGroupBy {
	wdgb.fns = append(wdgb.fns, fns...)
	return wdgb
}

// Scan applies the selector query and scans the result into the given value.
func (wdgb *WorkoutDataGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wdgb.build.ctx, "GroupBy")
	if err := wdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkoutDataQuery, *WorkoutDataGroupBy](ctx, wdgb.build, wdgb, wdgb.build.inters, v)
}

func (wdgb *WorkoutDataGroupBy) sqlScan(ctx context.Context, root *WorkoutDataQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wdgb.fns))
	for _, fn := range wdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wdgb.flds)+len(wdgb.fns))
		for _, f := range *wdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkoutDataSelect is the builder for selecting fields of WorkoutData entities.
type WorkoutDataSelect struct {
	*WorkoutDataQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wds *WorkoutDataSelect) Aggregate(fns ...AggregateFunc) *WorkoutDataSelect {
	wds.fns = append(wds.fns, fns...)
	return wds
}

// Scan applies the selector query and scans the result into the given value.
func (wds *WorkoutDataSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wds.ctx, "Select")
	if err := wds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkoutDataQuery, *WorkoutDataSelect](ctx, wds.WorkoutDataQuery, wds, wds.inters, v)
}

func (wds *WorkoutDataSelect) sqlScan(ctx context.Context, root *WorkoutDataQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wds.fns))
	for _, fn := range wds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wds *WorkoutDataSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkoutDataSelect {
	wds.modifiers = append(wds.modifiers, modifiers...)
	return wds
}