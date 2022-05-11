package repository

import (
	"context"
	"fmt"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	ctx context.Context
}

func NewUserRepository(collection *mongo.Collection, ctx context.Context) *UserRepository {
	return &UserRepository{collection, ctx}
}

func (userRepository *UserRepository) Create(user *models.User) {
	insertResult, err := userRepository.collection.InsertOne(userRepository.ctx, user)
	if err != nil {
			panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}