// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"todogolist/ent/predicate"
	"todogolist/ent/task"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where adds a new predicate for the builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.predicates = append(tu.mutation.predicates, ps...)
	return tu
}

// SetTaskName sets the TaskName field.
func (tu *TaskUpdate) SetTaskName(s string) *TaskUpdate {
	tu.mutation.SetTaskName(s)
	return tu
}

// SetTaskLabel sets the TaskLabel field.
func (tu *TaskUpdate) SetTaskLabel(s string) *TaskUpdate {
	tu.mutation.SetTaskLabel(s)
	return tu
}

// SetNillableTaskLabel sets the TaskLabel field if the given value is not nil.
func (tu *TaskUpdate) SetNillableTaskLabel(s *string) *TaskUpdate {
	if s != nil {
		tu.SetTaskLabel(*s)
	}
	return tu
}

// SetTaskDate sets the TaskDate field.
func (tu *TaskUpdate) SetTaskDate(t time.Time) *TaskUpdate {
	tu.mutation.SetTaskDate(t)
	return tu
}

// SetNillableTaskDate sets the TaskDate field if the given value is not nil.
func (tu *TaskUpdate) SetNillableTaskDate(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetTaskDate(*t)
	}
	return tu
}

// SetRecurrentTask sets the RecurrentTask field.
func (tu *TaskUpdate) SetRecurrentTask(b bool) *TaskUpdate {
	tu.mutation.SetRecurrentTask(b)
	return tu
}

// SetNillableRecurrentTask sets the RecurrentTask field if the given value is not nil.
func (tu *TaskUpdate) SetNillableRecurrentTask(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetRecurrentTask(*b)
	}
	return tu
}

// SetTaskTimeInvested sets the TaskTimeInvested field.
func (tu *TaskUpdate) SetTaskTimeInvested(i int) *TaskUpdate {
	tu.mutation.ResetTaskTimeInvested()
	tu.mutation.SetTaskTimeInvested(i)
	return tu
}

// SetNillableTaskTimeInvested sets the TaskTimeInvested field if the given value is not nil.
func (tu *TaskUpdate) SetNillableTaskTimeInvested(i *int) *TaskUpdate {
	if i != nil {
		tu.SetTaskTimeInvested(*i)
	}
	return tu
}

// AddTaskTimeInvested adds i to TaskTimeInvested.
func (tu *TaskUpdate) AddTaskTimeInvested(i int) *TaskUpdate {
	tu.mutation.AddTaskTimeInvested(i)
	return tu
}

// ClearTaskTimeInvested clears the value of TaskTimeInvested.
func (tu *TaskUpdate) ClearTaskTimeInvested() *TaskUpdate {
	tu.mutation.ClearTaskTimeInvested()
	return tu
}

// SetPeriodicity sets the Periodicity field.
func (tu *TaskUpdate) SetPeriodicity(s string) *TaskUpdate {
	tu.mutation.SetPeriodicity(s)
	return tu
}

// SetNillablePeriodicity sets the Periodicity field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePeriodicity(s *string) *TaskUpdate {
	if s != nil {
		tu.SetPeriodicity(*s)
	}
	return tu
}

// ClearPeriodicity clears the value of Periodicity.
func (tu *TaskUpdate) ClearPeriodicity() *TaskUpdate {
	tu.mutation.ClearPeriodicity()
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.TaskName(); ok {
		if err := task.TaskNameValidator(v); err != nil {
			return &ValidationError{Name: "TaskName", err: fmt.Errorf("ent: validator failed for field \"TaskName\": %w", err)}
		}
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.TaskName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTaskName,
		})
	}
	if value, ok := tu.mutation.TaskLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTaskLabel,
		})
	}
	if value, ok := tu.mutation.TaskDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldTaskDate,
		})
	}
	if value, ok := tu.mutation.RecurrentTask(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldRecurrentTask,
		})
	}
	if value, ok := tu.mutation.TaskTimeInvested(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: task.FieldTaskTimeInvested,
		})
	}
	if value, ok := tu.mutation.AddedTaskTimeInvested(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: task.FieldTaskTimeInvested,
		})
	}
	if tu.mutation.TaskTimeInvestedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: task.FieldTaskTimeInvested,
		})
	}
	if value, ok := tu.mutation.Periodicity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldPeriodicity,
		})
	}
	if tu.mutation.PeriodicityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldPeriodicity,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// SetTaskName sets the TaskName field.
