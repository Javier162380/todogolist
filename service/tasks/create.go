package service

import (
	"time"
	"todogolist/models"
	"todogolist/repository"
	"todogolist/service"

	"github.com/jinzhu/copier"
)

var validTaskLabels = `(unknow$|sport$|work$|tech$|friends$)`
var validTaskPeriodicity = `(w$|weekly$|d$|daily$|m$|monthly$)`

func generateDateSeries(startDate time.Time, endDate time.Time, periodicity string) []time.Time {

	dateRanges := []time.Time{}
	switch periodicity {
	case "w", "weekly":
		dateRange := int(endDate.Sub(startDate).Hours() / (24 * 7))

		for date := 1; date <= dateRange; date++ {
			dateRanges = append(dateRanges, startDate.AddDate(0, 0, 7*date))

		}

		return dateRanges

	case "m", "monthly":
		dateRange := int(endDate.Sub(startDate).Hours() / (24 * 30))

		for date := 1; date <= dateRange; date++ {
			dateRanges = append(dateRanges, startDate.AddDate(0, date, 0))

		}
		return dateRanges

	case "d", "daily":
		dateRange := int(endDate.Sub(startDate).Hours() / 24)

		for date := 1; date <= dateRange; date++ {
			dateRanges = append(dateRanges, startDate.AddDate(date, 0, 0))

		}
		return dateRanges

	default:
		return dateRanges

	}
}

func generateTaskDomainSeries(taskList []models.TaskDomain, taskDomain models.TaskDomain, dateseries []time.Time) []models.TaskDomain {

	for _, date := range dateseries {
		newtaskDomain := &models.TaskDomain{}
		copier.Copy(newtaskDomain, taskDomain)
		newtaskDomain.TaskDate = date
		taskList = append(taskList, *newtaskDomain)

	}

	return taskList
}

//Create a new task in the repository db.
func Create(taskName string, taskLabel string, taskDate string, taskTimeInvested int, taskPeriodicity string, debugMode bool, repository repository.Repository) (string, error) {
	taskDateFormated, err := service.FormatTaskTimestamp(taskDate)
	if err != nil {
		return "", err
	}

	taskInstances := []models.TaskDomain{}

	if taskLabel == "" {
		taskLabel = "unknow"
	}

	err = service.ValidateStringField(taskLabel, validTaskLabels, `Task Label do not match a valid task label`)

	if err != nil {
		return "", err
	}
	if taskPeriodicity != "" {
		err = service.ValidateStringField(taskPeriodicity, validTaskPeriodicity, `Task Periodicity do not match a valid task periodicity`)
		if err != nil {
			return "", err
		}

		dateSeries := generateDateSeries(taskDateFormated, taskDateFormated.AddDate(1, 0, 0), taskPeriodicity)

		taskInstance := &models.TaskDomain{
			TaskName:         taskName,
			TaskLabel:        taskLabel,
			TaskDate:         taskDateFormated,
			TaskTimeInvested: taskTimeInvested,
			Periodicity:      taskPeriodicity,
		}

		taskInstances := generateTaskDomainSeries(taskInstances, *taskInstance, dateSeries)

		taskInstances = append(taskInstances, *taskInstance)

		return repository.CreateTask(*&taskInstances, debugMode)

	} else {
		taskInstance := &models.TaskDomain{
			TaskName:         taskName,
			TaskLabel:        taskLabel,
			TaskDate:         taskDateFormated,
			TaskTimeInvested: taskTimeInvested,
			Periodicity:      taskPeriodicity,
		}

		taskInstances = append(taskInstances, *taskInstance)

	}

	return repository.CreateTask(*&taskInstances, debugMode)

}
