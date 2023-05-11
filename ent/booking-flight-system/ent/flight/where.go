// Code generated by ent, DO NOT EDIT.

package flight

import (
	"booking-flight-system/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldName, v))
}

// DepartAt applies equality check predicate on the "depart_at" field. It's identical to DepartAtEQ.
func DepartAt(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDepartAt, v))
}

// LandAt applies equality check predicate on the "land_at" field. It's identical to LandAtEQ.
func LandAt(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldLandAt, v))
}

// AvailableEcSlot applies equality check predicate on the "available_ec_slot" field. It's identical to AvailableEcSlotEQ.
func AvailableEcSlot(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAvailableEcSlot, v))
}

// AvailableBcSlot applies equality check predicate on the "available_bc_slot" field. It's identical to AvailableBcSlotEQ.
func AvailableBcSlot(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAvailableBcSlot, v))
}

// PlaneID applies equality check predicate on the "plane_id" field. It's identical to PlaneIDEQ.
func PlaneID(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldPlaneID, v))
}

// AirportID applies equality check predicate on the "airport_id" field. It's identical to AirportIDEQ.
func AirportID(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAirportID, v))
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldCustomerID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Flight {
	return predicate.Flight(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Flight {
	return predicate.Flight(sql.FieldContainsFold(FieldName, v))
}

// DepartAtEQ applies the EQ predicate on the "depart_at" field.
func DepartAtEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldDepartAt, v))
}

// DepartAtNEQ applies the NEQ predicate on the "depart_at" field.
func DepartAtNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldDepartAt, v))
}

// DepartAtIn applies the In predicate on the "depart_at" field.
func DepartAtIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldDepartAt, vs...))
}

// DepartAtNotIn applies the NotIn predicate on the "depart_at" field.
func DepartAtNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldDepartAt, vs...))
}

// DepartAtGT applies the GT predicate on the "depart_at" field.
func DepartAtGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldDepartAt, v))
}

// DepartAtGTE applies the GTE predicate on the "depart_at" field.
func DepartAtGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldDepartAt, v))
}

// DepartAtLT applies the LT predicate on the "depart_at" field.
func DepartAtLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldDepartAt, v))
}

// DepartAtLTE applies the LTE predicate on the "depart_at" field.
func DepartAtLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldDepartAt, v))
}

// LandAtEQ applies the EQ predicate on the "land_at" field.
func LandAtEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldLandAt, v))
}

// LandAtNEQ applies the NEQ predicate on the "land_at" field.
func LandAtNEQ(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldLandAt, v))
}

// LandAtIn applies the In predicate on the "land_at" field.
func LandAtIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldLandAt, vs...))
}

// LandAtNotIn applies the NotIn predicate on the "land_at" field.
func LandAtNotIn(vs ...time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldLandAt, vs...))
}

// LandAtGT applies the GT predicate on the "land_at" field.
func LandAtGT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldLandAt, v))
}

// LandAtGTE applies the GTE predicate on the "land_at" field.
func LandAtGTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldLandAt, v))
}

// LandAtLT applies the LT predicate on the "land_at" field.
func LandAtLT(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldLandAt, v))
}

// LandAtLTE applies the LTE predicate on the "land_at" field.
func LandAtLTE(v time.Time) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldLandAt, v))
}

// AvailableEcSlotEQ applies the EQ predicate on the "available_ec_slot" field.
func AvailableEcSlotEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAvailableEcSlot, v))
}

// AvailableEcSlotNEQ applies the NEQ predicate on the "available_ec_slot" field.
func AvailableEcSlotNEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldAvailableEcSlot, v))
}

// AvailableEcSlotIn applies the In predicate on the "available_ec_slot" field.
func AvailableEcSlotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldAvailableEcSlot, vs...))
}

// AvailableEcSlotNotIn applies the NotIn predicate on the "available_ec_slot" field.
func AvailableEcSlotNotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldAvailableEcSlot, vs...))
}

// AvailableEcSlotGT applies the GT predicate on the "available_ec_slot" field.
func AvailableEcSlotGT(v int) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldAvailableEcSlot, v))
}

// AvailableEcSlotGTE applies the GTE predicate on the "available_ec_slot" field.
func AvailableEcSlotGTE(v int) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldAvailableEcSlot, v))
}

// AvailableEcSlotLT applies the LT predicate on the "available_ec_slot" field.
func AvailableEcSlotLT(v int) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldAvailableEcSlot, v))
}

// AvailableEcSlotLTE applies the LTE predicate on the "available_ec_slot" field.
func AvailableEcSlotLTE(v int) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldAvailableEcSlot, v))
}

// AvailableBcSlotEQ applies the EQ predicate on the "available_bc_slot" field.
func AvailableBcSlotEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAvailableBcSlot, v))
}

// AvailableBcSlotNEQ applies the NEQ predicate on the "available_bc_slot" field.
func AvailableBcSlotNEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldAvailableBcSlot, v))
}

// AvailableBcSlotIn applies the In predicate on the "available_bc_slot" field.
func AvailableBcSlotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldAvailableBcSlot, vs...))
}

// AvailableBcSlotNotIn applies the NotIn predicate on the "available_bc_slot" field.
func AvailableBcSlotNotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldAvailableBcSlot, vs...))
}

