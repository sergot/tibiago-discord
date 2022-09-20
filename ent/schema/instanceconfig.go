package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// InstanceConfig holds the schema definition for the InstanceConfig entity.
type InstanceConfig struct {
	ent.Schema
}

// Fields of the InstanceConfig.
func (InstanceConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("uuid"),

		field.String("key"),
		field.String("value"),
	}
}

// Edges of the InstanceConfig.
func (InstanceConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("instance", Instance.Type).
			Ref("configs").
			Unique(),
	}
}

func (InstanceConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").
			Edges("instance").
			Unique(),
	}
}
