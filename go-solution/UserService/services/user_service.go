package services

import (
	"fmt"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/dto"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/helpers"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/models"
	"github.com/albertmakan/scipaper.io/go-solution/UserService/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService *UserService) Register(user* models.User) error {
	if userService.FindByUsername(user.Username) != nil {
		return fmt.Errorf("user with username %v already exists", user.Username)
	}
	user.Password, _ = helpers.HashPassword(user.Password)
	userService.userRepository.Create(user)
	return nil
}

func (userService *UserService) GetAll() *[]models.User {
	return userService.userRepository.GetAll()
}

func (userService *UserService) FindByUsername(username string) *models.User {
	return userService.userRepository.FindByUsername(username)
}

func (userService *UserService) Authenticate(username, password string) (*dto.AuthenticatedUser, error) {
	user := userService.FindByUsername(username)
	if user == nil {
		return nil, fmt.Errorf("invalid username or password")
	}
	if !helpers.CheckPasswordHash(password, user.Password) {
		return nil, fmt.Errorf("invalid username or password")
	}
	claims := helpers.Claims{
		Name: username,
	}
	jwt, _ := helpers.GetJwtToken(claims)
	return &dto.AuthenticatedUser{Jwt:jwt}, nil
}