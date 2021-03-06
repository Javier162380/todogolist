// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"
	"todogolist/ent/predicate"
	"todogolist/ent/task"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTask = "Task"
)

// TaskMutation represents an operation that mutate the Tasks
// nodes in the graph.
type TaskMutation struct {
	config
	op                   Op
	typ                  string
	id                   *int
	_TaskName            *string
	_TaskLabel           *string
	_TaskDate            *time.Time
	_RecurrentTask       *bool
	_TaskTimeInvested    *int
	add_TaskTimeInvested *int
	_Periodicity         *string
	clearedFields        map[string]struct{}
	done                 bool
	oldValue             func(context.Context) (*Task, error)
	predicates           []predicate.Task
}

var _ ent.Mutation = (*TaskMutation)(nil)

// taskOption allows to manage the mutation configuration using functional options.
type taskOption func(*TaskMutation)

// newTaskMutation creates new mutation for $n.Name.
func newTaskMutation(c config, op Op, opts ...taskOption) *TaskMutation {
	m := &TaskMutation{
		config:        c,
		op:            op,
		typ:           TypeTask,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTaskID sets the id field of the mutation.
func withTaskID(id int) taskOption {
	return func(m *TaskMutation) {
		var (
			err   error
			once  sync.Once
			value *Task
		)
		m.oldValue = func(ctx context.Context) (*Task, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Task.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTask sets the old Task of the mutation.
func withTask(node *Task) taskOption {
	return func(m *TaskMutation) {
		m.oldValue = func(context.Context) (*Task, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TaskMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TaskMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *TaskMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTaskName sets the TaskName field.
func (m *TaskMutation) SetTaskName(s string) {
	m._TaskName = &s
}

// TaskName returns the TaskName value in the mutation.
func (m *TaskMutation) TaskName() (r string, exists bool) {
	v := m._TaskName
	if v == nil {
		return
	}
	return *v, true
}

// OldTaskName returns the old TaskName value of the Task.
// If the Task object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TaskMutation) OldTaskName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTaskName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTaskName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTaskName: %w", err)
	}
	return oldValue.TaskName, nil
}

// ResetTaskName reset all changes of the "TaskName" field.
func (m *TaskMutation) ResetTaskName() {
	m._TaskName = nil
}

// SetTaskLabel sets the TaskLabel field.
func (m *TaskMutation) SetTaskLabel(s string) {
	m._TaskLabel = &s
}

// TaskLabel returns the TaskLabel value in the mutation.
func (m *TaskMutation) TaskLabel() (r string, exists bool) {
	v := m._TaskLabel
	if v == nil {
		return
	}
	return *v, true
}

// OldTaskLabel returns the old TaskLabel value of the Task.
// If the Task object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TaskMutation) OldTaskLabel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTaskLabel is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTaskLabel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTaskLabel: %w", err)
	}
	return oldValue.TaskLabel, nil
}

// ResetTaskLabel reset all changes of the "TaskLabel" field.
func (m *TaskMutation) ResetTaskLabel() {
	m._TaskLabel = nil
}

// SetTaskDate sets the TaskDate field.
func (m *TaskMutation) SetTaskDate(t time.Time) {
	m._TaskDate = &t
}

// TaskDate returns the TaskDate value in the mutation.
func (m *TaskMutation) TaskDate() (r time.Time, exists bool) {
	v := m._TaskDate
	if v == nil {
		return
	}
	return *v, true
}

// OldTaskDate returns the old TaskDate value of the Task.
// If the Task object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TaskMutation) OldTaskDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTaskDate is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTaskDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTaskDate: %w", err)
	}
	return oldValue.TaskDate, nil
}

// ResetTaskDate reset all changes of the "TaskDate" field.
func (m *TaskMutation) ResetTaskDate() {
	m._TaskDate = nil
}

// SetRecurrentTask sets the RecurrentTask field.
func (m *TaskMutation) SetRecurrentTask(b bool) {
	m._RecurrentTask = &b
}

// RecurrentTask returns the RecurrentTask value in the mutation.
func (m *TaskMutation) RecurrentTask() (r bool, exists bool) {
	v := m._RecurrentTask
	if v == nil {
		return
	}
	return *v, true
}

// OldRecurrentTask returns the old RecurrentTask value of the Task.
// If the Task object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TaskMutation) OldRecurrentTask(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRecurrentTask is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRecurrentTask requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecurrentTask: %w", err)
	}
	return oldValue.RecurrentTask, nil
}

