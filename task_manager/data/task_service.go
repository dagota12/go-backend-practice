package data

import (
	"fmt"
	"goPractice/task_manager/models"
	"time"
)

type TaskService interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTasks() map[string]models.Task
	GetTask(id string) (models.Task, error)
	UpdateTask(id string, task models.Task) (models.Task, error)
	DeleteTask(id string) error
}
type taskService struct {
	Tasks map[string]models.Task
}

func NewTaskService() *taskService {
	return &taskService{
		Tasks: map[string]models.Task{
			"1": {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
			"2": {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
			"3": {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
		},
	}
}

// CreateTask implements TaskService.
func (t *taskService) CreateTask(task models.Task) (models.Task, error) {

	if _, ok := t.Tasks[task.ID]; ok {
		return models.Task{}, fmt.Errorf("task with id %s already exists", task.ID)
	}
	t.Tasks[task.ID] = task
	return task, nil
}

// DeleteTask implements TaskService.
func (t *taskService) DeleteTask(id string) error {
	if _, ok := t.Tasks[id]; !ok {
		return fmt.Errorf("task with id %s not found", id)
	}
	delete(t.Tasks, id)
	return nil
}

// GetTask implements TaskService.
func (t *taskService) GetTask(id string) (models.Task, error) {
	if _, ok := t.Tasks[id]; !ok {
		return models.Task{}, fmt.Errorf("task with id %s not found", id)
	}
	return t.Tasks[id], nil
}

// GetTasks implements TaskService.
func (t *taskService) GetTasks() map[string]models.Task {
	return t.Tasks
}

// UpdateTask implements TaskService.
func (t *taskService) UpdateTask(id string, task models.Task) (models.Task, error) {
	if _, ok := t.Tasks[id]; !ok {
		return models.Task{}, fmt.Errorf("task with id %s not found", id)
	}
	if task.ID != id {
		return models.Task{}, fmt.Errorf("task ID in request body (%s) does not match task ID in URL (%s)", task.ID, id)
	}
	t.Tasks[id] = task
	return task, nil
}
