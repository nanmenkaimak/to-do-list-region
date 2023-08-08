package main

import (
	_ "github.com/nanmenkaimak/to-do-list-region/docs"
	"github.com/nanmenkaimak/to-do-list-region/internal/dbs/mongodb"
	"github.com/nanmenkaimak/to-do-list-region/internal/handlers"
)

// @title ToDo List Region
// @version 1.0
// @description to do list for region

// @host localhost:8080
// @BasePath /

func main() {
	// подключает в базу данных
	db, err := mongodb.New()
	if err != nil {
		panic(err)
	}

	// создаем новую репозиторию
	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	// вызываем функцию routes
	routes()
}
