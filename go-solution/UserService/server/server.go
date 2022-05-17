package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/controllers"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/services"
)

type Server struct {
	router *http.ServeMux
	userController *controllers.UserController
	rpc *services.RPC
}

func New(userService *services.UserService) *Server {
	server := &Server{
		http.NewServeMux(),
		controllers.NewUserController(userService),
		&services.RPC{UserService: userService},
	}
	server.addHandlers()
	return server
}

func (server *Server) addHandlers() {
	router := server.router
	router.HandleFunc("/register", server.userController.Register())
	router.HandleFunc("/auth", server.userController.Authenticate())
	router.HandleFunc("/hello", server.userController.Hello())
}

func (server *Server) Start() {
	rpc.Register(server.rpc)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	go http.Serve(listener, nil)
	log.Printf("serving http on port %d", 8000)
	http.ListenAndServe(":8000", server.router)
}