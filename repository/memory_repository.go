package repository

import (
	"fmt"
	"time"
	"todogolist/models"
)

type memoryApp struct {
	db string
}

func (m *memoryApp) Setup(dryRun bool) (string, error) {

	switch dryRun {
	case false:
		return "DB Setup Succesfully", nil
	default:
		return "BEGIN;DB Setup Succesfully;COMMIT", nil
	}
}

func (m *memoryApp) CreateTask(taskDomains []models.TaskDomain, debugMode bool) (string, error) {
	switch debugMode {
	case false:
		return "Task Created Succesfully", nil

	default:
		return fmt.Sprintf(`Task Created Succesfully with: %v`, taskDomains), nil

	}
}

func (m *memoryApp) GetTask(models.TaskFilterDomain) ([]models.TaskDomainResponse, error) {
	return []models.TaskDomainResponse{
		{TaskID: "1234", TaskName: "test1", TaskLabel: "work", TaskDate: time.Date(2020, 1, 1, 10, 10, 10, 10, time.UTC), RecurrentTask: true, Periodicity: "w", TaskTimeInvested: 600},
		{TaskID: "1235", TaskName: "test2", TaskLabel: "sport", TaskDate: time.Date(2020, 2, 10, 10, 10, 10, 10, time.UTC), RecurrentTask: false, Periodicity: "", TaskTimeInvested: 10},
		{TaskID: "1236", TaskName: "test3", TaskLabel: "tech", TaskDate: time.Date(2020, 8, 1, 10, 10, 10, 10, time.UTC), RecurrentTask: false, Periodicity: "w", TaskTimeInvested: 10},
	}, nil
}

func (m *memoryApp) ListTask(models.TaskFilterListDomain) ([]models.TaskDomainResponse, error) {
	sampleEvents := []models.TaskDomainResponse{
		{TaskID: "1234", TaskName: "test1", TaskLabel: "work", TaskDate: time.Date(2020, 1, 1, 10, 10, 10, 10, time.UTC), RecurrentTask: true, Periodicity: "w", TaskTimeInvested: 600},
		{TaskID: "1235", TaskName: "test2", TaskLabel: "sport", TaskDate: time.Date(2020, 2, 10, 10, 10, 10, 10, time.UTC), RecurrentTask: false, Periodicity: "", TaskTimeInvested: 10},
		{TaskID: "1236", TaskName: "test3", TaskLabel: "tech", TaskDate: time.Date(2020, 8, 1, 10, 10, 10, 10, time.UTC), RecurrentTask: false, Periodicity: "w", TaskTimeInvested: 10},
		{TaskID: "1234", TaskName: "test1", TaskLabel: "work", TaskDate: time.Date(2020, 1, 1, 10, 10, 10, 10, time.UTC), RecurrentTask: true, Periodicity: "w", TaskTimeInvested: 600},
		{TaskID: "1235", TaskName: "test2", TaskLabel: "sport", TaskDate: time.Date(2020, 2, 10, 10, 10, 10, 10, time.UTC), RecurrentTask: false, Periodicity: "", TaskTimeInvested: 10},
		{TaskID: "1236", TaskName: "test3", TaskLabel: "tech", TaskDate: time.Date(2020, 8, 1, 10, 10, 10, 10, time.UTC), RecurrentTask: false, Periodicity: "w", TaskTimeInvested: 10},
	}

	return sampleEvents, nil
}

func (m *memoryApp) DeleteTask(models.TaskDeleteDomain) (string, error) {

	return "Task Deleted Succesfully", nil

}

func newMemoryRepository() Repository {
	return &memoryApp{}

}
