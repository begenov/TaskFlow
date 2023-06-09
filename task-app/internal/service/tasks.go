package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

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
	log.Println(task.UserID, "check task1")

	if err := checkTask(task); err != nil {
		return err
	}
	log.Println(task.UserID, "check task2")
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

func (t *tasksService) UpdateTask(ctx context.Context, task models.Todo, taskID int, userID int) error {
	dbtask, err := t.task.TaskByID(ctx, taskID)

	if err != nil {
		return err
	}
	if dbtask.UserID == userID {

		if strings.TrimSpace(task.Title) == "" {
			task.Title = dbtask.Title
		}

		if strings.TrimSpace(task.Description) == "" {
			task.Description = dbtask.Description
		}
		task.CreatedAt = time.Now()
		task.ID = taskID
		log.Println("services", task)
		if err = t.task.UpdateTask(ctx, task); err != nil {
			return err
		}

		return nil
	} else {
		return fmt.Errorf("incorct error name")
	}
}

func (t *tasksService) DeleteTask(ctx context.Context, taskID int, userID int) error {

	dbtask, err := t.task.TaskByID(ctx, taskID)

	if err != nil {
		return err
	}
	if dbtask.UserID == userID {

		if err := t.task.DeleteTask(ctx, taskID); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("incorct error name")
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
