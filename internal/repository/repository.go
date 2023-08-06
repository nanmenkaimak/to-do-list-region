package repository

import (
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// функция для создания mocks для тестирования
//go:generate mockgen -source=repository.go -destination=mocks/mock.go

// DatabaseRepo интерфейс где хранится функций для работы с база данными
type DatabaseRepo interface {
	CreateTask(task models.Task) error
	UpdateTask(updatedTask models.Task) error
	DeleteTask(id primitive.ObjectID) error
	UpdateStatus(id primitive.ObjectID) error
	GetAllTasks(status bool) ([]models.Task, error)
}
