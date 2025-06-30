package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title    string             `bson:"title" json:"title"`
	Company  string             `bson:"company" json:"company"`
	Location string             `bson:"location" json:"location"`
}
