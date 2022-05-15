package services

import (
	"encoding/json"
	"fmt"

	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/messaging"
	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/models"
	"github.com/albertmakan/scipaper.io/go-solution/SciPaperService/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SciPaperService struct {
	sciPaperRepository *repository.SciPaperRepository
	sender *messaging.AMQPSender
}

func NewSciPaperService(sciPaperRepository *repository.SciPaperRepository) *SciPaperService {
	return &SciPaperService{sciPaperRepository: sciPaperRepository}
}

func (sciPaperService *SciPaperService) InitializeSender() {
	sciPaperService.sender = &messaging.AMQPSender{}
	sciPaperService.sender.Initialize("PUBLISH_PAPER")
}

func (sciPaperService *SciPaperService) DeinitializeSender() {
	sciPaperService.sender.Deinitialize()
}

func (sciPaperService *SciPaperService) Create(paper *models.Paper) error {
	return sciPaperService.sciPaperRepository.Create(paper)
}

func (sciPaperService *SciPaperService) GetAllByAuthor(author string) *[]models.Paper {
	return sciPaperService.sciPaperRepository.GetAllByAuthor(author)
}

func (sciPaperService *SciPaperService) Update(paper *models.Paper) error {
	return sciPaperService.sciPaperRepository.Update(paper)
}

func (sciPaperService *SciPaperService) FindByID(id primitive.ObjectID) *models.Paper {
	return sciPaperService.sciPaperRepository.FindByID(id)
}

func (sciPaperService *SciPaperService) Publish(paperID primitive.ObjectID) error {
	paper := sciPaperService.FindByID(paperID)
	if paper == nil { return fmt.Errorf("paper not found") }
	info := struct {PaperID primitive.ObjectID; Author string; Title string
		} {paper.ID, paper.Author, paper.Title}
	body, _ := json.Marshal(info)
	sciPaperService.sender.Send(body)
	return nil
}
