package repository

import "github.com/nanmenkaimak/to-do-list-region/internal/models"

type DatabaseRepo interface {
	CreateTask(task models.Task) error
	UpdateTask(updatedTask models.Task) error
}
