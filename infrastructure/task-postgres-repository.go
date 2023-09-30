package infrastructure

import (
	"fmt"
	"go-todo/domain"

	"gorm.io/gorm"
)

type TaskPostgresRepository struct {
	db *gorm.DB
}

func NewTaskPostgresRepository(db *gorm.DB) domain.TaskRepository {
	return &TaskPostgresRepository{
		db,
	}
}

func (taskRepository *TaskPostgresRepository) Create(task *domain.Task) error {
	fmt.Printf("%+v", task)
	taskRepository.db.Create(&task)
	return nil
}
