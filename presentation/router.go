package presentation

import (
	"fmt"
	"go-todo/application"
	"go-todo/infrastructure"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	fmt.Print(123)
	router := gin.Default()
	taskController := NewTaskController(application.NewTaskUseCase(infrastructure.NewTaskRepository()))
	router.GET("/tasks", taskController.NewTask())
	return router
}
