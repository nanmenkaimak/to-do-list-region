package main

import (
	"github.com/nanmenkaimak/to-do-list-region/internal/dbs/mongodb"
	"github.com/nanmenkaimak/to-do-list-region/internal/handlers"
)

func main() {
	db, err := mongodb.New()
	if err != nil {
		panic(err)
	}

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	routes()
}
