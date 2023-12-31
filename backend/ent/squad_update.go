// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

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

// SquadUpdate is the builder for updating Squad entities.
type SquadUpdate struct {
	config
	hooks     []Hook
	mutation  *SquadMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SquadUpdate builder.
func (su *SquadUpdate) Where(ps ...predicate.Squad) *SquadUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SquadUpdate) SetUpdatedAt(t time.Time) *SquadUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetPublic sets the "public" field.
func (su *SquadUpdate) SetPublic(b bool) *SquadUpdate {
	su.mutation.SetPublic(b)
	return su
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (su *SquadUpdate) SetNillablePublic(b *bool) *SquadUpdate {
	if b != nil {
		su.SetPublic(*b)
	}
	return su
}

// SetTitle sets the "title" field.
func (su *SquadUpdate) SetTitle(s string) *SquadUpdate {
	su.mutation.SetTitle(s)
	return su
}

// AddMemberIDs adds the "members" edge to the Profile entity by IDs.
func (su *SquadUpdate) AddMemberIDs(ids ...uuid.UUID) *SquadUpdate {
	su.mutation.AddMemberIDs(ids...)
	return su
}

// AddMembers adds the "members" edges to the Profile entity.
func (su *SquadUpdate) AddMembers(p ...*Profile) *SquadUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddMemberIDs(ids...)
}

// AddInviteIDs adds the "invites" edge to the Invite entity by IDs.
func (su *SquadUpdate) AddInviteIDs(ids ...uuid.UUID) *SquadUpdate {
	su.mutation.AddInviteIDs(ids...)
	return su
}

// AddInvites adds the "invites" edges to the Invite entity.
func (su *SquadUpdate) AddInvites(i ...*Invite) *SquadUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return su.AddInviteIDs(ids...)
}

// AddCompetitionResultIDs adds the "competition_results" edge to the CompetitionResult entity by IDs.
func (su *SquadUpdate) AddCompetitionResultIDs(ids ...uuid.UUID) *SquadUpdate {
	su.mutation.AddCompetitionResultIDs(ids...)
	return su
}

// AddCompetitionResults adds the "competition_results" edges to the CompetitionResult entity.
func (su *SquadUpdate) AddCompetitionResults(c ...*CompetitionResult) *SquadUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCompetitionResultIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Profile entity by ID.
func (su *SquadUpdate) SetOwnerID(id uuid.UUID) *SquadUpdate {
	su.mutation.SetOwnerID(id)
	return su
}

// SetNillableOwnerID sets the "owner" edge to the Profile entity by ID if the given value is not nil.
func (su *SquadUpdate) SetNillableOwnerID(id *uuid.UUID) *SquadUpdate {
	if id != nil {
		su = su.SetOwnerID(*id)
	}
	return su
}

// SetOwner sets the "owner" edge to the Profile entity.
func (su *SquadUpdate) SetOwner(p *Profile) *SquadUpdate {
	return su.SetOwnerID(p.ID)
}

// Mutation returns the SquadMutation object of the builder.
func (su *SquadUpdate) Mutation() *SquadMutation {
	return su.mutation
}

// ClearMembers clears all "members" edges to the Profile entity.
func (su *SquadUpdate) ClearMembers() *SquadUpdate {
	su.mutation.ClearMembers()
	return su
}

// RemoveMemberIDs removes the "members" edge to Profile entities by IDs.
func (su *SquadUpdate) RemoveMemberIDs(ids ...uuid.UUID) *SquadUpdate {
	su.mutation.RemoveMemberIDs(ids...)
	return su
}

// RemoveMembers removes "members" edges to Profile entities.
func (su *SquadUpdate) RemoveMembers(p ...*Profile) *SquadUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemoveMemberIDs(ids...)
}

// ClearInvites clears all "invites" edges to the Invite entity.
func (su *SquadUpdate) ClearInvites() *SquadUpdate {
	su.mutation.ClearInvites()
	return su
}

// RemoveInviteIDs removes the "invites" edge to Invite entities by IDs.
func (su *SquadUpdate) RemoveInviteIDs(ids ...uuid.UUID) *SquadUpdate {
	su.mutation.RemoveInviteIDs(ids...)
	return su
}

// RemoveInvites removes "invites" edges to Invite entities.
func (su *SquadUpdate) RemoveInvites(i ...*Invite) *SquadUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return su.RemoveInviteIDs(ids...)
}

// ClearCompetitionResults clears all "competition_results" edges to the CompetitionResult entity.
func (su *SquadUpdate) ClearCompetitionResults() *SquadUpdate {
	su.mutation.ClearCompetitionResults()
	return su
}

// RemoveCompetitionResultIDs removes the "competition_results" edge to CompetitionResult entities by IDs.
func (su *SquadUpdate) RemoveCompetitionResultIDs(ids ...uuid.UUID) *SquadUpdate {
	su.mutation.RemoveCompetitionResultIDs(ids...)
	return su
}

