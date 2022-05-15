package services

import (
	"encoding/json"

	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/messaging"
	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/models"
	"github.com/albertmakan/scipaper.io/go-solution/LibraryService/repository"
)

type LibraryService struct {
	libraryRepository *repository.LibraryRepository
	receiver *messaging.AMQPReceiver
}

func NewLibraryService(libraryRepository *repository.LibraryRepository) *LibraryService {
	return &LibraryService{libraryRepository: libraryRepository}
}

func (libraryService *LibraryService) InitializeReceiver() {
	libraryService.receiver = &messaging.AMQPReceiver{}
	libraryService.receiver.Initialize("PUBLISH_PAPER")
}

func (libraryService *LibraryService) DeinitializeReceiver() {
	libraryService.receiver.Deinitialize()
}

func (libraryService *LibraryService) StartConsuming() {
	libraryService.receiver.Consume(func(b []byte) {
		var publication models.Publication
		json.Unmarshal(b, &publication)
		libraryService.libraryRepository.Create(&publication)
	})
}

func (libraryService *LibraryService) Search(query string) *[]models.Publication {
	return libraryService.libraryRepository.Search(query)
}