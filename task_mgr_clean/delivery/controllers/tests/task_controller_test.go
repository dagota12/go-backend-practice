package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goPractice/task_manager/delivery/controllers"
	"goPractice/task_manager/domain"
	"goPractice/task_manager/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseMock struct {
	mock.Mock
}

func (m *TaskUsecaseMock) CreateTask(task domain.Task) (domain.Task, error) {
	args := m.Called(task)

	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *TaskUsecaseMock) GetTask(id string) (domain.Task, error) {
	args := m.Called(id)

	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *TaskUsecaseMock) GetTasks() []domain.Task {
	args := m.Called()

	return args.Get(0).([]domain.Task)
}
func (m *TaskUsecaseMock) GetUserTasks(user_id string) []domain.Task {
	args := m.Called(user_id)

	return args.Get(0).([]domain.Task)
}

func (m *TaskUsecaseMock) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	args := m.Called(id, task)

	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *TaskUsecaseMock) DeleteTask(id string) error {
	args := m.Called(id)

	return args.Error(0)
}

type TaskControllerSuite struct {
	suite.Suite
	controller controllers.TaskController
	usecase    *TaskUsecaseMock
	httpServer *httptest.Server
}

func (s *TaskControllerSuite) SetupSuite() {
	usecase := new(TaskUsecaseMock)
	controller := *controllers.NewTaskController(usecase)

	router := gin.Default()
	taskG := router.Group("/task")
	taskG.Use(mocks.Authorize("admin"))

	taskG.GET("/", controller.GetUserTasks)
	taskG.GET("/:id", controller.GetTask)

	//protected
	adminOnly := taskG.Group("/")
	// adminOnly.Use(infrastructure.Authorize("admin"))

	adminOnly.GET("/all", controller.GetTasks)
	adminOnly.POST("/", controller.CreateTask)
	adminOnly.PUT("/:id", controller.UpdateTask)
	adminOnly.DELETE("/:id", controller.DeleteTask)

	s.httpServer = httptest.NewServer(router)
	fmt.Println(s.httpServer.URL)

	s.usecase = usecase
	s.controller = controller

	// s.usecase.On("CreateTask", task).Return(task, nil)

}

func (s *TaskControllerSuite) TearDownSuite() {
	defer s.httpServer.Close()
}
func (s *TaskControllerSuite) TestCreateTask() {

	task := domain.Task{
		ID:          primitive.NewObjectID(),
		UserId:      primitive.NewObjectID(),
		Title:       "demo Title",
		Description: "erathostes",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "pending",
	}

	s.usecase.On("CreateTask", mock.Anything).Return(task, nil)

	body, err := json.Marshal(task)
	require.NoError(s.T(), err, "Failed to marshal task")

	response, err := http.Post(s.httpServer.URL+"/task", "application/json", bytes.NewBuffer(body))
	require.NoError(s.T(), err, "faild to create task can'tsend request")
	defer response.Body.Close()

	require.Equal(s.T(), http.StatusCreated, response.StatusCode)

	s.usecase.AssertExpectations(s.T())

}
func (s *TaskControllerSuite) TestGetTask() {
	s.usecase.On("GetTask", "1").Return(domain.Task{Title: "task 1"}, nil)

	response, err := http.Get(s.httpServer.URL + "/task/1")

	require.NoError(s.T(), err, "faild to make request to get task")
	defer response.Body.Close()
	require.Equal(s.T(), http.StatusOK, response.StatusCode)
	s.usecase.AssertExpectations(s.T())
}
func (s *TaskControllerSuite) TestGetUserTasks() {
	s.usecase.On("GetUserTasks", "1").Return([]domain.Task{{Title: "task 1"}, {Title: "task 2"}}, nil)

	response, err := http.Get(s.httpServer.URL + "/task")

	require.NoError(s.T(), err, "faild to make request to get task")
	defer response.Body.Close()

	tasks := []domain.Task{}
	json.NewDecoder(response.Body).Decode(&tasks)
	s.T().Log(tasks)

	require.Equal(s.T(), http.StatusOK, response.StatusCode)
	s.usecase.AssertExpectations(s.T())
}
func (s *TaskControllerSuite) TestGetTasks() {
	s.usecase.On("GetTasks").Return([]domain.Task{{Title: "task 1"}, {Title: "task 2"}}, nil)

	response, err := http.Get(s.httpServer.URL + "/task/all")
	require.NoError(s.T(), err, "faild to make request to get tasks")

	defer response.Body.Close()

	require.Equal(s.T(), http.StatusOK, response.StatusCode)
	s.usecase.AssertExpectations(s.T())
}
func (s *TaskControllerSuite) TestUpdateTask() {

	task := domain.Task{
		Title:       "demo Title",
		Description: "erathostes",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      "pending",
	}

	s.usecase.On("UpdateTask", "1", mock.Anything).Return(task, nil)

	body, err := json.Marshal((task))
	require.NoError(s.T(), err, "Failed to marshal task")

	req, err := http.NewRequest("PUT", s.httpServer.URL+"/task/1", bytes.NewBuffer(body))
	require.NoError(s.T(), err, "faild to make request to get task")

	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	require.NoError(s.T(), err, "faild to make request to update task")
	defer response.Body.Close()

	require.Equal(s.T(), http.StatusOK, response.StatusCode)

}
func (s *TaskControllerSuite) TestDeleteTask() {

	s.usecase.On("DeleteTask", "1").Return(nil)

	req, err := http.NewRequest("DELETE", s.httpServer.URL+"/task/1", nil)

	require.NoError(s.T(), err, "faild to make request to get task")

	req.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	require.NoError(s.T(), err, "faild to make task delete request")

	require.Equal(s.T(), http.StatusNoContent, response.StatusCode)
}
func TestTaskController(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
