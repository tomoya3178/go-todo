package domain

type TaskRepository interface {
	Create(task *Task) (*Task, error)
}
