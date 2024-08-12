package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goPractice/task_manager/delivery/controllers"
	"goPractice/task_manager/domain"
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
	// taskG.Use(infrastructure.Authorize("user", "admin"))

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

func (s *TaskControllerSuite) TearDownTest() {
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

	s.usecase.On("CreateTask", task).Return(task, nil)

	body, err := json.Marshal(task)
	require.NoError(s.T(), err, "Failed to marshal task")

	response, err := http.Post(s.httpServer.URL+"/task", "application/json", bytes.NewBuffer(body))
	require.NoError(s.T(), err, "faild to create task can'tsend request")
	defer response.Body.Close()

	require.Equal(s.T(), http.StatusCreated, response.StatusCode)

	s.usecase.AssertExpectations(s.T())

}

func TestTaskCntroller(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
