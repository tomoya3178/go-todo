package presentation

import (
	"fmt"
	"go-todo/application"
	"go-todo/infrastructure"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	dbName := "postgres"
	password := "password"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, dbName, port)
	var err error
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("failed to connect database", err.Error())
		return nil, err
	}
	return db, nil
}

func GetRouter() *gin.Engine {
	/*
		db, err := initDB()
		if err != nil {
			panic(err)
		}
	*/
	taskController := NewTaskController(application.NewTaskUseCase(infrastructure.NewTaskMockRepository()))
	router := gin.Default()
	router.GET("/tasks", taskController.NewTask())
	return router
}
