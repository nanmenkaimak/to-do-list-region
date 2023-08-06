package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Task структура таблицы
type Task struct {
	ID        primitive.ObjectID `json:"-" bson:"_id, omitempty"`
	Title     string             `json:"title" bson:"title"`
	ActiveAt  string             `json:"activeAt" bson:"activeAt"`
	IsDone    bool               `json:"-" bson:"isDone"`
	CreatedAt time.Time          `json:"-" bson:"createdAt"`
}
