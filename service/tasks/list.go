package service

import (
	"log"
	"todogolist/models"
	"todogolist/repository"
	"todogolist/utils"
)

//List all the objects on the DB.
func List(limit int, repository repository.Repository) string {
	taskFilterListDomainInstance := models.TaskFilterListDomain{Limit: limit}
	res, err := repository.ListTask(taskFilterListDomainInstance)
	if err != nil {
		log.Fatalf("%s", err)
	}

	printer := utils.NewTaskDomainPrinter(res)

	return printer.PrintTable().String()

}
