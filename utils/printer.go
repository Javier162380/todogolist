package utils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"todogolist/models"
)

//Printer default type to print results in the terminal.
type Printer struct {
	SourceColumns []string
	ColumsWidth   map[string]int
	Data          []map[string]string
}

func resizeString(inputstring string, maxWidth int) string {

	if len(inputstring) <= maxWidth {
		return fmt.Sprintf("%s%s", inputstring, strings.Repeat(" ", maxWidth-len(inputstring)))
	}

	return inputstring[:maxWidth]

}

//PrintTable returns a buffer of string
func (p *Printer) PrintTable() *bytes.Buffer {
	b := new(bytes.Buffer)

	formatColumns := []string{}
	for _, column := range p.SourceColumns {
		formatColumn := resizeString(column, p.ColumsWidth[column])
		formatColumns = append(formatColumns, formatColumn)

	}

	for _, formatColumn := range formatColumns {
		b.WriteString(formatColumn)
		b.WriteString("\t")

	}

	b.WriteString("\n")

	for _, row := range p.Data {
		for _, column := range p.SourceColumns {
			record := row[column]
			formattedRecord := resizeString(record, p.ColumsWidth[column])
			b.WriteString(formattedRecord)
			b.WriteString("\t")

		}
		b.WriteString("\n")

	}

	return b

}

//NewTaskDomainPrinter generates a new task Printer with the default TaskDomain results.
func NewTaskDomainPrinter(tasks []models.TaskDomainResponse) *Printer {

	taskDataCollection := []map[string]string{}
	for _, task := range tasks {
		taskProccessed := make(map[string]string)
		taskProccessed["TaskID"] = task.TaskID
		taskProccessed["TaskName"] = task.TaskName
		taskProccessed["TaskLabel"] = task.TaskLabel
		taskProccessed["TaskDate"] = task.TaskDate.String()
		taskProccessed["RecurrentTask"] = strconv.FormatBool(task.RecurrentTask)
		taskProccessed["Periodicity"] = task.Periodicity
		taskProccessed["TaskTimeInvested"] = strconv.FormatInt(int64(task.TaskTimeInvested), 10)

		taskDataCollection = append(taskDataCollection, taskProccessed)

	}
	return &Printer{
		SourceColumns: []string{"TaskID", "TaskName", "TaskLabel", "TaskDate", "RecurrentTask", "Periodicity", "TaskTimeInvested"},
		ColumsWidth:   map[string]int{"TaskID": 8, "TaskName": 20, "TaskLabel": 15, "TaskDate": 25, "Periodicity": 12, "TaskTimeInvested": 20},
		Data:          taskDataCollection,
	}
}
