package handlers

import (
	"github.com/nanmenkaimak/to-do-list-region/internal/repository"
	"github.com/nanmenkaimak/to-do-list-region/internal/repository/dbrepo"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repo is repository used by handlers
var Repo *Repository

// Repository is repository type
type Repository struct {
	DB repository.DatabaseRepo
}

// NewRepo creates new repository
func NewRepo(db *mongo.Collection) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db),
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
