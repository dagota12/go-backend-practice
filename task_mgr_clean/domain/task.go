package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Title       string             `bson:"title,omitempty" json:"title"`
	Description string             `bson:"description,omitempty" json:"description"`
	DueDate     time.Time          `bson:"due_date,omitempty" json:"due_date"`
	Status      string             `bson:"status,omitempty" json:"status"`
}
