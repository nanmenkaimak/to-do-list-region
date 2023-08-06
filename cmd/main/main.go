package main

import (
	"github.com/nanmenkaimak/to-do-list-region/internal/dbs/mongodb"
	"github.com/nanmenkaimak/to-do-list-region/internal/handlers"
)

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
