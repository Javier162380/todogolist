package service

import (
	"log"
	"todogolist/models"
	"todogolist/repository"
)

//Delete a task in a repository
func Delete(taskName string, repository repository.Repository) string {
	if taskName == "" {
		return "Empty task nothing is deleted from the repository"
	}

	deletetaskInstance := &models.TaskDeleteDomain{
		TaskName: taskName,
	}
	resp, err := repository.DeleteTask(*deletetaskInstance)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}
