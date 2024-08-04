package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username,omitempty" json:"username"`
	Password string             `bson:"password,omitempty" json:"-"`
	Role     string             `bson:"role,omitempty" json:"role"`
}
