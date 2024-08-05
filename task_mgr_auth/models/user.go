package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username,omitempty" json:"username" binding:"required"`
	Password string             `bson:"password,omitempty" json:"password" binding:"required"`
	Role     string             `bson:"role,omitempty" json:"role" binding:"required"`
}
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
