// Code generated by entc, DO NOT EDIT.

package task

import (
	"time"
	"todogolist/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TaskName applies equality check predicate on the "TaskName" field. It's identical to TaskNameEQ.
func TaskName(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskName), v))
	})
}

// TaskLabel applies equality check predicate on the "TaskLabel" field. It's identical to TaskLabelEQ.
func TaskLabel(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskLabel), v))
	})
}

// TaskDate applies equality check predicate on the "TaskDate" field. It's identical to TaskDateEQ.
func TaskDate(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskDate), v))
	})
}

// RecurrentTask applies equality check predicate on the "RecurrentTask" field. It's identical to RecurrentTaskEQ.
func RecurrentTask(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecurrentTask), v))
	})
}

// TaskTimeInvested applies equality check predicate on the "TaskTimeInvested" field. It's identical to TaskTimeInvestedEQ.
func TaskTimeInvested(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskTimeInvested), v))
	})
}

// Periodicity applies equality check predicate on the "Periodicity" field. It's identical to PeriodicityEQ.
func Periodicity(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPeriodicity), v))
	})
}

// TaskNameEQ applies the EQ predicate on the "TaskName" field.
func TaskNameEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskName), v))
	})
}

// TaskNameNEQ applies the NEQ predicate on the "TaskName" field.
func TaskNameNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskName), v))
	})
}

// TaskNameIn applies the In predicate on the "TaskName" field.
func TaskNameIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskName), v...))
	})
}

// TaskNameNotIn applies the NotIn predicate on the "TaskName" field.
func TaskNameNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskName), v...))
	})
}

// TaskNameGT applies the GT predicate on the "TaskName" field.
func TaskNameGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskName), v))
	})
}

// TaskNameGTE applies the GTE predicate on the "TaskName" field.
func TaskNameGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskName), v))
	})
}

// TaskNameLT applies the LT predicate on the "TaskName" field.
func TaskNameLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskName), v))
	})
}

// TaskNameLTE applies the LTE predicate on the "TaskName" field.
func TaskNameLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskName), v))
	})
}

// TaskNameContains applies the Contains predicate on the "TaskName" field.
func TaskNameContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskName), v))
	})
}

// TaskNameHasPrefix applies the HasPrefix predicate on the "TaskName" field.
func TaskNameHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskName), v))
	})
}

// TaskNameHasSuffix applies the HasSuffix predicate on the "TaskName" field.
func TaskNameHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskName), v))
	})
}

// TaskNameEqualFold applies the EqualFold predicate on the "TaskName" field.
func TaskNameEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskName), v))
	})
}

// TaskNameContainsFold applies the ContainsFold predicate on the "TaskName" field.
func TaskNameContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskName), v))
	})
}

// TaskLabelEQ applies the EQ predicate on the "TaskLabel" field.
func TaskLabelEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelNEQ applies the NEQ predicate on the "TaskLabel" field.
func TaskLabelNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelIn applies the In predicate on the "TaskLabel" field.
func TaskLabelIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskLabel), v...))
	})
}

// TaskLabelNotIn applies the NotIn predicate on the "TaskLabel" field.
func TaskLabelNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskLabel), v...))
	})
}

// TaskLabelGT applies the GT predicate on the "TaskLabel" field.
func TaskLabelGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelGTE applies the GTE predicate on the "TaskLabel" field.
func TaskLabelGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelLT applies the LT predicate on the "TaskLabel" field.
func TaskLabelLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelLTE applies the LTE predicate on the "TaskLabel" field.
func TaskLabelLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelContains applies the Contains predicate on the "TaskLabel" field.
func TaskLabelContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelHasPrefix applies the HasPrefix predicate on the "TaskLabel" field.
func TaskLabelHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelHasSuffix applies the HasSuffix predicate on the "TaskLabel" field.
func TaskLabelHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelEqualFold applies the EqualFold predicate on the "TaskLabel" field.
func TaskLabelEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTaskLabel), v))
	})
}

// TaskLabelContainsFold applies the ContainsFold predicate on the "TaskLabel" field.
func TaskLabelContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTaskLabel), v))
	})
}

// TaskDateEQ applies the EQ predicate on the "TaskDate" field.
func TaskDateEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskDate), v))
	})
}

