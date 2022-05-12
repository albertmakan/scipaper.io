package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/controllers"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/repository"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/server"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("go.env")
  if err != nil {
    panic("Error loading .env file")
  }

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("CONNECTION_STRING")))
	if err != nil {panic(err)}
	defer client.Disconnect(ctx)
	database := client.Database("scipaper-io")
	userService := services.NewUserService(
		repository.NewUserRepository(database.Collection("user"), ctx),
	)
	controller := controllers.NewUserController(userService)


	api := services.RPC{UserService:  userService}
	err = rpc.Register(&api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	go http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}


	s := server.New()
	s.AddHandlers(controller)
	s.Start()
}