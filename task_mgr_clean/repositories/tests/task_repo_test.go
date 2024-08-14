package tests

import (
	"context"
	"fmt"
	"goPractice/task_manager/domain"
	"goPractice/task_manager/repositories"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepoSuite struct {
	suite.Suite

	repository repositories.TaskRepository
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func (suite *TaskRepoSuite) SetupSuite() {
	fmt.Println("Setting up test suite")
	godotenv.Load("../../.env")
	options := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.TODO(), options)
	require.NoError(suite.T(), err, "Failed to connect to MongoDB")

	suite.client = client
	db := suite.client.Database("task-manager-test")
	suite.db = db
	suite.collection = db.Collection("tasks")
	suite.repository = repositories.NewTaskRepository(db)
}
func (suite *TaskRepoSuite) TearDownSuite() {
	//drop task collection
	defer suite.client.Disconnect(context.TODO())
	defer suite.collection.Drop(context.TODO())

}

// delete all tasks after each test for consistency
func (suite *TaskRepoSuite) TeaDownTest() {
	fmt.Println("Tearing down test")
	suite.collection.DeleteMany(context.TODO(), bson.M{})

}
func (suite *TaskRepoSuite) TestCreateTask() {

	task := domain.Task{
		ID:          primitive.NewObjectID(),
		UserId:      primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "todo",
	}
	_, err := suite.repository.CreateTask(task)
	require.NoError(suite.T(), err, "Failed to create task")
}

func (suite *TaskRepoSuite) TestGetTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		UserId:      primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "todo",
	}

	_, err := suite.repository.CreateTask(task)
	require.NoError(suite.T(), err, "Failed to create task")

	task, err = suite.repository.GetTask(task.ID.Hex())
	require.NoError(suite.T(), err, "Failed to get task")

	require.Equal(suite.T(), task.ID, task.ID)

}
func (suite *TaskRepoSuite) TestUpdateTask() {

	task := domain.Task{
		ID:          primitive.NewObjectID(),
		UserId:      primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "todo",
	}
	//updated task
	_, err := suite.repository.CreateTask(task)
	updatedTask := domain.Task{
		Title:       "Updated Title",
		Description: "task1 description updated",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "todo",
	}
	require.NoError(suite.T(), err, "Failed to create task")

	task, err = suite.repository.UpdateTask(task.ID.Hex(), updatedTask)

	require.NoError(suite.T(), err, "Failed to update task")

	require.Equal(suite.T(), updatedTask.Title, task.Title)
}

func (suite *TaskRepoSuite) TestGetUserTasks() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		UserId:      primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "todo",
	}

	_, err := suite.repository.CreateTask(task)
	require.NoError(suite.T(), err, "Failed to create task")

	tasks := suite.repository.GetUserTasks(task.UserId.Hex())
	require.Equal(suite.T(), 1, len(tasks))
}
func (suite *TaskRepoSuite) TestDeleteTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		UserId:      primitive.NewObjectID(),
		Title:       "task1",
		Description: "task1 description",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "todo",
	}

	_, err := suite.repository.CreateTask(task)
	require.NoError(suite.T(), err, "Failed to create task")

	err = suite.repository.DeleteTask(task.ID.Hex())
	require.NoError(suite.T(), err, "Failed to delete task")

}

func TestTaskRepo(t *testing.T) {
	suite.Run(t, new(TaskRepoSuite))
}
