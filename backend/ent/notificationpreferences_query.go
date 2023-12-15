// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/notificationpreferences"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
)

// NotificationPreferencesQuery is the builder for querying NotificationPreferences entities.
type NotificationPreferencesQuery struct {
	config
	ctx         *QueryContext
	order       []notificationpreferences.OrderOption
	inters      []Interceptor
	predicates  []predicate.NotificationPreferences
	withProfile *ProfileQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotificationPreferencesQuery builder.
func (npq *NotificationPreferencesQuery) Where(ps ...predicate.NotificationPreferences) *NotificationPreferencesQuery {
	npq.predicates = append(npq.predicates, ps...)
	return npq
}

// Limit the number of records to be returned by this query.
func (npq *NotificationPreferencesQuery) Limit(limit int) *NotificationPreferencesQuery {
	npq.ctx.Limit = &limit
	return npq
}

// Offset to start from.
func (npq *NotificationPreferencesQuery) Offset(offset int) *NotificationPreferencesQuery {
	npq.ctx.Offset = &offset
	return npq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (npq *NotificationPreferencesQuery) Unique(unique bool) *NotificationPreferencesQuery {
	npq.ctx.Unique = &unique
	return npq
}

// Order specifies how the records should be ordered.
func (npq *NotificationPreferencesQuery) Order(o ...notificationpreferences.OrderOption) *NotificationPreferencesQuery {
	npq.order = append(npq.order, o...)
	return npq
}

// QueryProfile chains the current query on the "profile" edge.
func (npq *NotificationPreferencesQuery) QueryProfile() *ProfileQuery {
	query := (&ProfileClient{config: npq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := npq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := npq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notificationpreferences.Table, notificationpreferences.FieldID, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, notificationpreferences.ProfileTable, notificationpreferences.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(npq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NotificationPreferences entity from the query.
// Returns a *NotFoundError when no NotificationPreferences was found.
func (npq *NotificationPreferencesQuery) First(ctx context.Context) (*NotificationPreferences, error) {
	nodes, err := npq.Limit(1).All(setContextOp(ctx, npq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notificationpreferences.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (npq *NotificationPreferencesQuery) FirstX(ctx context.Context) *NotificationPreferences {
	node, err := npq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NotificationPreferences ID from the query.
// Returns a *NotFoundError when no NotificationPreferences ID was found.
func (npq *NotificationPreferencesQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = npq.Limit(1).IDs(setContextOp(ctx, npq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notificationpreferences.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (npq *NotificationPreferencesQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := npq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NotificationPreferences entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NotificationPreferences entity is found.
// Returns a *NotFoundError when no NotificationPreferences entities are found.
func (npq *NotificationPreferencesQuery) Only(ctx context.Context) (*NotificationPreferences, error) {
	nodes, err := npq.Limit(2).All(setContextOp(ctx, npq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notificationpreferences.Label}
	default:
		return nil, &NotSingularError{notificationpreferences.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (npq *NotificationPreferencesQuery) OnlyX(ctx context.Context) *NotificationPreferences {
	node, err := npq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NotificationPreferences ID in the query.
// Returns a *NotSingularError when more than one NotificationPreferences ID is found.
// Returns a *NotFoundError when no entities are found.
func (npq *NotificationPreferencesQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = npq.Limit(2).IDs(setContextOp(ctx, npq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notificationpreferences.Label}
	default:
		err = &NotSingularError{notificationpreferences.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (npq *NotificationPreferencesQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := npq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NotificationPreferencesSlice.
func (npq *NotificationPreferencesQuery) All(ctx context.Context) ([]*NotificationPreferences, error) {
	ctx = setContextOp(ctx, npq.ctx, "All")
	if err := npq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NotificationPreferences, *NotificationPreferencesQuery]()
	return withInterceptors[[]*NotificationPreferences](ctx, npq, qr, npq.inters)
}

// AllX is like All, but panics if an error occurs.
func (npq *NotificationPreferencesQuery) AllX(ctx context.Context) []*NotificationPreferences {
	nodes, err := npq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NotificationPreferences IDs.
func (npq *NotificationPreferencesQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if npq.ctx.Unique == nil && npq.path != nil {
		npq.Unique(true)
	}
	ctx = setContextOp(ctx, npq.ctx, "IDs")
	if err = npq.Select(notificationpreferences.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (npq *NotificationPreferencesQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := npq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (npq *NotificationPreferencesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, npq.ctx, "Count")
	if err := npq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, npq, querierCount[*NotificationPreferencesQuery](), npq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (npq *NotificationPreferencesQuery) CountX(ctx context.Context) int {
	count, err := npq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (npq *NotificationPreferencesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, npq.ctx, "Exist")
	switch _, err := npq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (npq *NotificationPreferencesQuery) ExistX(ctx context.Context) bool {
	exist, err := npq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotificationPreferencesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (npq *NotificationPreferencesQuery) Clone() *NotificationPreferencesQuery {
	if npq == nil {
		return nil
	}
	return &NotificationPreferencesQuery{
		config:      npq.config,
		ctx:         npq.ctx.Clone(),
		order:       append([]notificationpreferences.OrderOption{}, npq.order...),
		inters:      append([]Interceptor{}, npq.inters...),
		predicates:  append([]predicate.NotificationPreferences{}, npq.predicates...),
		withProfile: npq.withProfile.Clone(),
		// clone intermediate query.
		sql:  npq.sql.Clone(),
		path: npq.path,
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (npq *NotificationPreferencesQuery) WithProfile(opts ...func(*ProfileQuery)) *NotificationPreferencesQuery {
	query := (&ProfileClient{config: npq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	npq.withProfile = query
	return npq
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
//	client.NotificationPreferences.Query().
//		GroupBy(notificationpreferences.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (npq *NotificationPreferencesQuery) GroupBy(field string, fields ...string) *NotificationPreferencesGroupBy {
	npq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NotificationPreferencesGroupBy{build: npq}
	grbuild.flds = &npq.ctx.Fields
	grbuild.label = notificationpreferences.Label
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
//	client.NotificationPreferences.Query().
//		Select(notificationpreferences.FieldCreatedAt).
//		Scan(ctx, &v)
func (npq *NotificationPreferencesQuery) Select(fields ...string) *NotificationPreferencesSelect {
	npq.ctx.Fields = append(npq.ctx.Fields, fields...)
	sbuild := &NotificationPreferencesSelect{NotificationPreferencesQuery: npq}
	sbuild.label = notificationpreferences.Label
	sbuild.flds, sbuild.scan = &npq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NotificationPreferencesSelect configured with the given aggregations.
func (npq *NotificationPreferencesQuery) Aggregate(fns ...AggregateFunc) *NotificationPreferencesSelect {
	return npq.Select().Aggregate(fns...)
}

func (npq *NotificationPreferencesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range npq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, npq); err != nil {
				return err
			}
		}
	}
	for _, f := range npq.ctx.Fields {
		if !notificationpreferences.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if npq.path != nil {
		prev, err := npq.path(ctx)
		if err != nil {
			return err
		}
		npq.sql = prev
	}
	return nil
}

func (npq *NotificationPreferencesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NotificationPreferences, error) {
	var (
		nodes       = []*NotificationPreferences{}
		withFKs     = npq.withFKs
		_spec       = npq.querySpec()
		loadedTypes = [1]bool{
			npq.withProfile != nil,
		}
	)
	if npq.withProfile != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, notificationpreferences.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NotificationPreferences).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NotificationPreferences{config: npq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(npq.modifiers) > 0 {
		_spec.Modifiers = npq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, npq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := npq.withProfile; query != nil {
		if err := npq.loadProfile(ctx, query, nodes, nil,
			func(n *NotificationPreferences, e *Profile) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (npq *NotificationPreferencesQuery) loadProfile(ctx context.Context, query *ProfileQuery, nodes []*NotificationPreferences, init func(*NotificationPreferences), assign func(*NotificationPreferences, *Profile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*NotificationPreferences)
	for i := range nodes {
		if nodes[i].profile_notification_preferences == nil {
			continue
		}
		fk := *nodes[i].profile_notification_preferences
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
			return fmt.Errorf(`unexpected foreign-key "profile_notification_preferences" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (npq *NotificationPreferencesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := npq.querySpec()
	if len(npq.modifiers) > 0 {
		_spec.Modifiers = npq.modifiers
	}
	_spec.Node.Columns = npq.ctx.Fields
	if len(npq.ctx.Fields) > 0 {
		_spec.Unique = npq.ctx.Unique != nil && *npq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, npq.driver, _spec)
}

func (npq *NotificationPreferencesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(notificationpreferences.Table, notificationpreferences.Columns, sqlgraph.NewFieldSpec(notificationpreferences.FieldID, field.TypeUUID))
	_spec.From = npq.sql
	if unique := npq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if npq.path != nil {
		_spec.Unique = true
	}
	if fields := npq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationpreferences.FieldID)
		for i := range fields {
			if fields[i] != notificationpreferences.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := npq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := npq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := npq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := npq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (npq *NotificationPreferencesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(npq.driver.Dialect())
	t1 := builder.Table(notificationpreferences.Table)
	columns := npq.ctx.Fields
	if len(columns) == 0 {
		columns = notificationpreferences.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if npq.sql != nil {
		selector = npq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if npq.ctx.Unique != nil && *npq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range npq.modifiers {
		m(selector)
	}
	for _, p := range npq.predicates {
		p(selector)
	}
	for _, p := range npq.order {
		p(selector)
	}
	if offset := npq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := npq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (npq *NotificationPreferencesQuery) Modify(modifiers ...func(s *sql.Selector)) *NotificationPreferencesSelect {
	npq.modifiers = append(npq.modifiers, modifiers...)
	return npq.Select()
}

// NotificationPreferencesGroupBy is the group-by builder for NotificationPreferences entities.
type NotificationPreferencesGroupBy struct {
	selector
	build *NotificationPreferencesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (npgb *NotificationPreferencesGroupBy) Aggregate(fns ...AggregateFunc) *NotificationPreferencesGroupBy {
	npgb.fns = append(npgb.fns, fns...)
	return npgb
}

// Scan applies the selector query and scans the result into the given value.
func (npgb *NotificationPreferencesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, npgb.build.ctx, "GroupBy")
	if err := npgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationPreferencesQuery, *NotificationPreferencesGroupBy](ctx, npgb.build, npgb, npgb.build.inters, v)
}

func (npgb *NotificationPreferencesGroupBy) sqlScan(ctx context.Context, root *NotificationPreferencesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(npgb.fns))
	for _, fn := range npgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*npgb.flds)+len(npgb.fns))
		for _, f := range *npgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*npgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := npgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NotificationPreferencesSelect is the builder for selecting fields of NotificationPreferences entities.
type NotificationPreferencesSelect struct {
	*NotificationPreferencesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nps *NotificationPreferencesSelect) Aggregate(fns ...AggregateFunc) *NotificationPreferencesSelect {
	nps.fns = append(nps.fns, fns...)
	return nps
}

// Scan applies the selector query and scans the result into the given value.
func (nps *NotificationPreferencesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nps.ctx, "Select")
	if err := nps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationPreferencesQuery, *NotificationPreferencesSelect](ctx, nps.NotificationPreferencesQuery, nps, nps.inters, v)
}

func (nps *NotificationPreferencesSelect) sqlScan(ctx context.Context, root *NotificationPreferencesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nps.fns))
	for _, fn := range nps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (nps *NotificationPreferencesSelect) Modify(modifiers ...func(s *sql.Selector)) *NotificationPreferencesSelect {
	nps.modifiers = append(nps.modifiers, modifiers...)
	return nps
}