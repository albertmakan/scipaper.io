package services

import (
	"github.com/albertmakan/scipaper.io/go-solution/UserService/models"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService *UserService) Create(user* models.User) {
	userService.userRepository.Create(user)
}