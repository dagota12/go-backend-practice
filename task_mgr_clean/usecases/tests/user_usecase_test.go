package tests

import (
	"goPractice/task_manager/domain"
	"goPractice/task_manager/infrastructure"
	"goPractice/task_manager/usecases"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) CreateUser(user domain.User) (domain.User, error) {
	args := u.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}
func (u *UserRepoMock) GetUsers() ([]domain.UserOut, error) {
	args := u.Called()
	return args.Get(0).([]domain.UserOut), args.Error(1)
}
func (u *UserRepoMock) FilterUser(username string) (domain.User, error) {
	args := u.Called(username)
	return args.Get(0).(domain.User), args.Error(1)
}
func (u *UserRepoMock) PromoteUser(id string) error {
	args := u.Called(id)
	return args.Error(0)
}

// func (u *UserRepoMock) Login(data domain.LoginForm) {

// }

func TestUserLogin(t *testing.T) {
	userRepo := new(UserRepoMock)
	data := domain.LoginForm{
		Username: "brad",
		Password: "brad",
	}
	hashedPwd, _ := infrastructure.HashPassword(data.Password)
	user := domain.User{
		Username: "brad",
		Password: hashedPwd,
		Role:     "user",
	}
	userRepo.On("FilterUser", data.Username).Return(user, nil)

	// userRepo.On("Login", data).Return("token", nil)

	userUsecase := usecases.NewUserUsecase(userRepo)
	// userUsecase.
	token, err := userUsecase.Login(data)
	require.NoError(t, err, "failed to login")
	t.Log(token)
	require.Greater(t, len(token), 0)
	userRepo.AssertExpectations(t)
}
