package service

import "github.com/begenov/TaskFlow/task-app/internal/storage"

type Tasks interface {
}

type Service struct {
	Task Tasks
}

func NewService(s storage.Storage) *Service {
	return &Service{
		Task: NewTaskService(s.Task),
	}
}
