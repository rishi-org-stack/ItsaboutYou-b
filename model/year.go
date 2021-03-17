package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Year struct{
	ID primitive.ObjectID	`bson:"_id,omitempty"`
	TotalSavings float64`bson:"total_saving"`
	TotalEarning float64`bson:"total_earning"`
	NoofTasksCompleted int64`bson:"no_of_tasks_completed"`

}