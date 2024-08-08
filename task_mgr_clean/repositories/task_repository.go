package repositories

//import mongidb driver
import (
	"goPractice/task_manager/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	CreateTask(task domain.Task) (domain.Task, error)
	GetTasks() []domain.Task
	GetUserTasks(id string) []domain.Task
	GetTask(id string) (domain.Task, error)
	UpdateTask(id string, task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
}

type taskRepository struct {
	db    *mongo.Database
	tasks *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepository{db: db, tasks: db.Collection("tasks")}
}

// CreateTask implements TaskRepository.
func (t *taskRepository) CreateTask(task domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

// DeleteTask implements TaskRepository.
func (t *taskRepository) DeleteTask(id string) error {
	panic("unimplemented")
}

// GetTask implements TaskRepository.
func (t *taskRepository) GetTask(id string) (domain.Task, error) {
	panic("unimplemented")
}

// GetTasks implements TaskRepository.
func (t *taskRepository) GetTasks() []domain.Task {
	panic("unimplemented")
}

// GetUserTasks implements TaskRepository.
func (t *taskRepository) GetUserTasks(id string) []domain.Task {
	panic("unimplemented")
}

// UpdateTask implements TaskRepository.
func (t *taskRepository) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

// getTasks implements TaskRepository.
