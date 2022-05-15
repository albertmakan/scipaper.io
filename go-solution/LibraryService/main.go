package main

import (
	"context"
	"os"

	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/repository"
	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/server"
	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load("go.env")
  if err != nil {
    panic("Error loading .env file")
  }

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION")))
	if err != nil {panic(err)}
	defer client.Disconnect(ctx)
	database := client.Database("scipaper-io")
	libraryService := services.NewLibraryService(
		repository.NewLibraryRepository(database.Collection("library"), ctx),
	)
	libraryService.InitializeReceiver()
	libraryService.StartConsuming()
	defer libraryService.DeinitializeReceiver()
	server.New(libraryService).Start()
}