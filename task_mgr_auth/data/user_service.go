package data

import (
	"goPractice/task_manager/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
}
type userService struct {
	users *mongo.Collection
}

func NewUserService(collection *mongo.Collection) *userService {

	return &userService{
		users: collection,
	}
}
func (us *userService) CreateUser(user models.User) (models.User, error) {

	return models.User{}, nil

}
