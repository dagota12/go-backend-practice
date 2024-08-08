package usecases

import (
	"goPractice/task_manager/domain"
	"goPractice/task_manager/repositories"
)

type TaksUsecase interface {
	GetTask(id string) (domain.Task, error)
	GetTasks() []domain.Task
	GetUserTasks(user_id string) []domain.Task
	CreateTask(task domain.Task) (domain.Task, error)
	UpdateTask(id string, task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
}

type taskUsecase struct {
	taskRepository repositories.TaskRepository
}

func NewTaskUsecase(tasks repositories.TaskRepository) TaksUsecase {
	return &taskUsecase{taskRepository: tasks}
}

// GetTask implements TaksUsecase.
func (t *taskUsecase) GetTask(id string) (domain.Task, error) {
	return t.taskRepository.GetTask(id)
}

// CreateTask implements TaksUsecase.
func (t *taskUsecase) CreateTask(task domain.Task) (domain.Task, error) {
	task, err := t.taskRepository.CreateTask(task)
	return task, err
}

// GetTask implements TaksUsecase.
func (t *taskUsecase) GetTasks() []domain.Task {
	return t.taskRepository.GetTasks()
}

// GetUserTasks implements TaksUsecase.
func (t *taskUsecase) GetUserTasks(user_id string) []domain.Task {
	return t.taskRepository.GetUserTasks(user_id)
}

// UpdateTask implements TaksUsecase.
func (t *taskUsecase) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	return t.taskRepository.UpdateTask(id, task)
}

// DeleteTask implements TaksUsecase.
func (t *taskUsecase) DeleteTask(id string) error {
	return t.taskRepository.DeleteTask(id)
}
