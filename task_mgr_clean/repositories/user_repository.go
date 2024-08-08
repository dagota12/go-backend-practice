package repositories

import (
	"goPractice/task_manager/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FilterUser(username string) (domain.User, error)
	GetUsers() []domain.User
	PromoteUser(string) (bool, error)
}
type userRepository struct {
	db    *mongo.Database
	users *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db, users: db.Collection("users")}
}

// CreateUser implements UserRepository.
func (u *userRepository) CreateUser(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// FilterUser implements UserRepository.
func (u *userRepository) FilterUser(username string) (domain.User, error) {
	panic("unimplemented")
}

// GetUsers implements UserRepository.
func (u *userRepository) GetUsers() []domain.User {
	panic("unimplemented")
}

// PromoteUser implements UserRepository.
func (u *userRepository) PromoteUser(string) (bool, error) {
	panic("unimplemented")
}
