using LibraryService.Messaging;
using LibraryService.Messaging.Messages;
using LibraryService.Models;
using LibraryService.Repository.Contracts;
using LibraryService.Services.Base;
using System.Collections.Generic;

namespace LibraryService.Services.Impl
{
    public class PublicationService : IPublicationService
    {
        private readonly IPublicationRepository _publicationRepository;
        private readonly IKafkaConsumer<PaperPublished> _kafkaConsumer;

        public PublicationService(IPublicationRepository publicationRepository)
        {
            _publicationRepository = publicationRepository;
            _kafkaConsumer = new KafkaConsumer<PaperPublished>("PUBLISH_PAPER");
            _kafkaConsumer.AddListener(PublishListener);
        }

        public void PublishListener(PaperPublished paperPublished)
        {
            Publication publication = new()
            {
                Author = paperPublished.Author, 
                Title = paperPublished.Title,
                PaperId = paperPublished.PaperId,
            };
            _publicationRepository.InsertOne(publication);
        }

        public IEnumerable<Publication> Search()
        {
            return _publicationRepository.AsQueryable();
        }
    }
}
