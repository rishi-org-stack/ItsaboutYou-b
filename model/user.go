package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Age      int    `bson:"age"`
	Gender   string `bson:"gender"`
	Unique   string `bson:"unique"`
}
