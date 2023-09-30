package presentation

import (
	"go-todo/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase application.TaskUseCase
}

func NewTaskController(taskUseCase application.TaskUseCase) *TaskController {
	controller := new(TaskController)
	controller.taskUseCase = taskUseCase
	return controller
}

func (controller *TaskController) NewTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.taskUseCase.Create("name", "contents")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
