package server

import (
	"log"
	"net/http"

	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/controllers"
	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/services"
)

type Server struct {
	router             *http.ServeMux
	libraryController *controllers.LibraryController
}

func New(libraryService *services.LibraryService) *Server {
	server := &Server{
		http.NewServeMux(),
		controllers.NewLibraryController(libraryService),
	}
	server.addHandlers()
	return server
}

func (server *Server) addHandlers() {
	router := server.router
	router.HandleFunc("/search", server.libraryController.Search())
	router.HandleFunc("/hello", server.libraryController.Hello())
}

func (server *Server) Start() {
	log.Printf("serving http on port %d", 8002)
	http.ListenAndServe(":8002", server.router)
}