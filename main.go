package main

import (
	"go-todo/application"
	"go-todo/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	taskUsecase := application.NewTaskUseCase(infrastructure.NewTaskRepository())
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		taskUsecase.Create("name", "contents")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":3000")
}
