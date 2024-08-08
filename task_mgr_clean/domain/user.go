package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username,omitempty" json:"username" binding:"required"`
	Password string             `bson:"password,omitempty" json:"password" binding:"required"`
	Role     string             `bson:"role,omitempty" json:"role"`
}

// UserOut defines a structure for the user output in the http response
type UserOut struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username,omitempty" json:"username" binding:"required"`
	Role     string             `bson:"role,omitempty" json:"role"`
}
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
