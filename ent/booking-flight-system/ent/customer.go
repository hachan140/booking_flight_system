// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/customer"
	"booking-flight-system/ent/member"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// Dob holds the value of the "dob" field.
	Dob time.Time `json:"dob,omitempty"`
	// Cid holds the value of the "cid" field.
	Cid string `json:"cid,omitempty"`
	// MemberID holds the value of the "member_id" field.
	MemberID int `json:"member_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges        CustomerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// HasMember holds the value of the has_Member edge.
	HasMember *Member `json:"has_Member,omitempty"`
	// HasFlight holds the value of the has_Flight edge.
	HasFlight []*Flight `json:"has_Flight,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedHasFlight map[string][]*Flight
}

// HasMemberOrErr returns the HasMember value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) HasMemberOrErr() (*Member, error) {
	if e.loadedTypes[0] {
		if e.HasMember == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.HasMember, nil
	}
	return nil, &NotLoadedError{edge: "has_Member"}
}

// HasFlightOrErr returns the HasFlight value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) HasFlightOrErr() ([]*Flight, error) {
	if e.loadedTypes[1] {
		return e.HasFlight, nil
	}
	return nil, &NotLoadedError{edge: "has_Flight"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldID, customer.FieldMemberID:
			values[i] = new(sql.NullInt64)
		case customer.FieldEmail, customer.FieldPhoneNumber, customer.FieldFullName, customer.FieldCid:
			values[i] = new(sql.NullString)
		case customer.FieldCreatedAt, customer.FieldUpdatedAt, customer.FieldDob:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case customer.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case customer.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				c.PhoneNumber = value.String
			}
		case customer.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				c.FullName = value.String
			}
		case customer.FieldDob:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dob", values[i])
			} else if value.Valid {
				c.Dob = value.Time
			}
		case customer.FieldCid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cid", values[i])
			} else if value.Valid {
				c.Cid = value.String
			}
		case customer.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				c.MemberID = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Customer.
// This includes values selected through modifiers, order, etc.
func (c *Customer) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryHasMember queries the "has_Member" edge of the Customer entity.
func (c *Customer) QueryHasMember() *MemberQuery {
	return NewCustomerClient(c.config).QueryHasMember(c)
}

// QueryHasFlight queries the "has_Flight" edge of the Customer entity.
func (c *Customer) QueryHasFlight() *FlightQuery {
	return NewCustomerClient(c.config).QueryHasFlight(c)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return NewCustomerClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(c.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(c.FullName)
	builder.WriteString(", ")
	builder.WriteString("dob=")
	builder.WriteString(c.Dob.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("cid=")
	builder.WriteString(c.Cid)
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", c.MemberID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedHasFlight returns the HasFlight named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Customer) NamedHasFlight(name string) ([]*Flight, error) {
	if c.Edges.namedHasFlight == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedHasFlight[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Customer) appendNamedHasFlight(name string, edges ...*Flight) {
	if c.Edges.namedHasFlight == nil {
		c.Edges.namedHasFlight = make(map[string][]*Flight)
	}
	if len(edges) == 0 {
		c.Edges.namedHasFlight[name] = []*Flight{}
	} else {
		c.Edges.namedHasFlight[name] = append(c.Edges.namedHasFlight[name], edges...)
	}
}

// Customers is a parsable slice of Customer.
type Customers []*Customer
