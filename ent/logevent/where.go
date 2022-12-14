// Code generated by ent, DO NOT EDIT.

package logevent

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/velcro-xiv/velcro/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.LogEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.LogEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLevel), v))
	})
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.LogEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLevel), v...))
	})
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.LogEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLevel), v...))
	})
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLevel), v))
	})
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLevel), v))
	})
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLevel), v))
	})
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLevel), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.LogEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.LogEvent {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LogEvent) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LogEvent) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
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
func Not(p predicate.LogEvent) predicate.LogEvent {
	return predicate.LogEvent(func(s *sql.Selector) {
		p(s.Not())
	})
}
