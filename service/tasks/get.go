package service

import (
	"todogolist/models"
	"todogolist/repository"
	"todogolist/service"
	_ "todogolist/service"
	"todogolist/utils"
)

//Get task results from the repository
func Get(taskFilterName string, taskFilterLabel string, startDate string, endDate string, limit int, repository repository.Repository) (string, error) {

	taskFilterModelInstance := models.TaskFilterDomain{
		TaskName:  taskFilterName,
		TaskLabel: taskFilterLabel,
	}

	if startDate != "" {
		startFormatDate, err := service.FormatTaskTimestamp(startDate)
		if err != nil {
			return "", err
		}
		taskFilterModelInstance.StartDate = startFormatDate
	}

	if endDate != "" {
		endFormatDate, err := service.FormatTaskTimestamp(endDate)
		if err != nil {
			return "", err
		}
		taskFilterModelInstance.EndDate = endFormatDate
	}

	if limit != 0 {
		taskFilterModelInstance.Limit = limit

	}

	res, err := repository.GetTask(taskFilterModelInstance)
	if err != nil {
		return "", err
	}

	printer := utils.NewTaskDomainPrinter(res)

	return printer.PrintTable().String(), nil

}
