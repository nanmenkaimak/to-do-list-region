package repository

import (
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatabaseRepo interface {
	CreateTask(task models.Task) error
	UpdateTask(updatedTask models.Task) error
	DeleteTask(id primitive.ObjectID) error
	UpdateStatus(id primitive.ObjectID) error
	GetAllTask(status bool) ([]models.Task, error)
}
