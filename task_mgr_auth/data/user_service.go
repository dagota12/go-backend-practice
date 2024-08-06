package data

import (
	"context"
	"fmt"
	"goPractice/task_manager/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	FilterUser(username string) (models.User, error)
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
	res, err := us.users.InsertOne(context.TODO(), user)
	if err != nil {
		return models.User{}, fmt.Errorf("error to create user: %v", err.Error())
	}
	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil

}
func (us *userService) FilterUser(username string) (models.User, error) {
	filter := bson.M{"username": username}
	var user models.User
	err := us.users.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println(err.Error())
		return models.User{}, err
	}
	return user, nil

}
