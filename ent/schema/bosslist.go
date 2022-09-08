package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Bosslist holds the schema definition for the Bosslist entity.
type Bosslist struct {
	ent.Schema
}

// Fields of the Bosslist.
func (Bosslist) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("uuid"),

		field.Time("created_at").
			Default(time.Now),

		field.Time("starts_at"),

		field.String("custom_template").
			Optional(),

		field.String("discord_message_id").
			Optional(),
	}
}

// Edges of the Bosslist.
func (Bosslist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("boss", Boss.Type).
			Ref("bosslists").
			Unique(),
		edge.To("participants", Participant.Type),
	}
}
