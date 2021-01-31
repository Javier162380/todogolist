package api

import (
	"net/http"
	"todogolist/repository"
	service "todogolist/service/tasks"

	"github.com/gin-gonic/gin"
)

type taskRequest struct {
	TaskName         string `json:"task_name" binding:"required"`
	TaskLabel        string `json:"task_label"`
	TaskDate         string `json:"task_date"`
	TaskTimeInvested int    `json:"task_time_invested"`
	TaskPeriodicity  string `json:"task_periodicity"`
}

//CreateHandler create a task via API.
func CreateHandler(repository repository.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		th := &taskRequest{}

		if err := ctx.BindJSON(th); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return

		}

		_, err := service.Create(th.TaskName, th.TaskLabel, th.TaskDate, th.TaskTimeInvested, th.TaskPeriodicity, false, repository)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(200, gin.H{"message": "Task/s created succesfully"})

	}

}
