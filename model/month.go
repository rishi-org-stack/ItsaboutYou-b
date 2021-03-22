package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Month struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Saving float64`bson:"saving"`
	Earning float64`bson:"earning"`
	TotlTasksComplted int `bson:"total_tasks_completed"`
}
//TODO:look on the diagram and make changes
