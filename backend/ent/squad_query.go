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
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competitionresult"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/invite"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/squad"
)

// SquadQuery is the builder for querying Squad entities.
type SquadQuery struct {
	config
	ctx                    *QueryContext
	order                  []squad.OrderOption
	inters                 []Interceptor
	predicates             []predicate.Squad
	withMembers            *ProfileQuery
	withInvites            *InviteQuery
	withCompetitionResults *CompetitionResultQuery
	withOwner              *ProfileQuery
	withFKs                bool
	modifiers              []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SquadQuery builder.
func (sq *SquadQuery) Where(ps ...predicate.Squad) *SquadQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SquadQuery) Limit(limit int) *SquadQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SquadQuery) Offset(offset int) *SquadQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SquadQuery) Unique(unique bool) *SquadQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SquadQuery) Order(o ...squad.OrderOption) *SquadQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryMembers chains the current query on the "members" edge.
func (sq *SquadQuery) QueryMembers() *ProfileQuery {
	query := (&ProfileClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(squad.Table, squad.FieldID, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, squad.MembersTable, squad.MembersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvites chains the current query on the "invites" edge.
func (sq *SquadQuery) QueryInvites() *InviteQuery {
	query := (&InviteClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(squad.Table, squad.FieldID, selector),
			sqlgraph.To(invite.Table, invite.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, squad.InvitesTable, squad.InvitesColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompetitionResults chains the current query on the "competition_results" edge.
func (sq *SquadQuery) QueryCompetitionResults() *CompetitionResultQuery {
	query := (&CompetitionResultClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(squad.Table, squad.FieldID, selector),
			sqlgraph.To(competitionresult.Table, competitionresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, squad.CompetitionResultsTable, squad.CompetitionResultsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (sq *SquadQuery) QueryOwner() *ProfileQuery {
	query := (&ProfileClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(squad.Table, squad.FieldID, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, squad.OwnerTable, squad.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Squad entity from the query.
// Returns a *NotFoundError when no Squad was found.
func (sq *SquadQuery) First(ctx context.Context) (*Squad, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{squad.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SquadQuery) FirstX(ctx context.Context) *Squad {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Squad ID from the query.
// Returns a *NotFoundError when no Squad ID was found.
func (sq *SquadQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{squad.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SquadQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Squad entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Squad entity is found.
// Returns a *NotFoundError when no Squad entities are found.
func (sq *SquadQuery) Only(ctx context.Context) (*Squad, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{squad.Label}
	default:
		return nil, &NotSingularError{squad.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SquadQuery) OnlyX(ctx context.Context) *Squad {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Squad ID in the query.
// Returns a *NotSingularError when more than one Squad ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SquadQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{squad.Label}
	default:
		err = &NotSingularError{squad.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SquadQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Squads.
func (sq *SquadQuery) All(ctx context.Context) ([]*Squad, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Squad, *SquadQuery]()
	return withInterceptors[[]*Squad](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SquadQuery) AllX(ctx context.Context) []*Squad {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Squad IDs.
func (sq *SquadQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(squad.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SquadQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SquadQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SquadQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SquadQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SquadQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SquadQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SquadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SquadQuery) Clone() *SquadQuery {
	if sq == nil {
		return nil
	}
	return &SquadQuery{
		config:                 sq.config,
		ctx:                    sq.ctx.Clone(),
		order:                  append([]squad.OrderOption{}, sq.order...),
		inters:                 append([]Interceptor{}, sq.inters...),
		predicates:             append([]predicate.Squad{}, sq.predicates...),
		withMembers:            sq.withMembers.Clone(),
		withInvites:            sq.withInvites.Clone(),
		withCompetitionResults: sq.withCompetitionResults.Clone(),
		withOwner:              sq.withOwner.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithMembers tells the query-builder to eager-load the nodes that are connected to
// the "members" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SquadQuery) WithMembers(opts ...func(*ProfileQuery)) *SquadQuery {
	query := (&ProfileClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withMembers = query
	return sq
}

// WithInvites tells the query-builder to eager-load the nodes that are connected to
// the "invites" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SquadQuery) WithInvites(opts ...func(*InviteQuery)) *SquadQuery {
	query := (&InviteClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withInvites = query
	return sq
}

// WithCompetitionResults tells the query-builder to eager-load the nodes that are connected to
// the "competition_results" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SquadQuery) WithCompetitionResults(opts ...func(*CompetitionResultQuery)) *SquadQuery {
	query := (&CompetitionResultClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withCompetitionResults = query
	return sq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SquadQuery) WithOwner(opts ...func(*ProfileQuery)) *SquadQuery {
	query := (&ProfileClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withOwner = query
	return sq
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
//	client.Squad.Query().
//		GroupBy(squad.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SquadQuery) GroupBy(field string, fields ...string) *SquadGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SquadGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = squad.Label
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
//	client.Squad.Query().
//		Select(squad.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *SquadQuery) Select(fields ...string) *SquadSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SquadSelect{SquadQuery: sq}
	sbuild.label = squad.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SquadSelect configured with the given aggregations.
func (sq *SquadQuery) Aggregate(fns ...AggregateFunc) *SquadSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SquadQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !squad.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SquadQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Squad, error) {
	var (
		nodes       = []*Squad{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [4]bool{
			sq.withMembers != nil,
			sq.withInvites != nil,
			sq.withCompetitionResults != nil,
			sq.withOwner != nil,
		}
	)
	if sq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, squad.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Squad).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Squad{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withMembers; query != nil {
		if err := sq.loadMembers(ctx, query, nodes,
			func(n *Squad) { n.Edges.Members = []*Profile{} },
			func(n *Squad, e *Profile) { n.Edges.Members = append(n.Edges.Members, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withInvites; query != nil {
		if err := sq.loadInvites(ctx, query, nodes,
			func(n *Squad) { n.Edges.Invites = []*Invite{} },
			func(n *Squad, e *Invite) { n.Edges.Invites = append(n.Edges.Invites, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withCompetitionResults; query != nil {
		if err := sq.loadCompetitionResults(ctx, query, nodes,
			func(n *Squad) { n.Edges.CompetitionResults = []*CompetitionResult{} },
			func(n *Squad, e *CompetitionResult) {
				n.Edges.CompetitionResults = append(n.Edges.CompetitionResults, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := sq.withOwner; query != nil {
		if err := sq.loadOwner(ctx, query, nodes, nil,
			func(n *Squad, e *Profile) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SquadQuery) loadMembers(ctx context.Context, query *ProfileQuery, nodes []*Squad, init func(*Squad), assign func(*Squad, *Profile)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Squad)
	nids := make(map[uuid.UUID]map[*Squad]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(squad.MembersTable)
		s.Join(joinT).On(s.C(profile.FieldID), joinT.C(squad.MembersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(squad.MembersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(squad.MembersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Squad]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Profile](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "members" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (sq *SquadQuery) loadInvites(ctx context.Context, query *InviteQuery, nodes []*Squad, init func(*Squad), assign func(*Squad, *Invite)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Squad)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Invite(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(squad.InvitesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.invite_squad
		if fk == nil {
			return fmt.Errorf(`foreign-key "invite_squad" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "invite_squad" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SquadQuery) loadCompetitionResults(ctx context.Context, query *CompetitionResultQuery, nodes []*Squad, init func(*Squad), assign func(*Squad, *CompetitionResult)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Squad)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CompetitionResult(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(squad.CompetitionResultsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.competition_result_squad
		if fk == nil {
			return fmt.Errorf(`foreign-key "competition_result_squad" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "competition_result_squad" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SquadQuery) loadOwner(ctx context.Context, query *ProfileQuery, nodes []*Squad, init func(*Squad), assign func(*Squad, *Profile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Squad)
	for i := range nodes {
		if nodes[i].squad_owner == nil {
			continue
		}
		fk := *nodes[i].squad_owner
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
			return fmt.Errorf(`unexpected foreign-key "squad_owner" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SquadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SquadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(squad.Table, squad.Columns, sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, squad.FieldID)
		for i := range fields {
			if fields[i] != squad.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SquadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(squad.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = squad.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range sq.modifiers {
		m(selector)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sq *SquadQuery) Modify(modifiers ...func(s *sql.Selector)) *SquadSelect {
	sq.modifiers = append(sq.modifiers, modifiers...)
	return sq.Select()
}

// SquadGroupBy is the group-by builder for Squad entities.
type SquadGroupBy struct {
	selector
	build *SquadQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SquadGroupBy) Aggregate(fns ...AggregateFunc) *SquadGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SquadGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SquadQuery, *SquadGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SquadGroupBy) sqlScan(ctx context.Context, root *SquadQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SquadSelect is the builder for selecting fields of Squad entities.
type SquadSelect struct {
	*SquadQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SquadSelect) Aggregate(fns ...AggregateFunc) *SquadSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SquadSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SquadQuery, *SquadSelect](ctx, ss.SquadQuery, ss, ss.inters, v)
}

func (ss *SquadSelect) sqlScan(ctx context.Context, root *SquadQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ss *SquadSelect) Modify(modifiers ...func(s *sql.Selector)) *SquadSelect {
	ss.modifiers = append(ss.modifiers, modifiers...)
	return ss
}