// AvailableBcSlotGT applies the GT predicate on the "available_bc_slot" field.
func AvailableBcSlotGT(v int) predicate.Flight {
	return predicate.Flight(sql.FieldGT(FieldAvailableBcSlot, v))
}

// AvailableBcSlotGTE applies the GTE predicate on the "available_bc_slot" field.
func AvailableBcSlotGTE(v int) predicate.Flight {
	return predicate.Flight(sql.FieldGTE(FieldAvailableBcSlot, v))
}

// AvailableBcSlotLT applies the LT predicate on the "available_bc_slot" field.
func AvailableBcSlotLT(v int) predicate.Flight {
	return predicate.Flight(sql.FieldLT(FieldAvailableBcSlot, v))
}

// AvailableBcSlotLTE applies the LTE predicate on the "available_bc_slot" field.
func AvailableBcSlotLTE(v int) predicate.Flight {
	return predicate.Flight(sql.FieldLTE(FieldAvailableBcSlot, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldStatus, vs...))
}

// PlaneIDEQ applies the EQ predicate on the "plane_id" field.
func PlaneIDEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldPlaneID, v))
}

// PlaneIDNEQ applies the NEQ predicate on the "plane_id" field.
func PlaneIDNEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldPlaneID, v))
}

// PlaneIDIn applies the In predicate on the "plane_id" field.
func PlaneIDIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldPlaneID, vs...))
}

// PlaneIDNotIn applies the NotIn predicate on the "plane_id" field.
func PlaneIDNotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldPlaneID, vs...))
}

// PlaneIDIsNil applies the IsNil predicate on the "plane_id" field.
func PlaneIDIsNil() predicate.Flight {
	return predicate.Flight(sql.FieldIsNull(FieldPlaneID))
}

// PlaneIDNotNil applies the NotNil predicate on the "plane_id" field.
func PlaneIDNotNil() predicate.Flight {
	return predicate.Flight(sql.FieldNotNull(FieldPlaneID))
}

// AirportIDEQ applies the EQ predicate on the "airport_id" field.
func AirportIDEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldAirportID, v))
}

// AirportIDNEQ applies the NEQ predicate on the "airport_id" field.
func AirportIDNEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldAirportID, v))
}

// AirportIDIn applies the In predicate on the "airport_id" field.
func AirportIDIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldAirportID, vs...))
}

// AirportIDNotIn applies the NotIn predicate on the "airport_id" field.
func AirportIDNotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldAirportID, vs...))
}

// AirportIDIsNil applies the IsNil predicate on the "airport_id" field.
func AirportIDIsNil() predicate.Flight {
	return predicate.Flight(sql.FieldIsNull(FieldAirportID))
}

// AirportIDNotNil applies the NotNil predicate on the "airport_id" field.
func AirportIDNotNil() predicate.Flight {
	return predicate.Flight(sql.FieldNotNull(FieldAirportID))
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldEQ(FieldCustomerID, v))
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v int) predicate.Flight {
	return predicate.Flight(sql.FieldNEQ(FieldCustomerID, v))
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldIn(FieldCustomerID, vs...))
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...int) predicate.Flight {
	return predicate.Flight(sql.FieldNotIn(FieldCustomerID, vs...))
}

// CustomerIDIsNil applies the IsNil predicate on the "customer_id" field.
func CustomerIDIsNil() predicate.Flight {
	return predicate.Flight(sql.FieldIsNull(FieldCustomerID))
}

// CustomerIDNotNil applies the NotNil predicate on the "customer_id" field.
func CustomerIDNotNil() predicate.Flight {
	return predicate.Flight(sql.FieldNotNull(FieldCustomerID))
}

// HasHasPlane applies the HasEdge predicate on the "has_plane" edge.
func HasHasPlane() predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HasPlaneTable, HasPlaneColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasPlaneWith applies the HasEdge predicate on the "has_plane" edge with a given conditions (other predicates).
func HasHasPlaneWith(preds ...predicate.Plane) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := newHasPlaneStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHasBooking applies the HasEdge predicate on the "has_booking" edge.
func HasHasBooking() predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HasBookingTable, HasBookingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasBookingWith applies the HasEdge predicate on the "has_booking" edge with a given conditions (other predicates).
func HasHasBookingWith(preds ...predicate.Booking) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := newHasBookingStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHasAirport applies the HasEdge predicate on the "has_Airport" edge.
func HasHasAirport() predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HasAirportTable, HasAirportColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasAirportWith applies the HasEdge predicate on the "has_Airport" edge with a given conditions (other predicates).
func HasHasAirportWith(preds ...predicate.Airport) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := newHasAirportStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHasCustomer applies the HasEdge predicate on the "has_Customer" edge.
func HasHasCustomer() predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HasCustomerTable, HasCustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasCustomerWith applies the HasEdge predicate on the "has_Customer" edge with a given conditions (other predicates).
func HasHasCustomerWith(preds ...predicate.Customer) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		step := newHasCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
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
func Not(p predicate.Flight) predicate.Flight {
	return predicate.Flight(func(s *sql.Selector) {
		p(s.Not())
	})
}
