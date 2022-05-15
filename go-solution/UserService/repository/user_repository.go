package repository

import (
	"context"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	ctx context.Context
}

func NewUserRepository(collection *mongo.Collection, ctx context.Context) *UserRepository {
	return &UserRepository{collection, ctx}
}

func (userRepository *UserRepository) Create(user *models.User) error {
	_, err := userRepository.collection.InsertOne(userRepository.ctx, user)
	return err
}

func (userRepository *UserRepository) GetAll() *[]models.User {
	cursor, _ := userRepository.collection.Find(userRepository.ctx, bson.M{})
	var users []models.User
	_ = cursor.All(userRepository.ctx, &users)
	return &users
}

func (userRepository *UserRepository) FindByUsername(username string) *models.User {
	var user models.User
	err := userRepository.collection.FindOne(userRepository.ctx, bson.M{"username":username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	return &user
}