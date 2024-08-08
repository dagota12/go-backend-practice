package repositories

import (
	"context"
	"errors"
	"goPractice/task_manager/domain"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FilterUser(username string) (domain.User, error)
	GetUsers() ([]domain.UserOut, error)
	PromoteUser(id string) error
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
	res, err := u.users.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err.Error())
		return domain.User{}, errors.New("error while createting user")
	}
	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil

}

// FilterUser implements UserRepository.
func (u *userRepository) FilterUser(username string) (domain.User, error) {
	filter := bson.M{"username": username}
	var user domain.User
	err := u.users.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println(err.Error())
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

// GetUsers implements UserRepository.
func (u *userRepository) GetUsers() ([]domain.UserOut, error) {
	var users []domain.UserOut
	cursor, err := u.users.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user domain.UserOut
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// PromoteUser implements UserRepository.
func (u *userRepository) PromoteUser(id string) error {
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	filter := bson.M{"_id": oId}
	update := bson.D{{Key: "$set", Value: bson.M{"role": "admin"}}}
	res, err := u.users.UpdateOne(context.TODO(), filter, update)
	if err != nil || res.ModifiedCount < 1 {
		log.Println("error on [promiting]", err)
		return err
	}
	return nil
}
