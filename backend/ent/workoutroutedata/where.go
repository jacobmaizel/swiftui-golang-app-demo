// Code generated by ent, DO NOT EDIT.

package workoutroutedata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasWorkout applies the HasEdge predicate on the "workout" edge.
func HasWorkout() predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, WorkoutTable, WorkoutColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutWith applies the HasEdge predicate on the "workout" edge with a given conditions (other predicates).
func HasWorkoutWith(preds ...predicate.Workout) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(func(s *sql.Selector) {
		step := newWorkoutStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkoutData applies the HasEdge predicate on the "workout_data" edge.
func HasWorkoutData() predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, WorkoutDataTable, WorkoutDataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutDataWith applies the HasEdge predicate on the "workout_data" edge with a given conditions (other predicates).
func HasWorkoutDataWith(preds ...predicate.WorkoutData) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(func(s *sql.Selector) {
		step := newWorkoutDataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkoutRouteData) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkoutRouteData) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WorkoutRouteData) predicate.WorkoutRouteData {
	return predicate.WorkoutRouteData(sql.NotPredicates(p))
}