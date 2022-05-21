package controllers

import (
	"net/http"

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
		enableCors(&w)
		if r.Method == http.MethodOptions {w.WriteHeader(http.StatusOK); return}
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

func (uc *UserController) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == http.MethodOptions {w.WriteHeader(http.StatusOK); return}
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
		helpers.JSONResponse(w, http.StatusOK, auth.Jwt)
	}
}

func (uc *UserController) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		helpers.JSONResponse(w, http.StatusOK, "Hello from UserService")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}