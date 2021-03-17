package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Meetup struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	With    Contact            `bson:"with"`
	Where   string             `bson:"where"`
	Subject []byte             `bson:"subject"`
}