func (tuo *TaskUpdateOne) SetTaskName(s string) *TaskUpdateOne {
	tuo.mutation.SetTaskName(s)
	return tuo
}

// SetTaskLabel sets the TaskLabel field.
func (tuo *TaskUpdateOne) SetTaskLabel(s string) *TaskUpdateOne {
	tuo.mutation.SetTaskLabel(s)
	return tuo
}

// SetNillableTaskLabel sets the TaskLabel field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTaskLabel(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetTaskLabel(*s)
	}
	return tuo
}

// SetTaskDate sets the TaskDate field.
func (tuo *TaskUpdateOne) SetTaskDate(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetTaskDate(t)
	return tuo
}

// SetNillableTaskDate sets the TaskDate field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTaskDate(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetTaskDate(*t)
	}
	return tuo
}

// SetRecurrentTask sets the RecurrentTask field.
func (tuo *TaskUpdateOne) SetRecurrentTask(b bool) *TaskUpdateOne {
	tuo.mutation.SetRecurrentTask(b)
	return tuo
}

// SetNillableRecurrentTask sets the RecurrentTask field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableRecurrentTask(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetRecurrentTask(*b)
	}
	return tuo
}

// SetTaskTimeInvested sets the TaskTimeInvested field.
func (tuo *TaskUpdateOne) SetTaskTimeInvested(i int) *TaskUpdateOne {
	tuo.mutation.ResetTaskTimeInvested()
	tuo.mutation.SetTaskTimeInvested(i)
	return tuo
}

// SetNillableTaskTimeInvested sets the TaskTimeInvested field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTaskTimeInvested(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetTaskTimeInvested(*i)
	}
	return tuo
}

// AddTaskTimeInvested adds i to TaskTimeInvested.
func (tuo *TaskUpdateOne) AddTaskTimeInvested(i int) *TaskUpdateOne {
	tuo.mutation.AddTaskTimeInvested(i)
	return tuo
}

// ClearTaskTimeInvested clears the value of TaskTimeInvested.
func (tuo *TaskUpdateOne) ClearTaskTimeInvested() *TaskUpdateOne {
	tuo.mutation.ClearTaskTimeInvested()
	return tuo
}

// SetPeriodicity sets the Periodicity field.
func (tuo *TaskUpdateOne) SetPeriodicity(s string) *TaskUpdateOne {
	tuo.mutation.SetPeriodicity(s)
	return tuo
}

// SetNillablePeriodicity sets the Periodicity field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePeriodicity(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetPeriodicity(*s)
	}
	return tuo
}

// ClearPeriodicity clears the value of Periodicity.
func (tuo *TaskUpdateOne) ClearPeriodicity() *TaskUpdateOne {
	tuo.mutation.ClearPeriodicity()
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Save executes the query and returns the updated entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.TaskName(); ok {
		if err := task.TaskNameValidator(v); err != nil {
			return &ValidationError{Name: "TaskName", err: fmt.Errorf("ent: validator failed for field \"TaskName\": %w", err)}
		}
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Task.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.TaskName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTaskName,
		})
	}
	if value, ok := tuo.mutation.TaskLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTaskLabel,
		})
	}
	if value, ok := tuo.mutation.TaskDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldTaskDate,
		})
	}
	if value, ok := tuo.mutation.RecurrentTask(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldRecurrentTask,
		})
	}
	if value, ok := tuo.mutation.TaskTimeInvested(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: task.FieldTaskTimeInvested,
		})
	}
	if value, ok := tuo.mutation.AddedTaskTimeInvested(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: task.FieldTaskTimeInvested,
		})
	}
	if tuo.mutation.TaskTimeInvestedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: task.FieldTaskTimeInvested,
		})
	}
	if value, ok := tuo.mutation.Periodicity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldPeriodicity,
		})
	}
	if tuo.mutation.PeriodicityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldPeriodicity,
		})
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
