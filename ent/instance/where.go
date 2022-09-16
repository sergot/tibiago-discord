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

// SessionID applies equality check predicate on the "session_id" field. It's identical to SessionIDEQ.
func SessionID(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionID), v))
	})
}

// SessionIDEQ applies the EQ predicate on the "session_id" field.
func SessionIDEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionID), v))
	})
}

// SessionIDNEQ applies the NEQ predicate on the "session_id" field.
func SessionIDNEQ(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessionID), v))
	})
}

// SessionIDIn applies the In predicate on the "session_id" field.
func SessionIDIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSessionID), v...))
	})
}

// SessionIDNotIn applies the NotIn predicate on the "session_id" field.
func SessionIDNotIn(vs ...string) predicate.Instance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSessionID), v...))
	})
}

// SessionIDGT applies the GT predicate on the "session_id" field.
func SessionIDGT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessionID), v))
	})
}

// SessionIDGTE applies the GTE predicate on the "session_id" field.
func SessionIDGTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessionID), v))
	})
}

// SessionIDLT applies the LT predicate on the "session_id" field.
func SessionIDLT(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessionID), v))
	})
}

// SessionIDLTE applies the LTE predicate on the "session_id" field.
func SessionIDLTE(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessionID), v))
	})
}

// SessionIDContains applies the Contains predicate on the "session_id" field.
func SessionIDContains(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSessionID), v))
	})
}

// SessionIDHasPrefix applies the HasPrefix predicate on the "session_id" field.
func SessionIDHasPrefix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSessionID), v))
	})
}

// SessionIDHasSuffix applies the HasSuffix predicate on the "session_id" field.
func SessionIDHasSuffix(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSessionID), v))
	})
}

// SessionIDEqualFold applies the EqualFold predicate on the "session_id" field.
func SessionIDEqualFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSessionID), v))
	})
}

// SessionIDContainsFold applies the ContainsFold predicate on the "session_id" field.
func SessionIDContainsFold(v string) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSessionID), v))
	})
}

// HasConfig applies the HasEdge predicate on the "config" edge.
func HasConfig() predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConfigTable, ConfigColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConfigWith applies the HasEdge predicate on the "config" edge with a given conditions (other predicates).
func HasConfigWith(preds ...predicate.InstanceConfig) predicate.Instance {
	return predicate.Instance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConfigTable, ConfigColumn),
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