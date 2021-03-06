// Code generated by ent, DO NOT EDIT.

package user

import (
	"dislinkt/connection_service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PrimaryKey applies equality check predicate on the "primary_key" field. It's identical to PrimaryKeyEQ.
func PrimaryKey(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryKey), v))
	})
}

// IsPrivate applies equality check predicate on the "is_private" field. It's identical to IsPrivateEQ.
func IsPrivate(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPrivate), v))
	})
}

// PrimaryKeyEQ applies the EQ predicate on the "primary_key" field.
func PrimaryKeyEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyNEQ applies the NEQ predicate on the "primary_key" field.
func PrimaryKeyNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyIn applies the In predicate on the "primary_key" field.
func PrimaryKeyIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrimaryKey), v...))
	})
}

// PrimaryKeyNotIn applies the NotIn predicate on the "primary_key" field.
func PrimaryKeyNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrimaryKey), v...))
	})
}

// PrimaryKeyGT applies the GT predicate on the "primary_key" field.
func PrimaryKeyGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyGTE applies the GTE predicate on the "primary_key" field.
func PrimaryKeyGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyLT applies the LT predicate on the "primary_key" field.
func PrimaryKeyLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyLTE applies the LTE predicate on the "primary_key" field.
func PrimaryKeyLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyContains applies the Contains predicate on the "primary_key" field.
func PrimaryKeyContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyHasPrefix applies the HasPrefix predicate on the "primary_key" field.
func PrimaryKeyHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyHasSuffix applies the HasSuffix predicate on the "primary_key" field.
func PrimaryKeyHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyEqualFold applies the EqualFold predicate on the "primary_key" field.
func PrimaryKeyEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrimaryKey), v))
	})
}

// PrimaryKeyContainsFold applies the ContainsFold predicate on the "primary_key" field.
func PrimaryKeyContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrimaryKey), v))
	})
}

// IsPrivateEQ applies the EQ predicate on the "is_private" field.
func IsPrivateEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPrivate), v))
	})
}

// IsPrivateNEQ applies the NEQ predicate on the "is_private" field.
func IsPrivateNEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsPrivate), v))
	})
}

// HasConnection applies the HasEdge predicate on the "connection" edge.
func HasConnection() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ConnectionTable, ConnectionPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConnectionWith applies the HasEdge predicate on the "connection" edge with a given conditions (other predicates).
func HasConnectionWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ConnectionTable, ConnectionPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBlock applies the HasEdge predicate on the "block" edge.
func HasBlock() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BlockTable, BlockColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlockWith applies the HasEdge predicate on the "block" edge with a given conditions (other predicates).
func HasBlockWith(preds ...predicate.BlockedUser) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BlockTable, BlockColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasConnections applies the HasEdge predicate on the "connections" edge.
func HasConnections() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ConnectionsTable, ConnectionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConnectionsWith applies the HasEdge predicate on the "connections" edge with a given conditions (other predicates).
func HasConnectionsWith(preds ...predicate.Connection) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ConnectionsTable, ConnectionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
