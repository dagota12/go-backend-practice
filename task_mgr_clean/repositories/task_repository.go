package repositories

//import mongidb driver
import (
	"context"
	"errors"
	"goPractice/task_manager/domain"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	CreateTask(task domain.Task) (domain.Task, error)
	GetTasks() []domain.Task
	GetUserTasks(user_id string) []domain.Task
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
	res, err := t.tasks.InsertOne(context.TODO(), task)

	if err != nil {
		log.Println(err.Error())
		return domain.Task{}, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID)
	return task, nil
}

// GetTask implements TaskRepository.
func (t *taskRepository) GetTask(id string) (domain.Task, error) {
	log.Println("getting task...")
	oId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err.Error())
		return domain.Task{}, err
	}
	var task domain.Task
	res := t.tasks.FindOne(context.TODO(), bson.M{"_id": oId})

	err = res.Decode(&task)
	if err != nil {
		log.Println(err.Error())
		return domain.Task{}, errors.New("task not found")
	}

	return task, nil
}

// GetTasks implements TaskRepository.
func (t *taskRepository) GetTasks() []domain.Task {
	cursor, err := t.tasks.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println(err.Error())
		return []domain.Task{}
	}
	var tasks []domain.Task

	for cursor.Next(context.TODO()) {
		var task domain.Task
		err := cursor.Decode(&task)
		if err != nil {
			log.Println(err.Error())
			return []domain.Task{}
		}

		tasks = append(tasks, task)
	}
	return tasks
}

// GetUserTasks implements TaskRepository.
func (t *taskRepository) GetUserTasks(user_id string) []domain.Task {
	oId, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		log.Println(err.Error())
		return []domain.Task{}
	}
	cursor, err := t.tasks.Find(context.TODO(), bson.M{"user_id": oId})
	if err != nil {
		log.Println(err.Error())
		return []domain.Task{}
	}
	var tasks []domain.Task

	for cursor.Next(context.TODO()) {
		var task domain.Task
		err := cursor.Decode(&task)
		if err != nil {
			log.Println(err.Error())
			return []domain.Task{}
		}

		tasks = append(tasks, task)
	}
	return tasks
}

// UpdateTask implements TaskRepository.
func (t *taskRepository) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	log.Println("updating task...")
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return domain.Task{}, err
	}
	res, err := t.tasks.UpdateOne(context.TODO(), bson.M{"_id": oId}, bson.D{{Key: "$set", Value: task}})
	if err != nil {
		log.Println(err.Error())
		return domain.Task{}, err
	}
	if res.MatchedCount < 1 {
		return domain.Task{}, mongo.ErrNoDocuments
	}
	task.ID = oId
	return task, nil
}

// DeleteTask implements TaskRepository.
func (t *taskRepository) DeleteTask(id string) error {
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	res, err := t.tasks.DeleteOne(context.TODO(), bson.M{"_id": oId})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if res.DeletedCount < 1 {
		return errors.New("task not found")
	}
	return nil
}
