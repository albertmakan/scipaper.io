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

func (sciPaperRepository *SciPaperRepository) Create(paper *models.Paper) (interface{}, error) {
	result, err := sciPaperRepository.collection.InsertOne(sciPaperRepository.ctx, paper)
	if err != nil {return nil, err}
	return result.InsertedID, err
}

func (sciPaperRepository *SciPaperRepository) GetAllByAuthorID(authorID string) *[]models.Paper {
	cursor, _ := sciPaperRepository.collection.Find(sciPaperRepository.ctx, bson.M{"authorId":authorID})
	var papers []models.Paper
	_ = cursor.All(sciPaperRepository.ctx, &papers)
	return &papers
}

func (sciPaperRepository *SciPaperRepository) Update(paper *models.Paper) (interface{}, error) {
	_, err := sciPaperRepository.collection.UpdateByID(sciPaperRepository.ctx, paper.ID, bson.M{"$set":paper})
	if err != nil {return nil, err}
	return paper.ID, err
}

func (sciPaperRepository *SciPaperRepository) FindByID(id primitive.ObjectID) *models.Paper {
	var paper models.Paper
	err := sciPaperRepository.collection.FindOne(sciPaperRepository.ctx, bson.M{"_id":id}).Decode(&paper)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	return &paper
}