package domain

type TaskRepository interface {
	Create(task *Task) error
}