// ResetRecurrentTask reset all changes of the "RecurrentTask" field.
func (m *TaskMutation) ResetRecurrentTask() {
	m._RecurrentTask = nil
}

// SetTaskTimeInvested sets the TaskTimeInvested field.
func (m *TaskMutation) SetTaskTimeInvested(i int) {
	m._TaskTimeInvested = &i
	m.add_TaskTimeInvested = nil
}

// TaskTimeInvested returns the TaskTimeInvested value in the mutation.
func (m *TaskMutation) TaskTimeInvested() (r int, exists bool) {
	v := m._TaskTimeInvested
	if v == nil {
		return
	}
	return *v, true
}

// OldTaskTimeInvested returns the old TaskTimeInvested value of the Task.
// If the Task object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TaskMutation) OldTaskTimeInvested(ctx context.Context) (v *int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTaskTimeInvested is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTaskTimeInvested requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTaskTimeInvested: %w", err)
	}
	return oldValue.TaskTimeInvested, nil
}

// AddTaskTimeInvested adds i to TaskTimeInvested.
func (m *TaskMutation) AddTaskTimeInvested(i int) {
	if m.add_TaskTimeInvested != nil {
		*m.add_TaskTimeInvested += i
	} else {
		m.add_TaskTimeInvested = &i
	}
}

// AddedTaskTimeInvested returns the value that was added to the TaskTimeInvested field in this mutation.
func (m *TaskMutation) AddedTaskTimeInvested() (r int, exists bool) {
	v := m.add_TaskTimeInvested
	if v == nil {
		return
	}
	return *v, true
}

// ClearTaskTimeInvested clears the value of TaskTimeInvested.
func (m *TaskMutation) ClearTaskTimeInvested() {
	m._TaskTimeInvested = nil
	m.add_TaskTimeInvested = nil
	m.clearedFields[task.FieldTaskTimeInvested] = struct{}{}
}

// TaskTimeInvestedCleared returns if the field TaskTimeInvested was cleared in this mutation.
func (m *TaskMutation) TaskTimeInvestedCleared() bool {
	_, ok := m.clearedFields[task.FieldTaskTimeInvested]
	return ok
}

// ResetTaskTimeInvested reset all changes of the "TaskTimeInvested" field.
func (m *TaskMutation) ResetTaskTimeInvested() {
	m._TaskTimeInvested = nil
	m.add_TaskTimeInvested = nil
	delete(m.clearedFields, task.FieldTaskTimeInvested)
}

// SetPeriodicity sets the Periodicity field.
func (m *TaskMutation) SetPeriodicity(s string) {
	m._Periodicity = &s
}

// Periodicity returns the Periodicity value in the mutation.
func (m *TaskMutation) Periodicity() (r string, exists bool) {
	v := m._Periodicity
	if v == nil {
		return
	}
	return *v, true
}

// OldPeriodicity returns the old Periodicity value of the Task.
// If the Task object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *TaskMutation) OldPeriodicity(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPeriodicity is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPeriodicity requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPeriodicity: %w", err)
	}
	return oldValue.Periodicity, nil
}

// ClearPeriodicity clears the value of Periodicity.
func (m *TaskMutation) ClearPeriodicity() {
	m._Periodicity = nil
	m.clearedFields[task.FieldPeriodicity] = struct{}{}
}

// PeriodicityCleared returns if the field Periodicity was cleared in this mutation.
func (m *TaskMutation) PeriodicityCleared() bool {
	_, ok := m.clearedFields[task.FieldPeriodicity]
	return ok
}

// ResetPeriodicity reset all changes of the "Periodicity" field.
func (m *TaskMutation) ResetPeriodicity() {
	m._Periodicity = nil
	delete(m.clearedFields, task.FieldPeriodicity)
}

