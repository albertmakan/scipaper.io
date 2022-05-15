package main

import (
	"context"
	"os"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/repository"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/server"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/services"
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
	userService := services.NewUserService(
		repository.NewUserRepository(database.Collection("user"), ctx),
	)
	server.New(userService).Start()
}