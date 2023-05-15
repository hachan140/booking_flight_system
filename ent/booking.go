// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/booking"
	"booking-flight-system/ent/customer"
	"booking-flight-system/ent/flight"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Booking is the model entity for the Booking schema.
type Booking struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Status holds the value of the "status" field.
	Status booking.Status `json:"status,omitempty"`
	// SeatType holds the value of the "seat_type" field.
	SeatType booking.SeatType `json:"seat_type,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID int `json:"customer_id,omitempty"`
	// FlightID holds the value of the "flight_id" field.
	FlightID int `json:"flight_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookingQuery when eager-loading is set.
	Edges        BookingEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BookingEdges holds the relations/edges for other nodes in the graph.
type BookingEdges struct {
	// HasFlight holds the value of the has_flight edge.
	HasFlight *Flight `json:"has_flight,omitempty"`
	// HasCustomer holds the value of the has_customer edge.
	HasCustomer *Customer `json:"has_customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// HasFlightOrErr returns the HasFlight value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) HasFlightOrErr() (*Flight, error) {
	if e.loadedTypes[0] {
		if e.HasFlight == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: flight.Label}
		}
		return e.HasFlight, nil
	}
	return nil, &NotLoadedError{edge: "has_flight"}
}

// HasCustomerOrErr returns the HasCustomer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) HasCustomerOrErr() (*Customer, error) {
	if e.loadedTypes[1] {
		if e.HasCustomer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.HasCustomer, nil
	}
	return nil, &NotLoadedError{edge: "has_customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Booking) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case booking.FieldID, booking.FieldCustomerID, booking.FieldFlightID:
			values[i] = new(sql.NullInt64)
		case booking.FieldCode, booking.FieldStatus, booking.FieldSeatType:
			values[i] = new(sql.NullString)
		case booking.FieldCreatedAt, booking.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Booking fields.
func (b *Booking) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case booking.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case booking.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case booking.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case booking.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				b.Code = value.String
			}
		case booking.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = booking.Status(value.String)
			}
		case booking.FieldSeatType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seat_type", values[i])
			} else if value.Valid {
				b.SeatType = booking.SeatType(value.String)
			}
		case booking.FieldCustomerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				b.CustomerID = int(value.Int64)
			}
		case booking.FieldFlightID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field flight_id", values[i])
			} else if value.Valid {
				b.FlightID = int(value.Int64)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Booking.
// This includes values selected through modifiers, order, etc.
func (b *Booking) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryHasFlight queries the "has_flight" edge of the Booking entity.
func (b *Booking) QueryHasFlight() *FlightQuery {
	return NewBookingClient(b.config).QueryHasFlight(b)
}

// QueryHasCustomer queries the "has_customer" edge of the Booking entity.
func (b *Booking) QueryHasCustomer() *CustomerQuery {
	return NewBookingClient(b.config).QueryHasCustomer(b)
}

// Update returns a builder for updating this Booking.
// Note that you need to call Booking.Unwrap() before calling this method if this Booking
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Booking) Update() *BookingUpdateOne {
	return NewBookingClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Booking entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Booking) Unwrap() *Booking {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Booking is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Booking) String() string {
	var builder strings.Builder
	builder.WriteString("Booking(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(b.Code)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("seat_type=")
	builder.WriteString(fmt.Sprintf("%v", b.SeatType))
	builder.WriteString(", ")
	builder.WriteString("customer_id=")
	builder.WriteString(fmt.Sprintf("%v", b.CustomerID))
	builder.WriteString(", ")
	builder.WriteString("flight_id=")
	builder.WriteString(fmt.Sprintf("%v", b.FlightID))
	builder.WriteByte(')')
	return builder.String()
}

// Bookings is a parsable slice of Booking.
type Bookings []*Booking
