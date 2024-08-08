package usecases

import (
	"goPractice/task_manager/domain"
	"goPractice/task_manager/repositories"
)

type UserUsecase interface {
	CreateUser(user domain.User) (domain.User, error)
	FilterUser(username string) (domain.User, error)
	GetUsers() []domain.UserOut
	PromoteUser(string) error
}
type userUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

// CreateUser implements UserUsecase.
func (u *userUsecase) CreateUser(user domain.User) (domain.User, error) {
	newUser, err := u.userRepository.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil

}

// FilterUser implements UserUsecase.
func (u *userUsecase) FilterUser(username string) (domain.User, error) {
	user, err := u.userRepository.FilterUser(username)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// GetUsers implements UserUsecase.
func (u *userUsecase) GetUsers() []domain.UserOut {
	users, _ := u.userRepository.GetUsers()
	return users
}

// PromoteUser implements UserUsecase.
func (u *userUsecase) PromoteUser(id string) error {
	err := u.userRepository.PromoteUser(id)
	if err != nil {
		return err
	}
	return nil
}
