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
		helpers.JSONResponse(w, http.StatusOK, lc.libraryService.Search(""))
	}
}

func (lc *LibraryController) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.JSONResponse(w, http.StatusOK, "Hello from LibraryService")
	}
}