package data

import (
	"context"
	"goPractice/task_manager/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTasks() []models.Task
	GetTask(id string) (models.Task, error)
	UpdateTask(id string, task models.Task) (models.Task, error)
	DeleteTask(id string) error
}
type taskService struct {
	tasks *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *taskService {
	return &taskService{
		tasks: collection,
	}
}

// CreateTask implements TaskService.
func (t *taskService) CreateTask(task models.Task) (models.Task, error) {
	res, err := t.tasks.InsertOne(context.TODO(), task)

	if err != nil {
		log.Println(err.Error())
		return models.Task{}, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID)
	return task, nil
}

// DeleteTask implements TaskService.
func (t *taskService) DeleteTask(id string) error {

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
		return mongo.ErrNoDocuments
	}
	return nil
}

// GetTask implements TaskService.
func (t *taskService) GetTask(id string) (models.Task, error) {
	log.Println("getting task...")
	oId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err.Error())
		return models.Task{}, err
	}
	var task models.Task
	res := t.tasks.FindOne(context.TODO(), bson.M{"_id": oId})

	err = res.Decode(&task)
	if err != nil {
		log.Println(err.Error())
		return models.Task{}, err
	}

	return task, nil
}

// GetTasks implements TaskService.
func (t *taskService) GetTasks() []models.Task {

	res, err := t.tasks.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println(err.Error())
		return []models.Task{}
	}
	var tasks []models.Task

	for res.Next(context.TODO()) {
		var task models.Task
		err := res.Decode(&task)
		if err != nil {
			log.Println(err.Error())
			return []models.Task{}
		}

		tasks = append(tasks, task)
	}
	return tasks
}

// UpdateTask implements TaskService.
func (t *taskService) UpdateTask(id string, task models.Task) (models.Task, error) {
	log.Println("updating task...")
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return models.Task{}, err
	}
	res, err := t.tasks.UpdateOne(context.TODO(), bson.M{"_id": oId}, bson.D{{Key: "$set", Value: task}})
	if err != nil {
		log.Println(err.Error())
		return models.Task{}, err
	}
	if res.MatchedCount < 1 {
		return models.Task{}, mongo.ErrNoDocuments
	}
	task.ID = oId
	return task, nil
}
