package data

import (
	"context"
	"goPractice/task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	FilterUser(filter bson.M) (models.User, error)
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
func (us *userService) FilterUser(filter bson.M) (models.User, error) {
	var user models.User
	res := us.users.FindOne(context.TODO(), filter)

	err := res.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}
