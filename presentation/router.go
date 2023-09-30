package presentation

import (
	"go-todo/application"
	"go-todo/infrastructure"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	taskController := NewTaskController(application.NewTaskUseCase(infrastructure.NewTaskRepository()))
	router.GET("/tasks", taskController.NewTask())
	return router
}
