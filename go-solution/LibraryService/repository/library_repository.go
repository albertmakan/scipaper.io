package repository

import (
	"context"

	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibraryRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewLibraryRepository(collection *mongo.Collection, ctx context.Context) *LibraryRepository {
	return &LibraryRepository{collection, ctx}
}

func (libraryRepository *LibraryRepository) Create(publication *models.Publication) error {
	_, err := libraryRepository.collection.InsertOne(libraryRepository.ctx, publication)
	return err
}

func (libraryRepository *LibraryRepository) Search(query string) *[]models.Publication {
	cursor, _ := libraryRepository.collection.Find(libraryRepository.ctx, bson.M{})
	var publications []models.Publication
	_ = cursor.All(libraryRepository.ctx, &publications)
	return &publications
}