package tests

import (
	"goPractice/task_manager/domain"

	"github.com/stretchr/testify/mock"
)

type TaskRepoMock struct {
	mock.Mock
}

func (m *TaskRepoMock) CreateTask(task domain.Task) (domain.Task, error) {
	args := m.Called(task)

	return args.Get(0).(domain.Task), args.Error(1)
}
