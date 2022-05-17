package controllers

import (
	"net/http"

	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/helpers"
	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/services"
)

type LibraryController struct {
	libraryService *services.LibraryService
}

func NewLibraryController(libraryService *services.LibraryService) *LibraryController {
	return &LibraryController{libraryService}
}

func (lc *LibraryController) Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == http.MethodOptions {w.WriteHeader(http.StatusOK); return}
		helpers.JSONResponse(w, http.StatusOK, lc.libraryService.Search(""))
	}
}

func (lc *LibraryController) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		helpers.JSONResponse(w, http.StatusOK, "Hello from LibraryService")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}