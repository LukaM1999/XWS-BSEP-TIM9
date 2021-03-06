// Code generated by ent, DO NOT EDIT.

package connection

import (
	"dislinkt/connection_service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// IsApproved applies equality check predicate on the "is_approved" field. It's identical to IsApprovedEQ.
func IsApproved(v bool) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsApproved), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// ConnectionID applies equality check predicate on the "connection_id" field. It's identical to ConnectionIDEQ.
func ConnectionID(v int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConnectionID), v))
	})
}

// IssuerPrimaryKey applies equality check predicate on the "issuer_primary_key" field. It's identical to IssuerPrimaryKeyEQ.
func IssuerPrimaryKey(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIssuerPrimaryKey), v))
	})
}

// SubjectPrimaryKey applies equality check predicate on the "subject_primary_key" field. It's identical to SubjectPrimaryKeyEQ.
func SubjectPrimaryKey(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubjectPrimaryKey), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// IsApprovedEQ applies the EQ predicate on the "is_approved" field.
func IsApprovedEQ(v bool) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsApproved), v))
	})
}

// IsApprovedNEQ applies the NEQ predicate on the "is_approved" field.
func IsApprovedNEQ(v bool) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsApproved), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// ConnectionIDEQ applies the EQ predicate on the "connection_id" field.
func ConnectionIDEQ(v int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConnectionID), v))
	})
}

// ConnectionIDNEQ applies the NEQ predicate on the "connection_id" field.
func ConnectionIDNEQ(v int) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConnectionID), v))
	})
}

// ConnectionIDIn applies the In predicate on the "connection_id" field.
func ConnectionIDIn(vs ...int) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConnectionID), v...))
	})
}

// ConnectionIDNotIn applies the NotIn predicate on the "connection_id" field.
func ConnectionIDNotIn(vs ...int) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConnectionID), v...))
	})
}

// IssuerPrimaryKeyEQ applies the EQ predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyEQ(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyNEQ applies the NEQ predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyNEQ(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyIn applies the In predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyIn(vs ...string) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIssuerPrimaryKey), v...))
	})
}

// IssuerPrimaryKeyNotIn applies the NotIn predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyNotIn(vs ...string) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIssuerPrimaryKey), v...))
	})
}

// IssuerPrimaryKeyGT applies the GT predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyGT(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyGTE applies the GTE predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyGTE(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyLT applies the LT predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyLT(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyLTE applies the LTE predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyLTE(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyContains applies the Contains predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyContains(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyHasPrefix applies the HasPrefix predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyHasPrefix(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyHasSuffix applies the HasSuffix predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyHasSuffix(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyEqualFold applies the EqualFold predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyEqualFold(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIssuerPrimaryKey), v))
	})
}

// IssuerPrimaryKeyContainsFold applies the ContainsFold predicate on the "issuer_primary_key" field.
func IssuerPrimaryKeyContainsFold(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIssuerPrimaryKey), v))
	})
}

// SubjectPrimaryKeyEQ applies the EQ predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyEQ(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyNEQ applies the NEQ predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyNEQ(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyIn applies the In predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyIn(vs ...string) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSubjectPrimaryKey), v...))
	})
}

// SubjectPrimaryKeyNotIn applies the NotIn predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyNotIn(vs ...string) predicate.Connection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connection(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSubjectPrimaryKey), v...))
	})
}

// SubjectPrimaryKeyGT applies the GT predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyGT(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyGTE applies the GTE predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyGTE(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyLT applies the LT predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyLT(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyLTE applies the LTE predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyLTE(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyContains applies the Contains predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyContains(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyHasPrefix applies the HasPrefix predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyHasPrefix(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyHasSuffix applies the HasSuffix predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyHasSuffix(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyEqualFold applies the EqualFold predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyEqualFold(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSubjectPrimaryKey), v))
	})
}

// SubjectPrimaryKeyContainsFold applies the ContainsFold predicate on the "subject_primary_key" field.
func SubjectPrimaryKeyContainsFold(v string) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSubjectPrimaryKey), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasConnection applies the HasEdge predicate on the "connection" edge.
func HasConnection() predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ConnectionTable, ConnectionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConnectionWith applies the HasEdge predicate on the "connection" edge with a given conditions (other predicates).
func HasConnectionWith(preds ...predicate.User) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ConnectionTable, ConnectionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
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
func Not(p predicate.Connection) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		p(s.Not())
	})
}
