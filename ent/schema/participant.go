package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Participant holds the schema definition for the Participant entity.
type Participant struct {
	ent.Schema
}

// Fields of the Participant.
func (Participant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("uuid"),

		field.Enum("vocation").
			Values("ek", "ed", "ms", "rp"),

		field.String("discord_id"),
	}
}

// Edges of the Participant.
func (Participant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bosslist", Bosslist.Type).
			Ref("participants"),
	}
}
