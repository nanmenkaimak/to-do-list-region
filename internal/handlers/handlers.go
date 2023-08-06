package handlers

import (
	"github.com/nanmenkaimak/to-do-list-region/internal/repository"
	"github.com/nanmenkaimak/to-do-list-region/internal/repository/dbrepo"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repo используется ли репозиторий обработчиками
var Repo *Repository

// Repository является типом репозиторий
type Repository struct {
	DB repository.DatabaseRepo
}

// NewRepo создает новый репозиторий
func NewRepo(db *mongo.Collection) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db),
	}
}

// NewHandlers устанавливает репозиторий для обработчиков
func NewHandlers(r *Repository) {
	Repo = r
}
