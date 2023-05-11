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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("email").NotEmpty().Unique(),
		field.String("phone_number").NotEmpty().Unique(),
		field.String("full_name").NotEmpty(),
		field.Time("dob"),
		field.String("cid").NotEmpty().Unique(),
		field.Int("member_id").Optional(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("has_Member", Member.Type).Ref("has_Customer").Field("member_id").Unique(),
		edge.To("has_Flight", Flight.Type),
	}
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
