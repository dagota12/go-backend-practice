package data

import (
	"fmt"
	"goPractice/task_manager/models"
)

type TaskService interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTasks() map[string]models.Task
	GetTask(id string) (models.Task, error)
	UpdateTask(id string, task models.Task) (models.Task, error)
	DeleteTask(id string) error
}
type taskService struct {
	tasks map[string]models.Task
}

func NewTaskService() TaskService {
	return &taskService{
		tasks: make(map[string]models.Task),
	}
}

// CreateTask implements TaskService.
func (t *taskService) CreateTask(task models.Task) (models.Task, error) {

	if _, ok := t.tasks[task.ID]; ok {
		return models.Task{}, fmt.Errorf("task with id %s already exists", task.ID)
	}
	t.tasks[task.ID] = task
	return task, nil
}

// DeleteTask implements TaskService.
func (t *taskService) DeleteTask(id string) error {
	if _, ok := t.tasks[id]; !ok {
		return fmt.Errorf("task with id %s not found", id)
	}
	delete(t.tasks, id)
	return nil
}

// GetTask implements TaskService.
func (t *taskService) GetTask(id string) (models.Task, error) {
	if _, ok := t.tasks[id]; !ok {
		return models.Task{}, fmt.Errorf("task with id %s not found", id)
	}
	return t.tasks[id], nil
}

// GetTasks implements TaskService.
func (t *taskService) GetTasks() map[string]models.Task {
	return t.tasks
}

// UpdateTask implements TaskService.
func (t *taskService) UpdateTask(id string, task models.Task) (models.Task, error) {
	if _, ok := t.tasks[id]; !ok {
		return models.Task{}, fmt.Errorf("task with id %s not found", id)
	}
	t.tasks[id] = task
	return task, nil
}
