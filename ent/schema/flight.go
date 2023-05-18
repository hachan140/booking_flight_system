package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.Time("depart_at").Annotations(entgql.OrderField("DEPART_AT")),
		field.Time("land_at").Annotations(entgql.OrderField("LAND_AT")),
		field.Int("available_ec_slot").Nillable().Positive().Annotations(entgql.OrderField("AVAILABLE_EC_SLOT")),
		field.Int("available_bc_slot").Nillable().Positive().Annotations(entgql.OrderField("AVAILABLE_BC_SLOT")),
		field.Enum("status").NamedValues("scheduled", "SCHEDULED", "canceled", "CANCELED", "delay", "DELAY", "landed", "LANDED", "flying", "FLYING").Default("SCHEDULED").Annotations(entgql.OrderField("STATUS")),
		field.Int("plane_id").Optional().Annotations(entgql.OrderField("PLANE_ID")),
		field.Int("from_airport_id").Optional().Annotations(entgql.OrderField("FROM_AIRPORT_ID")),
		field.Int("to_airport_id").Optional().Annotations(entgql.OrderField("TO_AIRPORT_ID")),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("has_plane", Plane.Type).Ref("flights").Field("plane_id").Unique(),
		edge.To("has_booking", Booking.Type),
		edge.From("from_airport", Airport.Type).Ref("from_flight").Field("from_airport_id").Unique(),
		edge.From("to_airport", Airport.Type).Ref("to_flight").Field("to_airport_id").Unique(),
	}
}

func (Flight) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
