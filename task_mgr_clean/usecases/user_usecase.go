package usecases

import (
	"goPractice/task_manager/domain"
	"goPractice/task_manager/repositories"
)

type UserUsecase interface {
	CreateUser(user domain.User) (domain.User, error)
	FilterUser(username string) (domain.User, error)
	GetUsers() []domain.User
	PromoteUser(string) (bool, error)
}
type userUsecase struct {
	userRepository repositories.UserRepository
}

// CreateUser implements UserUsecase.
func (u *userUsecase) CreateUser(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// FilterUser implements UserUsecase.
func (u *userUsecase) FilterUser(username string) (domain.User, error) {
	panic("unimplemented")
}

// GetUsers implements UserUsecase.
func (u *userUsecase) GetUsers() []domain.User {
	panic("unimplemented")
}

// PromoteUser implements UserUsecase.
func (u *userUsecase) PromoteUser(string) (bool, error) {
	panic("unimplemented")
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