// TaskDateNEQ applies the NEQ predicate on the "TaskDate" field.
func TaskDateNEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskDate), v))
	})
}

// TaskDateIn applies the In predicate on the "TaskDate" field.
func TaskDateIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskDate), v...))
	})
}

// TaskDateNotIn applies the NotIn predicate on the "TaskDate" field.
func TaskDateNotIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskDate), v...))
	})
}

// TaskDateGT applies the GT predicate on the "TaskDate" field.
func TaskDateGT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskDate), v))
	})
}

// TaskDateGTE applies the GTE predicate on the "TaskDate" field.
func TaskDateGTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskDate), v))
	})
}

// TaskDateLT applies the LT predicate on the "TaskDate" field.
func TaskDateLT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskDate), v))
	})
}

// TaskDateLTE applies the LTE predicate on the "TaskDate" field.
func TaskDateLTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskDate), v))
	})
}

// RecurrentTaskEQ applies the EQ predicate on the "RecurrentTask" field.
func RecurrentTaskEQ(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecurrentTask), v))
	})
}

// RecurrentTaskNEQ applies the NEQ predicate on the "RecurrentTask" field.
func RecurrentTaskNEQ(v bool) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecurrentTask), v))
	})
}

// TaskTimeInvestedEQ applies the EQ predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedEQ(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTaskTimeInvested), v))
	})
}

// TaskTimeInvestedNEQ applies the NEQ predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedNEQ(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTaskTimeInvested), v))
	})
}

// TaskTimeInvestedIn applies the In predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedIn(vs ...int) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTaskTimeInvested), v...))
	})
}

// TaskTimeInvestedNotIn applies the NotIn predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedNotIn(vs ...int) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTaskTimeInvested), v...))
	})
}

// TaskTimeInvestedGT applies the GT predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedGT(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTaskTimeInvested), v))
	})
}

// TaskTimeInvestedGTE applies the GTE predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedGTE(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTaskTimeInvested), v))
	})
}

// TaskTimeInvestedLT applies the LT predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedLT(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTaskTimeInvested), v))
	})
}

// TaskTimeInvestedLTE applies the LTE predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedLTE(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTaskTimeInvested), v))
	})
}

// TaskTimeInvestedIsNil applies the IsNil predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedIsNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTaskTimeInvested)))
	})
}

// TaskTimeInvestedNotNil applies the NotNil predicate on the "TaskTimeInvested" field.
func TaskTimeInvestedNotNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTaskTimeInvested)))
	})
}

// PeriodicityEQ applies the EQ predicate on the "Periodicity" field.
func PeriodicityEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityNEQ applies the NEQ predicate on the "Periodicity" field.
func PeriodicityNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityIn applies the In predicate on the "Periodicity" field.
func PeriodicityIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPeriodicity), v...))
	})
}

// PeriodicityNotIn applies the NotIn predicate on the "Periodicity" field.
func PeriodicityNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPeriodicity), v...))
	})
}

// PeriodicityGT applies the GT predicate on the "Periodicity" field.
func PeriodicityGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityGTE applies the GTE predicate on the "Periodicity" field.
func PeriodicityGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityLT applies the LT predicate on the "Periodicity" field.
func PeriodicityLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityLTE applies the LTE predicate on the "Periodicity" field.
func PeriodicityLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityContains applies the Contains predicate on the "Periodicity" field.
func PeriodicityContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityHasPrefix applies the HasPrefix predicate on the "Periodicity" field.
func PeriodicityHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityHasSuffix applies the HasSuffix predicate on the "Periodicity" field.
func PeriodicityHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityIsNil applies the IsNil predicate on the "Periodicity" field.
func PeriodicityIsNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPeriodicity)))
	})
}

// PeriodicityNotNil applies the NotNil predicate on the "Periodicity" field.
func PeriodicityNotNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPeriodicity)))
	})
}

// PeriodicityEqualFold applies the EqualFold predicate on the "Periodicity" field.
func PeriodicityEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPeriodicity), v))
	})
}

// PeriodicityContainsFold applies the ContainsFold predicate on the "Periodicity" field.
func PeriodicityContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPeriodicity), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}
