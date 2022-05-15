package controllers

import (
	"net/http"
	"strings"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/helpers"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/models"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var newUser models.User
		helpers.ReadJSONBody(r, &newUser)
		err := uc.userService.Register(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		helpers.JSONResponse(w, http.StatusCreated, nil)
	}
}

func (uc *UserController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.JSONResponse(w, http.StatusOK, uc.userService.GetAll())
	}
}

func (uc *UserController) FindByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := strings.TrimPrefix(r.URL.Path, "/users/")
		user := uc.userService.FindByUsername(username)
		if user == nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		helpers.JSONResponse(w, http.StatusOK, user)
	}
}

func (uc *UserController) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var login struct {Username, Password string}
		helpers.ReadJSONBody(r, &login)

		auth, err := uc.userService.Authenticate(login.Username, login.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		helpers.JSONResponse(w, http.StatusOK, auth)
	}
}

func (uc *UserController) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.JSONResponse(w, http.StatusOK, "Hello from UserService")
	}
}