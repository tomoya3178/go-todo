package infrastructure

import (
	"fmt"
	"go-todo/domain"
)

type TaskRepository struct {
}

func NewTaskRepository() domain.TaskRepository {
	return &TaskRepository{}
}

func (taskRepository *TaskRepository) Create(task *domain.Task) (*domain.Task, error) {
	fmt.Printf("%+v", task)
	return task, nil
}
