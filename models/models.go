package models

import "time"

// TaskDomain model
type TaskDomain struct {
	TaskName         string
	TaskLabel        string
	TaskDate         time.Time
	Periodicity      string
	TaskTimeInvested int
}

// TaskDomainResponse response
type TaskDomainResponse struct {
	TaskID           string
	TaskName         string
	TaskLabel        string
	TaskDate         time.Time
	RecurrentTask    bool
	Periodicity      string
	TaskTimeInvested int
}

// TaskFilterDomain model
type TaskFilterDomain struct {
	TaskName  string
	TaskLabel string
	StartDate time.Time
	EndDate   time.Time
	Limit     int
}

// TaskFilterListDomain model
type TaskFilterListDomain struct {
	Limit int
}

// TaskDeleteDomain model
type TaskDeleteDomain struct {
	TaskName string
}

// TaskStatsDomain model
type TaskStatsDomain struct {
	StartDate time.Time
	EndDate   time.Time
	Limit     int
}

// TaskStatsResponse model it is an aggregate over the task domain model.
type TaskStatsResponse struct {
	TaskDate              time.Time
	TaskLabel             string
	NumberofTasks         int
	NumberofPeriodicTasks int
	TaskTimeInvested      int
}
