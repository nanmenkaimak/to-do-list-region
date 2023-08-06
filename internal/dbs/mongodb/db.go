package mongodb

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// New подключается в базу данных MongoDB
func New() (*mongo.Collection, error) {
	// для чтение .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// создает клиента
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	// подключается в базу данных MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// проверяет связь
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// создает коллекцию
	collection := client.Database("to-do-list-region").Collection("task")

	return collection, nil
}
