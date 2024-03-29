// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/airport"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Airport is the model entity for the Airport schema.
type Airport struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Lat holds the value of the "lat" field.
	Lat *float64 `json:"lat,omitempty"`
	// Long holds the value of the "long" field.
	Long *float64 `json:"long,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AirportQuery when eager-loading is set.
	Edges        AirportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AirportEdges holds the relations/edges for other nodes in the graph.
type AirportEdges struct {
	// FromFlight holds the value of the from_flight edge.
	FromFlight []*Flight `json:"from_flight,omitempty"`
	// ToFlight holds the value of the to_flight edge.
	ToFlight []*Flight `json:"to_flight,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedFromFlight map[string][]*Flight
	namedToFlight   map[string][]*Flight
}

// FromFlightOrErr returns the FromFlight value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) FromFlightOrErr() ([]*Flight, error) {
	if e.loadedTypes[0] {
		return e.FromFlight, nil
	}
	return nil, &NotLoadedError{edge: "from_flight"}
}

// ToFlightOrErr returns the ToFlight value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) ToFlightOrErr() ([]*Flight, error) {
	if e.loadedTypes[1] {
		return e.ToFlight, nil
	}
	return nil, &NotLoadedError{edge: "to_flight"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Airport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case airport.FieldLat, airport.FieldLong:
			values[i] = new(sql.NullFloat64)
		case airport.FieldID:
			values[i] = new(sql.NullInt64)
		case airport.FieldName:
			values[i] = new(sql.NullString)
		case airport.FieldCreatedAt, airport.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Airport fields.
func (a *Airport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case airport.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case airport.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case airport.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case airport.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case airport.FieldLat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lat", values[i])
			} else if value.Valid {
				a.Lat = new(float64)
				*a.Lat = value.Float64
			}
		case airport.FieldLong:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field long", values[i])
			} else if value.Valid {
				a.Long = new(float64)
				*a.Long = value.Float64
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Airport.
// This includes values selected through modifiers, order, etc.
func (a *Airport) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryFromFlight queries the "from_flight" edge of the Airport entity.
func (a *Airport) QueryFromFlight() *FlightQuery {
	return NewAirportClient(a.config).QueryFromFlight(a)
}

// QueryToFlight queries the "to_flight" edge of the Airport entity.
func (a *Airport) QueryToFlight() *FlightQuery {
	return NewAirportClient(a.config).QueryToFlight(a)
}

// Update returns a builder for updating this Airport.
// Note that you need to call Airport.Unwrap() before calling this method if this Airport
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Airport) Update() *AirportUpdateOne {
	return NewAirportClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Airport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Airport) Unwrap() *Airport {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Airport is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Airport) String() string {
	var builder strings.Builder
	builder.WriteString("Airport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	if v := a.Lat; v != nil {
		builder.WriteString("lat=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := a.Long; v != nil {
		builder.WriteString("long=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedFromFlight returns the FromFlight named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Airport) NamedFromFlight(name string) ([]*Flight, error) {
	if a.Edges.namedFromFlight == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedFromFlight[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Airport) appendNamedFromFlight(name string, edges ...*Flight) {
	if a.Edges.namedFromFlight == nil {
		a.Edges.namedFromFlight = make(map[string][]*Flight)
	}
	if len(edges) == 0 {
		a.Edges.namedFromFlight[name] = []*Flight{}
	} else {
		a.Edges.namedFromFlight[name] = append(a.Edges.namedFromFlight[name], edges...)
	}
}

// NamedToFlight returns the ToFlight named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Airport) NamedToFlight(name string) ([]*Flight, error) {
	if a.Edges.namedToFlight == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedToFlight[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Airport) appendNamedToFlight(name string, edges ...*Flight) {
	if a.Edges.namedToFlight == nil {
		a.Edges.namedToFlight = make(map[string][]*Flight)
	}
	if len(edges) == 0 {
		a.Edges.namedToFlight[name] = []*Flight{}
	} else {
		a.Edges.namedToFlight[name] = append(a.Edges.namedToFlight[name], edges...)
	}
}

// Airports is a parsable slice of Airport.
type Airports []*Airport
