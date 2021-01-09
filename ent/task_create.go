// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todogolist/ent/task"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetTaskName sets the TaskName field.
func (tc *TaskCreate) SetTaskName(s string) *TaskCreate {
	tc.mutation.SetTaskName(s)
	return tc
}

// SetTaskLabel sets the TaskLabel field.
func (tc *TaskCreate) SetTaskLabel(s string) *TaskCreate {
	tc.mutation.SetTaskLabel(s)
	return tc
}

// SetNillableTaskLabel sets the TaskLabel field if the given value is not nil.
func (tc *TaskCreate) SetNillableTaskLabel(s *string) *TaskCreate {
	if s != nil {
		tc.SetTaskLabel(*s)
	}
	return tc
}

// SetTaskDate sets the TaskDate field.
func (tc *TaskCreate) SetTaskDate(t time.Time) *TaskCreate {
	tc.mutation.SetTaskDate(t)
	return tc
}

// SetNillableTaskDate sets the TaskDate field if the given value is not nil.
func (tc *TaskCreate) SetNillableTaskDate(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetTaskDate(*t)
	}
	return tc
}

// SetRecurrentTask sets the RecurrentTask field.
func (tc *TaskCreate) SetRecurrentTask(b bool) *TaskCreate {
	tc.mutation.SetRecurrentTask(b)
	return tc
}

// SetNillableRecurrentTask sets the RecurrentTask field if the given value is not nil.
func (tc *TaskCreate) SetNillableRecurrentTask(b *bool) *TaskCreate {
	if b != nil {
		tc.SetRecurrentTask(*b)
	}
	return tc
}

// SetTaskTimeInvested sets the TaskTimeInvested field.
func (tc *TaskCreate) SetTaskTimeInvested(i int) *TaskCreate {
	tc.mutation.SetTaskTimeInvested(i)
	return tc
}

// SetNillableTaskTimeInvested sets the TaskTimeInvested field if the given value is not nil.
func (tc *TaskCreate) SetNillableTaskTimeInvested(i *int) *TaskCreate {
	if i != nil {
		tc.SetTaskTimeInvested(*i)
	}
	return tc
}

// SetPeriodicity sets the Periodicity field.
func (tc *TaskCreate) SetPeriodicity(s string) *TaskCreate {
	tc.mutation.SetPeriodicity(s)
	return tc
}

// SetNillablePeriodicity sets the Periodicity field if the given value is not nil.
func (tc *TaskCreate) SetNillablePeriodicity(s *string) *TaskCreate {
	if s != nil {
		tc.SetPeriodicity(*s)
	}
	return tc
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.TaskLabel(); !ok {
		v := task.DefaultTaskLabel
		tc.mutation.SetTaskLabel(v)
	}
	if _, ok := tc.mutation.TaskDate(); !ok {
		v := task.DefaultTaskDate()
		tc.mutation.SetTaskDate(v)
	}
	if _, ok := tc.mutation.RecurrentTask(); !ok {
		v := task.DefaultRecurrentTask
		tc.mutation.SetRecurrentTask(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.TaskName(); !ok {
		return &ValidationError{Name: "TaskName", err: errors.New("ent: missing required field \"TaskName\"")}
	}
	if v, ok := tc.mutation.TaskName(); ok {
		if err := task.TaskNameValidator(v); err != nil {
			return &ValidationError{Name: "TaskName", err: fmt.Errorf("ent: validator failed for field \"TaskName\": %w", err)}
		}
	}
	if _, ok := tc.mutation.TaskLabel(); !ok {
		return &ValidationError{Name: "TaskLabel", err: errors.New("ent: missing required field \"TaskLabel\"")}
	}
	if _, ok := tc.mutation.TaskDate(); !ok {
		return &ValidationError{Name: "TaskDate", err: errors.New("ent: missing required field \"TaskDate\"")}
	}
	if _, ok := tc.mutation.RecurrentTask(); !ok {
		return &ValidationError{Name: "RecurrentTask", err: errors.New("ent: missing required field \"RecurrentTask\"")}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.TaskName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTaskName,
		})
		_node.TaskName = value
	}
	if value, ok := tc.mutation.TaskLabel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTaskLabel,
		})
		_node.TaskLabel = value
	}
	if value, ok := tc.mutation.TaskDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldTaskDate,
		})
		_node.TaskDate = value
	}
	if value, ok := tc.mutation.RecurrentTask(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldRecurrentTask,
		})
		_node.RecurrentTask = value
	}
	if value, ok := tc.mutation.TaskTimeInvested(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: task.FieldTaskTimeInvested,
		})
		_node.TaskTimeInvested = &value
	}
	if value, ok := tc.mutation.Periodicity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldPeriodicity,
		})
		_node.Periodicity = &value
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating a bulk of Task entities.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
