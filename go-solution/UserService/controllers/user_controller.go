package controllers

import (
	"net/http"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/models"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := models.User{
			Username:  "makanalbert",
			Password:  "string",
			Salt:  "string",
			FirstName :  "Albert",
			LastName:  "Makan",
			Email:  "makanalbert@gmail.com",
		}
		uc.userService.Create(&u);
		w.Write([]byte("CREATED"))
	}
}