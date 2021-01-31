package api

import (
	"net/http"
	"todogolist/repository"
	service "todogolist/service/tasks"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	TaskName      string `json:"task_name"`
	TaskLabel     string `json:"task_label"`
	TaskStartDate string `json:"task_start_date"`
	TaskEndDate   string `json:"task_end_date"`
	Limit         int    `json"limit"`
}

//GetHandler get a task via API.
func GetHandler(repository repository.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gr := &getRequest{}

		if err := ctx.BindJSON(gr); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return

		}

		// TODO: Create an interface for returning a json.
		gresponse, err := service.Get(gr.TaskName, gr.TaskLabel, gr.TaskStartDate, gr.TaskEndDate, gr.Limit, repository)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(200, gin.H{"message": gresponse})

	}

}
