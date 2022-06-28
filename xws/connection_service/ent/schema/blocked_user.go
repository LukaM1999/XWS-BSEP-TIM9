package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// BlockedUser holds the schema definition for the Block entity.
type BlockedUser struct {
	ent.Schema
}

// Fields of the Block.
func (BlockedUser) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.String("issuer_primary_key"),
		field.String("subject_primary_key"),
	}
}

// Edges of the Block.
func (BlockedUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("blocked_by", User.Type).
			Required().
			Unique().
			Ref("block"),
	}
}
