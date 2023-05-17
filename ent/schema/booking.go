package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("code").NotEmpty().Unique(),
		field.Enum("status").NamedValues("success", "SUCCESS", "cancel", "CANCEL"),
		field.Enum("seat_type").NamedValues("economic_class", "EC", "business_class", "BC"),
		field.Bool("is_round").Default(false),
		field.Int("customer_id").Optional(),
		field.Int("flight_id").Optional(),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("has_flight", Flight.Type).Ref("has_booking").Field("flight_id").Unique(),
		edge.From("has_customer", Customer.Type).Ref("has_booking").Field("customer_id").Unique(),
	}
}

func (Booking) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
