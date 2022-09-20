package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("uuid"),

		field.Enum("status").
			Values("active", "inactive"), // TODO: more statuses

		field.String("discord_guild_id"),
	}
}

// Edges of the Instance.
func (Instance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("configs", InstanceConfig.Type),
	}
}
