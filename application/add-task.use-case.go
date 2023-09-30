package application

import (
	"go-todo/domain"
)

type TaskUseCase interface {
	Create(name string, contents string) error
}

type taskUseCase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUseCase(taskRepository domain.TaskRepository) TaskUseCase {
	return &taskUseCase{taskRepository: taskRepository}
}

func (useCase *taskUseCase) Create(name string, contents string) error {
	task, err := domain.NewTask(name, contents)
	if err != nil {
		return err
	}
	err = useCase.taskRepository.Create(task)
	if err != nil {
		return err
	}
	return nil
}