// RemoveCompetitionResults removes "competition_results" edges to CompetitionResult entities.
func (su *SquadUpdate) RemoveCompetitionResults(c ...*CompetitionResult) *SquadUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCompetitionResultIDs(ids...)
}

// ClearOwner clears the "owner" edge to the Profile entity.
func (su *SquadUpdate) ClearOwner() *SquadUpdate {
	su.mutation.ClearOwner()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SquadUpdate) Save(ctx context.Context) (int, error) {
	if err := su.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SquadUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SquadUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SquadUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SquadUpdate) defaults() error {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		if squad.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized squad.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := squad.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SquadUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SquadUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SquadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(squad.Table, squad.Columns, sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(squad.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Public(); ok {
		_spec.SetField(squad.FieldPublic, field.TypeBool, value)
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.SetField(squad.FieldTitle, field.TypeString, value)
	}
	if su.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   squad.MembersTable,
			Columns: squad.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedMembersIDs(); len(nodes) > 0 && !su.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   squad.MembersTable,
			Columns: squad.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   squad.MembersTable,
			Columns: squad.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.InvitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.InvitesTable,
			Columns: []string{squad.InvitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invite.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedInvitesIDs(); len(nodes) > 0 && !su.mutation.InvitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.InvitesTable,
			Columns: []string{squad.InvitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.InvitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.InvitesTable,
			Columns: []string{squad.InvitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.CompetitionResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.CompetitionResultsTable,
			Columns: []string{squad.CompetitionResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competitionresult.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedCompetitionResultsIDs(); len(nodes) > 0 && !su.mutation.CompetitionResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.CompetitionResultsTable,
			Columns: []string{squad.CompetitionResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competitionresult.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CompetitionResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.CompetitionResultsTable,
			Columns: []string{squad.CompetitionResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competitionresult.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   squad.OwnerTable,
			Columns: []string{squad.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   squad.OwnerTable,
			Columns: []string{squad.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{squad.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SquadUpdateOne is the builder for updating a single Squad entity.
type SquadUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SquadMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SquadUpdateOne) SetUpdatedAt(t time.Time) *SquadUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetPublic sets the "public" field.
func (suo *SquadUpdateOne) SetPublic(b bool) *SquadUpdateOne {
	suo.mutation.SetPublic(b)
	return suo
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (suo *SquadUpdateOne) SetNillablePublic(b *bool) *SquadUpdateOne {
	if b != nil {
		suo.SetPublic(*b)
	}
	return suo
}

// SetTitle sets the "title" field.
func (suo *SquadUpdateOne) SetTitle(s string) *SquadUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// AddMemberIDs adds the "members" edge to the Profile entity by IDs.
func (suo *SquadUpdateOne) AddMemberIDs(ids ...uuid.UUID) *SquadUpdateOne {
	suo.mutation.AddMemberIDs(ids...)
	return suo
}

// AddMembers adds the "members" edges to the Profile entity.
func (suo *SquadUpdateOne) AddMembers(p ...*Profile) *SquadUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddMemberIDs(ids...)
}

// AddInviteIDs adds the "invites" edge to the Invite entity by IDs.
func (suo *SquadUpdateOne) AddInviteIDs(ids ...uuid.UUID) *SquadUpdateOne {
	suo.mutation.AddInviteIDs(ids...)
	return suo
}

// AddInvites adds the "invites" edges to the Invite entity.
func (suo *SquadUpdateOne) AddInvites(i ...*Invite) *SquadUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return suo.AddInviteIDs(ids...)
}

// AddCompetitionResultIDs adds the "competition_results" edge to the CompetitionResult entity by IDs.
func (suo *SquadUpdateOne) AddCompetitionResultIDs(ids ...uuid.UUID) *SquadUpdateOne {
	suo.mutation.AddCompetitionResultIDs(ids...)
	return suo
}

// AddCompetitionResults adds the "competition_results" edges to the CompetitionResult entity.
func (suo *SquadUpdateOne) AddCompetitionResults(c ...*CompetitionResult) *SquadUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCompetitionResultIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Profile entity by ID.
func (suo *SquadUpdateOne) SetOwnerID(id uuid.UUID) *SquadUpdateOne {
	suo.mutation.SetOwnerID(id)
	return suo
}

// SetNillableOwnerID sets the "owner" edge to the Profile entity by ID if the given value is not nil.
func (suo *SquadUpdateOne) SetNillableOwnerID(id *uuid.UUID) *SquadUpdateOne {
	if id != nil {
		suo = suo.SetOwnerID(*id)
	}
	return suo
}

// SetOwner sets the "owner" edge to the Profile entity.
func (suo *SquadUpdateOne) SetOwner(p *Profile) *SquadUpdateOne {
	return suo.SetOwnerID(p.ID)
}

// Mutation returns the SquadMutation object of the builder.
func (suo *SquadUpdateOne) Mutation() *SquadMutation {
	return suo.mutation
}

// ClearMembers clears all "members" edges to the Profile entity.
func (suo *SquadUpdateOne) ClearMembers() *SquadUpdateOne {
	suo.mutation.ClearMembers()
	return suo
}

// RemoveMemberIDs removes the "members" edge to Profile entities by IDs.
func (suo *SquadUpdateOne) RemoveMemberIDs(ids ...uuid.UUID) *SquadUpdateOne {
	suo.mutation.RemoveMemberIDs(ids...)
	return suo
}

// RemoveMembers removes "members" edges to Profile entities.
func (suo *SquadUpdateOne) RemoveMembers(p ...*Profile) *SquadUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemoveMemberIDs(ids...)
}

// ClearInvites clears all "invites" edges to the Invite entity.
func (suo *SquadUpdateOne) ClearInvites() *SquadUpdateOne {
	suo.mutation.ClearInvites()
	return suo
}

// RemoveInviteIDs removes the "invites" edge to Invite entities by IDs.
func (suo *SquadUpdateOne) RemoveInviteIDs(ids ...uuid.UUID) *SquadUpdateOne {
	suo.mutation.RemoveInviteIDs(ids...)
	return suo
}

// RemoveInvites removes "invites" edges to Invite entities.
func (suo *SquadUpdateOne) RemoveInvites(i ...*Invite) *SquadUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return suo.RemoveInviteIDs(ids...)
}

// ClearCompetitionResults clears all "competition_results" edges to the CompetitionResult entity.
func (suo *SquadUpdateOne) ClearCompetitionResults() *SquadUpdateOne {
	suo.mutation.ClearCompetitionResults()
	return suo
}

// RemoveCompetitionResultIDs removes the "competition_results" edge to CompetitionResult entities by IDs.
func (suo *SquadUpdateOne) RemoveCompetitionResultIDs(ids ...uuid.UUID) *SquadUpdateOne {
	suo.mutation.RemoveCompetitionResultIDs(ids...)
	return suo
}

// RemoveCompetitionResults removes "competition_results" edges to CompetitionResult entities.
func (suo *SquadUpdateOne) RemoveCompetitionResults(c ...*CompetitionResult) *SquadUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCompetitionResultIDs(ids...)
}

// ClearOwner clears the "owner" edge to the Profile entity.
func (suo *SquadUpdateOne) ClearOwner() *SquadUpdateOne {
	suo.mutation.ClearOwner()
	return suo
}

// Where appends a list predicates to the SquadUpdate builder.
func (suo *SquadUpdateOne) Where(ps ...predicate.Squad) *SquadUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SquadUpdateOne) Select(field string, fields ...string) *SquadUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Squad entity.
func (suo *SquadUpdateOne) Save(ctx context.Context) (*Squad, error) {
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SquadUpdateOne) SaveX(ctx context.Context) *Squad {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SquadUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SquadUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SquadUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		if squad.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized squad.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := squad.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SquadUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SquadUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SquadUpdateOne) sqlSave(ctx context.Context) (_node *Squad, err error) {
	_spec := sqlgraph.NewUpdateSpec(squad.Table, squad.Columns, sqlgraph.NewFieldSpec(squad.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Squad.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, squad.FieldID)
		for _, f := range fields {
			if !squad.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != squad.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(squad.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Public(); ok {
		_spec.SetField(squad.FieldPublic, field.TypeBool, value)
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.SetField(squad.FieldTitle, field.TypeString, value)
	}
	if suo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   squad.MembersTable,
			Columns: squad.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedMembersIDs(); len(nodes) > 0 && !suo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   squad.MembersTable,
			Columns: squad.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   squad.MembersTable,
			Columns: squad.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.InvitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.InvitesTable,
			Columns: []string{squad.InvitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invite.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedInvitesIDs(); len(nodes) > 0 && !suo.mutation.InvitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.InvitesTable,
			Columns: []string{squad.InvitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.InvitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.InvitesTable,
			Columns: []string{squad.InvitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.CompetitionResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.CompetitionResultsTable,
			Columns: []string{squad.CompetitionResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competitionresult.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedCompetitionResultsIDs(); len(nodes) > 0 && !suo.mutation.CompetitionResultsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.CompetitionResultsTable,
			Columns: []string{squad.CompetitionResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competitionresult.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CompetitionResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   squad.CompetitionResultsTable,
			Columns: []string{squad.CompetitionResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(competitionresult.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   squad.OwnerTable,
			Columns: []string{squad.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   squad.OwnerTable,
			Columns: []string{squad.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Squad{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{squad.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
