// Code generated by ent, DO NOT EDIT.

package plane

import (
	"booking-flight-system/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Plane {
	return predicate.Plane(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Plane {
	return predicate.Plane(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Plane {
	return predicate.Plane(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Plane {
	return predicate.Plane(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Plane {
	return predicate.Plane(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Plane {
	return predicate.Plane(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Plane {
	return predicate.Plane(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldName, v))
}

// EconomyClassSlots applies equality check predicate on the "economy_class_slots" field. It's identical to EconomyClassSlotsEQ.
func EconomyClassSlots(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldEconomyClassSlots, v))
}

// BusinessClassSlots applies equality check predicate on the "business_class_slots" field. It's identical to BusinessClassSlotsEQ.
func BusinessClassSlots(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldBusinessClassSlots, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Plane {
	return predicate.Plane(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Plane {
	return predicate.Plane(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Plane {
	return predicate.Plane(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Plane {
	return predicate.Plane(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Plane {
	return predicate.Plane(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Plane {
	return predicate.Plane(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Plane {
	return predicate.Plane(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Plane {
	return predicate.Plane(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Plane {
	return predicate.Plane(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Plane {
	return predicate.Plane(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Plane {
	return predicate.Plane(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Plane {
	return predicate.Plane(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Plane {
	return predicate.Plane(sql.FieldContainsFold(FieldName, v))
}

// EconomyClassSlotsEQ applies the EQ predicate on the "economy_class_slots" field.
func EconomyClassSlotsEQ(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldEconomyClassSlots, v))
}

// EconomyClassSlotsNEQ applies the NEQ predicate on the "economy_class_slots" field.
func EconomyClassSlotsNEQ(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldNEQ(FieldEconomyClassSlots, v))
}

// EconomyClassSlotsIn applies the In predicate on the "economy_class_slots" field.
func EconomyClassSlotsIn(vs ...int64) predicate.Plane {
	return predicate.Plane(sql.FieldIn(FieldEconomyClassSlots, vs...))
}

// EconomyClassSlotsNotIn applies the NotIn predicate on the "economy_class_slots" field.
func EconomyClassSlotsNotIn(vs ...int64) predicate.Plane {
	return predicate.Plane(sql.FieldNotIn(FieldEconomyClassSlots, vs...))
}

// EconomyClassSlotsGT applies the GT predicate on the "economy_class_slots" field.
func EconomyClassSlotsGT(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldGT(FieldEconomyClassSlots, v))
}

// EconomyClassSlotsGTE applies the GTE predicate on the "economy_class_slots" field.
func EconomyClassSlotsGTE(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldGTE(FieldEconomyClassSlots, v))
}

// EconomyClassSlotsLT applies the LT predicate on the "economy_class_slots" field.
func EconomyClassSlotsLT(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldLT(FieldEconomyClassSlots, v))
}

// EconomyClassSlotsLTE applies the LTE predicate on the "economy_class_slots" field.
func EconomyClassSlotsLTE(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldLTE(FieldEconomyClassSlots, v))
}

// BusinessClassSlotsEQ applies the EQ predicate on the "business_class_slots" field.
func BusinessClassSlotsEQ(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldBusinessClassSlots, v))
}

// BusinessClassSlotsNEQ applies the NEQ predicate on the "business_class_slots" field.
func BusinessClassSlotsNEQ(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldNEQ(FieldBusinessClassSlots, v))
}

// BusinessClassSlotsIn applies the In predicate on the "business_class_slots" field.
func BusinessClassSlotsIn(vs ...int64) predicate.Plane {
	return predicate.Plane(sql.FieldIn(FieldBusinessClassSlots, vs...))
}

// BusinessClassSlotsNotIn applies the NotIn predicate on the "business_class_slots" field.
func BusinessClassSlotsNotIn(vs ...int64) predicate.Plane {
	return predicate.Plane(sql.FieldNotIn(FieldBusinessClassSlots, vs...))
}

// BusinessClassSlotsGT applies the GT predicate on the "business_class_slots" field.
func BusinessClassSlotsGT(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldGT(FieldBusinessClassSlots, v))
}

// BusinessClassSlotsGTE applies the GTE predicate on the "business_class_slots" field.
func BusinessClassSlotsGTE(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldGTE(FieldBusinessClassSlots, v))
}

// BusinessClassSlotsLT applies the LT predicate on the "business_class_slots" field.
func BusinessClassSlotsLT(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldLT(FieldBusinessClassSlots, v))
}

// BusinessClassSlotsLTE applies the LTE predicate on the "business_class_slots" field.
func BusinessClassSlotsLTE(v int64) predicate.Plane {
	return predicate.Plane(sql.FieldLTE(FieldBusinessClassSlots, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Plane {
	return predicate.Plane(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Plane {
	return predicate.Plane(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Plane {
	return predicate.Plane(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Plane {
	return predicate.Plane(sql.FieldNotIn(FieldStatus, vs...))
}

// HasFlights applies the HasEdge predicate on the "flights" edge.
func HasFlights() predicate.Plane {
	return predicate.Plane(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FlightsTable, FlightsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFlightsWith applies the HasEdge predicate on the "flights" edge with a given conditions (other predicates).
func HasFlightsWith(preds ...predicate.Flight) predicate.Plane {
	return predicate.Plane(func(s *sql.Selector) {
		step := newFlightsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Plane) predicate.Plane {
	return predicate.Plane(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Plane) predicate.Plane {
	return predicate.Plane(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Plane) predicate.Plane {
	return predicate.Plane(func(s *sql.Selector) {
		p(s.Not())
	})
}
