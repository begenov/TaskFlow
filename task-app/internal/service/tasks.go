package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/begenov/TaskFlow/pkg/models"
	"github.com/begenov/TaskFlow/task-app/internal/storage"
)

type tasksService struct {
	task storage.Task
}

func NewTaskService(task storage.Task) Tasks {
	return &tasksService{
		task: task,
	}
}

func (t *tasksService) CreateTask(ctx context.Context, task models.Todo) error {

	if err := checkTask(task); err != nil {
		return err
	}

	if err := t.task.CreateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *tasksService) TaskByID(ctx context.Context, id int) (models.Todo, error) {

	if id <= 0 {
		return models.Todo{}, fmt.Errorf("does not exist id = %d", id)
	}

	task, err := t.task.TaskByID(ctx, id)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (t *tasksService) AllTasks(ctx context.Context) ([]models.Todo, error) {

	tasks, err := t.task.AllTask(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *tasksService) UpdateTask(ctx context.Context, task models.Todo) error {
	dbtask, err := t.task.TaskByID(ctx, task.ID)

	if err != nil {
		return err
	}

	if strings.TrimSpace(task.Title) == "" {
		task.Title = dbtask.Title
	}

	if strings.TrimSpace(task.Description) == "" {
		task.Description = dbtask.Description
	}

	if err = t.task.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *tasksService) DeleteTask(ctx context.Context, id int) error {
	if err := t.task.DeleteTask(ctx, id); err != nil {
		return err
	}
	return nil
}

func checkTask(task models.Todo) error {
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("incorecct: title empty")
	}
	if strings.TrimSpace(task.Description) == "" {
		return errors.New("incorrect: description empty")
	}

	return nil
}
