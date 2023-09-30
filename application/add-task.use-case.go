package application

import (
	"go-todo/domain"
)

type TaskUsecase interface {
	Create(name string, contents string) (*domain.Task, error)
}

type taskUseCase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUseCase(taskRepository domain.TaskRepository) TaskUsecase {
	return &taskUseCase{taskRepository: taskRepository}
}

func (useCase *taskUseCase) Create(name string, contents string) (*domain.Task, error) {
	task, err := domain.NewTask(name, contents)
	if err != nil {
		return nil, err
	}
	createdTask, err := useCase.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}
	return createdTask, nil
}
