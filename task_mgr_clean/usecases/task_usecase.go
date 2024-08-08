package usecases

import (
	"goPractice/task_manager/domain"
	"goPractice/task_manager/repositories"
)

type TaksUsecase interface {
	GetTask() []domain.Task
	GetUserTasks(id string) []domain.Task
	CreateTask(task domain.Task) (domain.Task, error)
	UpdateTask(id string, task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
}

type taskUsecase struct {
	tasks repositories.TaskRepository
}

// CreateTask implements TaksUsecase.
func (t *taskUsecase) CreateTask(task domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

// DeleteTask implements TaksUsecase.
func (t *taskUsecase) DeleteTask(id string) error {
	panic("unimplemented")
}

// GetTask implements TaksUsecase.
func (t *taskUsecase) GetTask() []domain.Task {
	panic("unimplemented")
}

// GetUserTasks implements TaksUsecase.
func (t *taskUsecase) GetUserTasks(id string) []domain.Task {
	panic("unimplemented")
}

// UpdateTask implements TaksUsecase.
func (t *taskUsecase) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

func NewTaskUsecase(tasks repositories.TaskRepository) TaksUsecase {
	return &taskUsecase{tasks: tasks}
}
