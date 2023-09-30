package infrastructure

import (
	"fmt"
	"go-todo/domain"
)

type TaskMockRepository struct {
}

func NewTaskMockRepository() domain.TaskRepository {
	return &TaskMockRepository{}
}

func (taskRepository *TaskMockRepository) Create(task *domain.Task) error {
	fmt.Printf("%+v", task)
	return nil
}