// Op returns the operation name.
func (m *TaskMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Task).
func (m *TaskMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *TaskMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m._TaskName != nil {
		fields = append(fields, task.FieldTaskName)
	}
	if m._TaskLabel != nil {
		fields = append(fields, task.FieldTaskLabel)
	}
	if m._TaskDate != nil {
		fields = append(fields, task.FieldTaskDate)
	}
	if m._RecurrentTask != nil {
		fields = append(fields, task.FieldRecurrentTask)
	}
	if m._TaskTimeInvested != nil {
		fields = append(fields, task.FieldTaskTimeInvested)
	}
	if m._Periodicity != nil {
		fields = append(fields, task.FieldPeriodicity)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *TaskMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case task.FieldTaskName:
		return m.TaskName()
	case task.FieldTaskLabel:
		return m.TaskLabel()
	case task.FieldTaskDate:
		return m.TaskDate()
	case task.FieldRecurrentTask:
		return m.RecurrentTask()
	case task.FieldTaskTimeInvested:
		return m.TaskTimeInvested()
	case task.FieldPeriodicity:
		return m.Periodicity()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *TaskMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case task.FieldTaskName:
		return m.OldTaskName(ctx)
	case task.FieldTaskLabel:
		return m.OldTaskLabel(ctx)
	case task.FieldTaskDate:
		return m.OldTaskDate(ctx)
	case task.FieldRecurrentTask:
		return m.OldRecurrentTask(ctx)
	case task.FieldTaskTimeInvested:
		return m.OldTaskTimeInvested(ctx)
	case task.FieldPeriodicity:
		return m.OldPeriodicity(ctx)
	}
	return nil, fmt.Errorf("unknown Task field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TaskMutation) SetField(name string, value ent.Value) error {
	switch name {
	case task.FieldTaskName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTaskName(v)
		return nil
	case task.FieldTaskLabel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTaskLabel(v)
		return nil
	case task.FieldTaskDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTaskDate(v)
		return nil
	case task.FieldRecurrentTask:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecurrentTask(v)
		return nil
	case task.FieldTaskTimeInvested:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTaskTimeInvested(v)
		return nil
	case task.FieldPeriodicity:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPeriodicity(v)
		return nil
	}
	return fmt.Errorf("unknown Task field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *TaskMutation) AddedFields() []string {
	var fields []string
	if m.add_TaskTimeInvested != nil {
		fields = append(fields, task.FieldTaskTimeInvested)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *TaskMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case task.FieldTaskTimeInvested:
		return m.AddedTaskTimeInvested()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TaskMutation) AddField(name string, value ent.Value) error {
	switch name {
	case task.FieldTaskTimeInvested:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTaskTimeInvested(v)
		return nil
	}
	return fmt.Errorf("unknown Task numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *TaskMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(task.FieldTaskTimeInvested) {
		fields = append(fields, task.FieldTaskTimeInvested)
	}
	if m.FieldCleared(task.FieldPeriodicity) {
		fields = append(fields, task.FieldPeriodicity)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *TaskMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *TaskMutation) ClearField(name string) error {
	switch name {
	case task.FieldTaskTimeInvested:
		m.ClearTaskTimeInvested()
		return nil
	case task.FieldPeriodicity:
		m.ClearPeriodicity()
		return nil
	}
	return fmt.Errorf("unknown Task nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *TaskMutation) ResetField(name string) error {
	switch name {
	case task.FieldTaskName:
		m.ResetTaskName()
		return nil
	case task.FieldTaskLabel:
		m.ResetTaskLabel()
		return nil
	case task.FieldTaskDate:
		m.ResetTaskDate()
		return nil
	case task.FieldRecurrentTask:
		m.ResetRecurrentTask()
		return nil
	case task.FieldTaskTimeInvested:
		m.ResetTaskTimeInvested()
		return nil
	case task.FieldPeriodicity:
		m.ResetPeriodicity()
		return nil
	}
	return fmt.Errorf("unknown Task field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *TaskMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *TaskMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *TaskMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *TaskMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *TaskMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *TaskMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *TaskMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Task unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *TaskMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Task edge %s", name)
}
