// Code generated by ent, DO NOT EDIT.

package booking

import (
	"booking-flight-system/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldUpdatedAt, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCode, v))
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCustomerID, v))
}

// FlightID applies equality check predicate on the "flight_id" field. It's identical to FlightIDEQ.
func FlightID(v int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldFlightID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldUpdatedAt, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContainsFold(FieldCode, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldStatus, vs...))
}

// SeatTypeEQ applies the EQ predicate on the "seat_type" field.
func SeatTypeEQ(v SeatType) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldSeatType, v))
}

// SeatTypeNEQ applies the NEQ predicate on the "seat_type" field.
func SeatTypeNEQ(v SeatType) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldSeatType, v))
}

// SeatTypeIn applies the In predicate on the "seat_type" field.
func SeatTypeIn(vs ...SeatType) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldSeatType, vs...))
}

// SeatTypeNotIn applies the NotIn predicate on the "seat_type" field.
func SeatTypeNotIn(vs ...SeatType) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldSeatType, vs...))
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCustomerID, v))
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v int) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldCustomerID, v))
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...int) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldCustomerID, vs...))
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...int) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldCustomerID, vs...))
}

// CustomerIDIsNil applies the IsNil predicate on the "customer_id" field.
func CustomerIDIsNil() predicate.Booking {
	return predicate.Booking(sql.FieldIsNull(FieldCustomerID))
}

// CustomerIDNotNil applies the NotNil predicate on the "customer_id" field.
func CustomerIDNotNil() predicate.Booking {
	return predicate.Booking(sql.FieldNotNull(FieldCustomerID))
}

// FlightIDEQ applies the EQ predicate on the "flight_id" field.
func FlightIDEQ(v int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldFlightID, v))
}

// FlightIDNEQ applies the NEQ predicate on the "flight_id" field.
func FlightIDNEQ(v int) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldFlightID, v))
}

// FlightIDIn applies the In predicate on the "flight_id" field.
func FlightIDIn(vs ...int) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldFlightID, vs...))
}

// FlightIDNotIn applies the NotIn predicate on the "flight_id" field.
func FlightIDNotIn(vs ...int) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldFlightID, vs...))
}

// FlightIDIsNil applies the IsNil predicate on the "flight_id" field.
func FlightIDIsNil() predicate.Booking {
	return predicate.Booking(sql.FieldIsNull(FieldFlightID))
}

// FlightIDNotNil applies the NotNil predicate on the "flight_id" field.
func FlightIDNotNil() predicate.Booking {
	return predicate.Booking(sql.FieldNotNull(FieldFlightID))
}

// HasHasFlight applies the HasEdge predicate on the "has_flight" edge.
func HasHasFlight() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HasFlightTable, HasFlightColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasFlightWith applies the HasEdge predicate on the "has_flight" edge with a given conditions (other predicates).
func HasHasFlightWith(preds ...predicate.Flight) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newHasFlightStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHasCustomer applies the HasEdge predicate on the "has_customer" edge.
func HasHasCustomer() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HasCustomerTable, HasCustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasCustomerWith applies the HasEdge predicate on the "has_customer" edge with a given conditions (other predicates).
func HasHasCustomerWith(preds ...predicate.Customer) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newHasCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
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
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		p(s.Not())
	})
}
