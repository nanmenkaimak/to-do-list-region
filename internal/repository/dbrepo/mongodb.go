package dbrepo

import (
	"context"
	"errors"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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
	task.CreatedAt = time.Now()
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
	if similarTasks.ID == primitive.NilObjectID {
		return errors.New("there is not a single task with this id")
	}
	if similarTasks.Title == updatedTask.Title && similarTasks.ActiveAt == updatedTask.ActiveAt {
		return errors.New("there are already similar task")
	}
	_, err = m.DB.UpdateOne(context.TODO(), filter, update)
	return err
}

func (m *postgresDBRepo) DeleteTask(id primitive.ObjectID) error {
	_, err := m.DB.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (m *postgresDBRepo) UpdateStatus(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.M{"isDone": true}}}

	_, err := m.DB.UpdateOne(context.TODO(), filter, update)
	return err
}

func (m *postgresDBRepo) GetAllTask(status bool) ([]models.Task, error) {
	var allTasks []models.Task
	filter := bson.D{{"isDone", status},
		{"activeAt", bson.M{"$gte": time.Now().Format("2006-01-02")}}}

	opts := options.Find().SetSort(bson.D{{"createdAt", 1}})

	res, err := m.DB.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	err = res.All(context.TODO(), &allTasks)
	if err != nil {
		return nil, err
	}
	return allTasks, err
}
