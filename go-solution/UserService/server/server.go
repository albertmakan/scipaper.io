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
	router.HandleFunc("/aaa", uc.Create())
}

func (server *Server) Start() {
	http.ListenAndServe(":8000", server.Router)
}