using LibraryService.Messaging;
using LibraryService.Messaging.Messages;
using LibraryService.Models;
using LibraryService.Services.Base;
using Microsoft.AspNetCore.Mvc;
using System.Collections.Generic;


namespace LibraryService.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class PublicationController : ControllerBase
    {
        private readonly IPublicationService _publicationService;
        private readonly IKafkaConsumer _kafkaConsumer;

        public PublicationController(IPublicationService publicationService)
        {
            _publicationService = publicationService;
            _kafkaConsumer = new KafkaConsumer("PUBLISH_PAPER");
            _kafkaConsumer.AddListener<PaperPublished>(PublishListener);
        }

        [HttpGet("search")]
        public IEnumerable<Publication> Search()
        {
            return _publicationService.Search();
        }

        private void PublishListener(PaperPublished message)
        {
            _publicationService.PublishListener(message);
        }
    }
}
