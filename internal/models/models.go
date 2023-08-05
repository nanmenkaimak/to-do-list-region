package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID       primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Title    string             `json:"title" bson:"title"`
	ActiveAt string             `json:"activeAt" bson:"activeAt"`
	IsDone   bool               `json:"isDone" bson:"isDone"`
}
