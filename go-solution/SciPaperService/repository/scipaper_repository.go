package repository

import (
	"context"

	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SciPaperRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewSciPaperRepository(collection *mongo.Collection, ctx context.Context) *SciPaperRepository {
	return &SciPaperRepository{collection, ctx}
}

func (sciPaperRepository *SciPaperRepository) Create(paper *models.Paper) error {
	_, err := sciPaperRepository.collection.InsertOne(sciPaperRepository.ctx, paper)
	return err
}

func (sciPaperRepository *SciPaperRepository) GetAllByAuthor(author string) *[]models.Paper {
	cursor, _ := sciPaperRepository.collection.Find(sciPaperRepository.ctx, bson.M{"author":author})
	var papers []models.Paper
	_ = cursor.All(sciPaperRepository.ctx, &papers)
	return &papers
}

func (sciPaperRepository *SciPaperRepository) Update(paper *models.Paper) error {
	_, err := sciPaperRepository.collection.UpdateByID(sciPaperRepository.ctx, paper.ID, paper)
	return err
}

func (sciPaperRepository *SciPaperRepository) FindByID(id primitive.ObjectID) *models.Paper {
	var paper models.Paper
	err := sciPaperRepository.collection.FindOne(sciPaperRepository.ctx, bson.M{"_id":id}).Decode(&paper)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	return &paper
}