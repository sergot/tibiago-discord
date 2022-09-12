package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Boss holds the schema definition for the Boss entity.
type Boss struct {
	ent.Schema
}

// Fields of the Boss.
func (Boss) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("uuid"),

		field.String("name").
			Unique(),

		field.String("template").
			Default("1ek1ed3shooter"),
	}
}

// Edges of the Boss.
func (Boss) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bosslists", Bosslist.Type),
	}
}
