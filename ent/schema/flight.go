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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("name").NotEmpty(),
		field.Time("depart_at"),
		field.Time("land_at"),
		field.Int("available_ec_slot").Nillable(),
		field.Int("available_bc_slot").Nillable(),
		field.Enum("status").Values("flying", "scheduled", "canceled", "landed").Default("landed"),
		field.Int("plane_id").Optional(),
		field.Int("from_airport_id").Optional(),
		field.Int("to_airport_id").Optional(),
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
