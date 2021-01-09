// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"todogolist/ent/task"

	"github.com/facebook/ent/dialect/sql"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TaskName holds the value of the "TaskName" field.
	TaskName string `json:"TaskName,omitempty"`
	// TaskLabel holds the value of the "TaskLabel" field.
	TaskLabel string `json:"TaskLabel,omitempty"`
	// TaskDate holds the value of the "TaskDate" field.
	TaskDate time.Time `json:"TaskDate,omitempty"`
	// RecurrentTask holds the value of the "RecurrentTask" field.
	RecurrentTask bool `json:"RecurrentTask,omitempty"`
	// TaskTimeInvested holds the value of the "TaskTimeInvested" field.
	TaskTimeInvested *int `json:"TaskTimeInvested,omitempty"`
	// Periodicity holds the value of the "Periodicity" field.
	Periodicity *string `json:"Periodicity,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // TaskName
		&sql.NullString{}, // TaskLabel
		&sql.NullTime{},   // TaskDate
		&sql.NullBool{},   // RecurrentTask
		&sql.NullInt64{},  // TaskTimeInvested
		&sql.NullString{}, // Periodicity
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(values ...interface{}) error {
	if m, n := len(values), len(task.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field TaskName", values[0])
	} else if value.Valid {
		t.TaskName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field TaskLabel", values[1])
	} else if value.Valid {
		t.TaskLabel = value.String
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field TaskDate", values[2])
	} else if value.Valid {
		t.TaskDate = value.Time
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field RecurrentTask", values[3])
	} else if value.Valid {
		t.RecurrentTask = value.Bool
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field TaskTimeInvested", values[4])
	} else if value.Valid {
		t.TaskTimeInvested = new(int)
		*t.TaskTimeInvested = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Periodicity", values[5])
	} else if value.Valid {
		t.Periodicity = new(string)
		*t.Periodicity = value.String
	}
	return nil
}

// Update returns a builder for updating this Task.
// Note that, you need to call Task.Unwrap() before calling this method, if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", TaskName=")
	builder.WriteString(t.TaskName)
	builder.WriteString(", TaskLabel=")
	builder.WriteString(t.TaskLabel)
	builder.WriteString(", TaskDate=")
	builder.WriteString(t.TaskDate.Format(time.ANSIC))
	builder.WriteString(", RecurrentTask=")
	builder.WriteString(fmt.Sprintf("%v", t.RecurrentTask))
	if v := t.TaskTimeInvested; v != nil {
		builder.WriteString(", TaskTimeInvested=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := t.Periodicity; v != nil {
		builder.WriteString(", Periodicity=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
