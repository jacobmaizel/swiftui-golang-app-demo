// Code generated by ent, DO NOT EDIT.

package competition

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldUpdatedAt, v))
}

// Public applies equality check predicate on the "public" field. It's identical to PublicEQ.
func Public(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldPublic, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldDescription, v))
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldStart, v))
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldEnd, v))
}

// Scheduled applies equality check predicate on the "scheduled" field. It's identical to ScheduledEQ.
func Scheduled(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldScheduled, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldStatus, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldUpdatedAt, v))
}

// PublicEQ applies the EQ predicate on the "public" field.
func PublicEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldPublic, v))
}

// PublicNEQ applies the NEQ predicate on the "public" field.
func PublicNEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldPublic, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContainsFold(FieldDescription, v))
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldStart, v))
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldStart, v))
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldStart, vs...))
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldStart, vs...))
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldStart, v))
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldStart, v))
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldStart, v))
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldStart, v))
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldEnd, v))
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldEnd, v))
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldEnd, vs...))
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldEnd, vs...))
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldEnd, v))
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldEnd, v))
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldEnd, v))
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldEnd, v))
}

// ScheduledEQ applies the EQ predicate on the "scheduled" field.
func ScheduledEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldScheduled, v))
}

// ScheduledNEQ applies the NEQ predicate on the "scheduled" field.
func ScheduledNEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldScheduled, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContainsFold(FieldStatus, v))
}

// ParticipantTypesIsNil applies the IsNil predicate on the "participant_types" field.
func ParticipantTypesIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldParticipantTypes))
}

// ParticipantTypesNotNil applies the NotNil predicate on the "participant_types" field.
func ParticipantTypesNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldParticipantTypes))
}

// WorkoutTypesIsNil applies the IsNil predicate on the "workout_types" field.
func WorkoutTypesIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldWorkoutTypes))
}

// WorkoutTypesNotNil applies the NotNil predicate on the "workout_types" field.
func WorkoutTypesNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldWorkoutTypes))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContainsFold(FieldType, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Profile) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParticipants applies the HasEdge predicate on the "participants" edge.
func HasParticipants() predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ParticipantsTable, ParticipantsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParticipantsWith applies the HasEdge predicate on the "participants" edge with a given conditions (other predicates).
func HasParticipantsWith(preds ...predicate.Profile) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := newParticipantsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkouts applies the HasEdge predicate on the "workouts" edge.
func HasWorkouts() predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, WorkoutsTable, WorkoutsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutsWith applies the HasEdge predicate on the "workouts" edge with a given conditions (other predicates).
func HasWorkoutsWith(preds ...predicate.Workout) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := newWorkoutsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkoutData applies the HasEdge predicate on the "workout_data" edge.
func HasWorkoutData() predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, WorkoutDataTable, WorkoutDataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutDataWith applies the HasEdge predicate on the "workout_data" edge with a given conditions (other predicates).
func HasWorkoutDataWith(preds ...predicate.WorkoutData) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := newWorkoutDataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResults applies the HasEdge predicate on the "results" edge.
func HasResults() predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ResultsTable, ResultsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResultsWith applies the HasEdge predicate on the "results" edge with a given conditions (other predicates).
func HasResultsWith(preds ...predicate.CompetitionResult) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := newResultsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Competition) predicate.Competition {
	return predicate.Competition(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Competition) predicate.Competition {
	return predicate.Competition(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Competition) predicate.Competition {
	return predicate.Competition(sql.NotPredicates(p))
}
