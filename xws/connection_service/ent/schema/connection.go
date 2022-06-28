package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Connection holds the schema definition for the Connection entity.
type Connection struct {
	ent.Schema
}

// Fields of the Connection.
func (Connection) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Bool("is_approved").
			Default(false),
		field.Int("user_id"),
		field.Int("connection_id"),
		field.String("issuer_primary_key"),
		field.String("subject_primary_key"),
	}
}

// Edges of the Connection.
func (Connection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("connection", User.Type).
			Required().
			Unique().
			Field("connection_id"),
	}
}
