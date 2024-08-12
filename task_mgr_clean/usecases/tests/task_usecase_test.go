package tests

import (
	"goPractice/task_manager/domain"
	"goPractice/task_manager/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//here we are testing the usecase layer
/*
	there are 2 versions of the mock repo one inside the mocks package
	and the other one is defined here

	since we are testing the usecase layer we need to mock the repository layer
	mocking is like trying to copy ones action
		go test -v -run=TestTaksUsecase
*/

type TaskRepoMock struct {
	mock.Mock
}

func (m *TaskRepoMock) CreateTask(task domain.Task) (domain.Task, error) {
	args := m.Called(task)

	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *TaskRepoMock) GetTasks() []domain.Task {
	args := m.Called()
	return args.Get(0).([]domain.Task)

}
func (m *TaskRepoMock) GetUserTasks(user_id string) []domain.Task {
	args := m.Called(user_id)
	return args.Get(0).([]domain.Task)
}
func (m *TaskRepoMock) GetTask(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *TaskRepoMock) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *TaskRepoMock) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

type TaskUsecaseSuite struct {
	suite.Suite
	usecase    usecases.TaksUsecase
	repository *TaskRepoMock
}

func (s *TaskUsecaseSuite) SetupSuite() {
	mockRepo := new(TaskRepoMock)
	s.usecase = usecases.NewTaskUsecase(mockRepo)
	s.repository = mockRepo
	s.repository.On("CreateTask", mock.Anything).Return(domain.Task{}, nil)

}
func (s *TaskUsecaseSuite) TestCreateTask() {
	newTask := domain.Task{
		UserId:      primitive.NewObjectID(),
		Title:       "Task Title",
		Description: "Task description",
		DueDate:     time.Time{},
		Status:      "pending",
	}

	_, err := s.usecase.CreateTask(newTask)
	require.NoError(s.T(), err, "Failed to create task")

	// require.Equal(s.T(), newTask.Title, created.Title)
}

func TestTaksUsecase(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
