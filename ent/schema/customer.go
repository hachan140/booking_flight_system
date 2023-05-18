package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
		field.String("email").NotEmpty().Unique().Annotations(entgql.OrderField("EMAIL")),
		field.String("phone_number").NotEmpty().Unique().Annotations(entgql.OrderField("PHONE_NUMBER")),
		field.String("full_name").NotEmpty().Annotations(entgql.OrderField("FULL_NAME")),
		field.Time("dob").Optional().Annotations(entgql.OrderField("DOB")),
		field.String("cid").NotEmpty().Unique().Annotations(entgql.OrderField("CID")),
		field.Int("member_id").Optional().Annotations(entgql.OrderField("MEMBER_ID")),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("has_member", Member.Type).Ref("has_customer").Field("member_id").Unique(),
		edge.To("has_booking", Booking.Type),
	}
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
