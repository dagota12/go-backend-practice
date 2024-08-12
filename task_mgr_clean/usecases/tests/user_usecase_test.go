package tests

import (
	"goPractice/task_manager/domain"
	"goPractice/task_manager/infrastructure"
	"goPractice/task_manager/usecases"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	userUsecase := usecases.NewUserUsecase(userRepo)
	// userUsecase.
	token, err := userUsecase.Login(data)
	require.NoError(t, err, "failed to login")
	// t.Log(token)
	require.Greater(t, len(token), 0)
	userRepo.AssertExpectations(t)
}
func TestGetUsers(t *testing.T) {
	userRepo := new(UserRepoMock)
	userRepo.On("GetUsers").Return([]domain.UserOut{
		{
			ID:       primitive.NewObjectID(),
			Username: "brad",
			Role:     "user",
		},
	}, nil)
	userUsecase := usecases.NewUserUsecase(userRepo)
	users := userUsecase.GetUsers()
	require.Equal(t, 1, len(users))
	userRepo.AssertExpectations(t)
}

func TestPromoteUser(t *testing.T) {
	userRepo := new(UserRepoMock)
	userRepo.On("PromoteUser", mock.Anything).Return(nil)
	userUsecase := usecases.NewUserUsecase(userRepo)
	err := userUsecase.PromoteUser("66b3876f2c78431256a57b41")
	require.NoError(t, err)
	userRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	newUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "user2",
		Password: "test",
		Role:     "user",
	}
	userRepo := new(UserRepoMock)
	userRepo.On("FilterUser", newUser.Username).Return(domain.User{}, nil)
	userRepo.On("GetUsers").Return([]domain.UserOut{}, nil)
	//when create user is called with any input just retunrn anything
	userRepo.On("CreateUser", mock.Anything).Return(newUser, nil)

	userUsecase := usecases.NewUserUsecase(userRepo)
	_, err := userUsecase.CreateUser(newUser)
	require.NoError(t, err)
	userRepo.AssertExpectations(t)
}
