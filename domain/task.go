package domain

import "github.com/google/uuid"

type Task struct {
	id       string
	name     string
	contents string
}

func NewTask(name, contents string) (*Task, error) {
	task := &Task{
		id:       uuid.New().String(),
		name:     name,
		contents: contents,
	}
	return task, nil
}
