package dbrepo

import (
	"context"
	"errors"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *postgresDBRepo) CreateTask(task models.Task) error {
	var similarTasks models.Task
	filter := bson.D{{"title", task.Title}, {"activeAt", task.ActiveAt}}
	err := m.DB.FindOne(context.TODO(), filter).Decode(&similarTasks)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	if similarTasks.Title == task.Title && similarTasks.ActiveAt == task.ActiveAt {
		return errors.New("there are already similar task")
	}
	_, err = m.DB.InsertOne(context.TODO(), task)
	return err
}

func (m *postgresDBRepo) UpdateTask(updatedTask models.Task) error {
	var similarTasks models.Task
	filter := bson.D{{"_id", updatedTask.ID}}
	findOneFilter := bson.D{{"title", updatedTask.Title}, {"activeAt", updatedTask.ActiveAt}}
	update := bson.D{{"$set", findOneFilter}}

	err := m.DB.FindOne(context.TODO(), findOneFilter).Decode(&similarTasks)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	if similarTasks.Title == updatedTask.Title && similarTasks.ActiveAt == updatedTask.ActiveAt {
		return errors.New("there are already similar task")
	}
	_, err = m.DB.UpdateOne(context.TODO(), filter, update)
	return err
}
