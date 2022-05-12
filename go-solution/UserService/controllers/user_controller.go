package controllers

import (
	"net/http"
	"strings"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/dto"
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
		helpers.JSONResponse(w, 201, nil)
	}
}

func (uc *UserController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.JSONResponse(w, 200, uc.userService.GetAll())
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
		helpers.JSONResponse(w, 200, user)
	}
}

func (uc *UserController) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var login dto.Login
		helpers.ReadJSONBody(r, &login)

		auth, err := uc.userService.Authenticate(login.Username, login.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		helpers.JSONResponse(w, 200, auth)
	}
}