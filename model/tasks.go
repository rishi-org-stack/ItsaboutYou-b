package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	DueDate string `bson:"duedate"`
	Subject []byte`bson:"subject"`
	Description []byte `bson:"description"`
	Priority int`bson:"priority"`
	Completed bool`bson:"completed"`
}
//TODO:look on the diagram and make changes