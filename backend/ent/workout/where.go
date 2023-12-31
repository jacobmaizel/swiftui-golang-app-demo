// Code generated by ent, DO NOT EDIT.

package workout

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldUpdatedAt, v))
}

// HealthkitWorkoutActivityType applies equality check predicate on the "healthkit_workout_activity_type" field. It's identical to HealthkitWorkoutActivityTypeEQ.
func HealthkitWorkoutActivityType(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldHealthkitWorkoutActivityType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldUpdatedAt, v))
}

// HealthkitWorkoutActivityTypeEQ applies the EQ predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeNEQ applies the NEQ predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeNEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeIn applies the In predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldHealthkitWorkoutActivityType, vs...))
}

// HealthkitWorkoutActivityTypeNotIn applies the NotIn predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeNotIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldHealthkitWorkoutActivityType, vs...))
}

// HealthkitWorkoutActivityTypeGT applies the GT predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeGT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeGTE applies the GTE predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeGTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeLT applies the LT predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeLT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeLTE applies the LTE predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeLTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeContains applies the Contains predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeContains(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContains(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeHasPrefix applies the HasPrefix predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeHasPrefix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasPrefix(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeHasSuffix applies the HasSuffix predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeHasSuffix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasSuffix(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeEqualFold applies the EqualFold predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeEqualFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEqualFold(FieldHealthkitWorkoutActivityType, v))
}

// HealthkitWorkoutActivityTypeContainsFold applies the ContainsFold predicate on the "healthkit_workout_activity_type" field.
func HealthkitWorkoutActivityTypeContainsFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContainsFold(FieldHealthkitWorkoutActivityType, v))
}

// HasInvite applies the HasEdge predicate on the "invite" edge.
func HasInvite() predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, InviteTable, InviteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInviteWith applies the HasEdge predicate on the "invite" edge with a given conditions (other predicates).
func HasInviteWith(preds ...predicate.Invite) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := newInviteStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLeader applies the HasEdge predicate on the "leader" edge.
func HasLeader() predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, LeaderTable, LeaderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLeaderWith applies the HasEdge predicate on the "leader" edge with a given conditions (other predicates).
func HasLeaderWith(preds ...predicate.Profile) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := newLeaderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompetition applies the HasEdge predicate on the "competition" edge.
func HasCompetition() predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CompetitionTable, CompetitionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompetitionWith applies the HasEdge predicate on the "competition" edge with a given conditions (other predicates).
func HasCompetitionWith(preds ...predicate.Competition) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := newCompetitionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkoutData applies the HasEdge predicate on the "workout_data" edge.
func HasWorkoutData() predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, WorkoutDataTable, WorkoutDataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutDataWith applies the HasEdge predicate on the "workout_data" edge with a given conditions (other predicates).
func HasWorkoutDataWith(preds ...predicate.WorkoutData) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := newWorkoutDataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkoutRouteData applies the HasEdge predicate on the "workout_route_data" edge.
func HasWorkoutRouteData() predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, WorkoutRouteDataTable, WorkoutRouteDataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutRouteDataWith applies the HasEdge predicate on the "workout_route_data" edge with a given conditions (other predicates).
func HasWorkoutRouteDataWith(preds ...predicate.WorkoutRouteData) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := newWorkoutRouteDataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Workout) predicate.Workout {
	return predicate.Workout(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Workout) predicate.Workout {
	return predicate.Workout(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Workout) predicate.Workout {
	return predicate.Workout(sql.NotPredicates(p))
}
