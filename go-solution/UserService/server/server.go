package server

import (
	"net/http"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/controllers"
)

type Server struct {
	Router *http.ServeMux
}

func New() *Server {
	server := &Server{http.NewServeMux()}
	// server.addHandlers()
	return server
}

func (server *Server) AddHandlers(uc *controllers.UserController) {
	router := server.Router
	router.HandleFunc("/register", uc.Register())
	router.HandleFunc("/all-users", uc.GetAll())
	router.HandleFunc("/users/", uc.FindByUsername())
	router.HandleFunc("/auth", uc.Authenticate())
}

func (server *Server) Start() {
	http.ListenAndServe(":8000", server.Router)
}