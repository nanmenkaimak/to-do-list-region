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

// CreateTask создает новые задачи
func (m *mongoDBRepo) CreateTask(task models.Task) error {
	// генерирует ID
	task.ID = primitive.NewObjectID()

	// проверяет из таблицы что нету такого же задачи
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

	// записывает в базу
	_, err = m.DB.InsertOne(context.TODO(), task)
	return err
}

// UpdateTask обновляет уже существующий задачи
func (m *mongoDBRepo) UpdateTask(updatedTask models.Task) error {
	var similarTasks models.Task
	filter := bson.D{{"_id", updatedTask.ID}}
	findOneFilter := bson.D{{"title", updatedTask.Title}, {"activeAt", updatedTask.ActiveAt}}
	update := bson.D{{"$set", findOneFilter}}

	// проверяет из таблицы что нету такого же задачи
	err := m.DB.FindOne(context.TODO(), findOneFilter).Decode(&similarTasks)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	// проверяет что из таблицы что нету такого же задачи
	if similarTasks.Title == updatedTask.Title && similarTasks.ActiveAt == updatedTask.ActiveAt {
		return errors.New("there are already similar task")
	}
	// обновляет задачу
	_, err = m.DB.UpdateOne(context.TODO(), filter, update)
	return err
}

// DeleteTask удаляет задачу
func (m *mongoDBRepo) DeleteTask(id primitive.ObjectID) error {
	// проверяет из таблицы что есть такое задача
	filter := bson.D{{"_id", id}}
	res := m.DB.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return mongo.ErrNoDocuments
		}
		return res.Err()
	}
	// удаляет из базы
	_, err := m.DB.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// UpdateStatus помечает задачу выполненной
func (m *mongoDBRepo) UpdateStatus(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.M{"isDone": true}}}

	// проверяет из таблицы что есть такое задача
	res := m.DB.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return mongo.ErrNoDocuments
		}
		return res.Err()
	}

	// обновляет статус
	_, err := m.DB.UpdateOne(context.TODO(), filter, update)
	return err
}

// GetAllTasks показывает список задач по статусу
func (m *mongoDBRepo) GetAllTasks(status bool) ([]models.Task, error) {
	var allTasks []models.Task
	filter := bson.D{{"isDone", status},
		{"activeAt", bson.M{"$gte": time.Now().Format("2006-01-02")}}}

	// сортировка по дате создания
	opts := options.Find().SetSort(bson.D{{"createdAt", 1}})

	// находит задачи
	res, err := m.DB.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	// записывает все в allTasks
	err = res.All(context.TODO(), &allTasks)
	if err != nil {
		return nil, err
	}
	return allTasks, err
}
