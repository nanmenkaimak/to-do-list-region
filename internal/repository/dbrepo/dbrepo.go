package dbrepo

import (
	"github.com/nanmenkaimak/to-do-list-region/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type postgresDBRepo struct {
	DB *mongo.Collection
}

func NewPostgresRepo(conn *mongo.Collection) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}
