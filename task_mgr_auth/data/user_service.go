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
	GetAllUsers() ([]models.UserOut, error)
	PromoteUser(string) (bool, error)
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

// gets all users from database
func (us *userService) GetAllUsers() ([]models.UserOut, error) {
	var users []models.UserOut
	cursor, err := us.users.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user models.UserOut
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// promote user
func (us *userService) PromoteUser(id string) (bool, error) {
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	filter := bson.M{"_id": oId}
	update := bson.D{{Key: "$set", Value: bson.M{"role": "admin"}}}
	res, err := us.users.UpdateOne(context.TODO(), filter, update)
	if err != nil || res.ModifiedCount < 1 {
		log.Println("promoting user", err)
		return false, err
	}
	return true, nil
}
