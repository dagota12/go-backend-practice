package usecases

import (
	"errors"
	"goPractice/task_manager/domain"
	"goPractice/task_manager/infrastructure"
	"goPractice/task_manager/repositories"
)

type UserUsecase interface {
	CreateUser(user domain.User) (domain.User, error)
	FilterUser(username string) (domain.User, error)
	GetUsers() []domain.UserOut
	PromoteUser(string) error
	Login(data domain.LoginForm) (string, error)
}
type userUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
func (u userUsecase) Login(data domain.LoginForm) (string, error) {

	user, err := u.userRepository.FilterUser(data.Username)
	if err != nil {
		return "", err
	}
	//check password
	if valid := infrastructure.CheckPasswordHash(data.Password, user.Password); !valid {
		return "", errors.New("password doesn't match")
	}
	//generate jwt token
	token, err := infrastructure.GenerateToken(user.ID.Hex(), user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

// CreateUser implements UserUsecase.
func (u *userUsecase) CreateUser(user domain.User) (domain.User, error) {
	//check if user aready exist
	existingUser, _ := u.userRepository.FilterUser(user.Username)

	if existingUser.Username != "" {
		return domain.User{}, errors.New("user already exist")
	}
	//hash password using bcrypt
	hashedPwd, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, errors.New("error while hashing the password")

	}
	user.Password = hashedPwd
	//create the user in the database
	//check if database is empty or if there are no users this user role becomes admin
	users, _ := u.userRepository.GetUsers()
	if len(users) == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

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
