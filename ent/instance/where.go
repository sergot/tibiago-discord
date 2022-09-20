// Code generated by ent, DO NOT EDIT.

package instance

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/sergot/tibiago/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DiscordGuildID applies equality check predicate on the "discord_guild_id" field. It's identical to DiscordGuildIDEQ.
func DiscordGuildID(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscordGuildID), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// DiscordGuildIDEQ applies the EQ predicate on the "discord_guild_id" field.
func DiscordGuildIDEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDNEQ applies the NEQ predicate on the "discord_guild_id" field.
func DiscordGuildIDNEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDIn applies the In predicate on the "discord_guild_id" field.
func DiscordGuildIDIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDiscordGuildID), v...))
	})
}

// DiscordGuildIDNotIn applies the NotIn predicate on the "discord_guild_id" field.
func DiscordGuildIDNotIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDiscordGuildID), v...))
	})
}

// DiscordGuildIDGT applies the GT predicate on the "discord_guild_id" field.
func DiscordGuildIDGT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDGTE applies the GTE predicate on the "discord_guild_id" field.
func DiscordGuildIDGTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDLT applies the LT predicate on the "discord_guild_id" field.
func DiscordGuildIDLT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDLTE applies the LTE predicate on the "discord_guild_id" field.
func DiscordGuildIDLTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDContains applies the Contains predicate on the "discord_guild_id" field.
func DiscordGuildIDContains(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDHasPrefix applies the HasPrefix predicate on the "discord_guild_id" field.
func DiscordGuildIDHasPrefix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDHasSuffix applies the HasSuffix predicate on the "discord_guild_id" field.
func DiscordGuildIDHasSuffix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDEqualFold applies the EqualFold predicate on the "discord_guild_id" field.
func DiscordGuildIDEqualFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDiscordGuildID), v))
	})
}

// DiscordGuildIDContainsFold applies the ContainsFold predicate on the "discord_guild_id" field.
func DiscordGuildIDContainsFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDiscordGuildID), v))
	})
}

// HasConfigs applies the HasEdge predicate on the "configs" edge.
func HasConfigs() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConfigsTable, ConfigsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConfigsWith applies the HasEdge predicate on the "configs" edge with a given conditions (other predicates).
func HasConfigsWith(preds ...predicate.InstanceConfig) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConfigsTable, ConfigsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Instance) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Instance) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Instance) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		p(s.Not())
	})
}
