package server

import (
	"log"
	"net/http"

	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/controllers"
	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/services"
)

type Server struct {
	router *http.ServeMux
	sciPaperController *controllers.SciPaperController
}

func New(sciPaperService *services.SciPaperService) *Server {
	server := &Server{
		http.NewServeMux(),
		controllers.NewSciPaperController(sciPaperService),
	}
	server.addHandlers()
	return server
}

func (server *Server) addHandlers() {
	router := server.router
	router.HandleFunc("/create-update", server.sciPaperController.CreateOrUpdate())
	router.HandleFunc("/my-papers", server.sciPaperController.GetAllByAuthor())
	router.HandleFunc("/hello", server.sciPaperController.Hello())
}

func (server *Server) Start() {
	log.Printf("serving http on port %d", 8001)
	http.ListenAndServe(":8001", server.router)
}