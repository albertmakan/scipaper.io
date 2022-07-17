using SciPaperService.Exceptions;
using SciPaperService.Messaging;
using SciPaperService.Messaging.Messages;
using SciPaperService.Models;
using SciPaperService.Repository.Contracts;
using SciPaperService.Services.Base;
using SciPaperService.Services.Dependencies;
using System.Collections.Generic;

namespace SciPaperService.Services.Impl
{
    public class PaperService : IPaperService
    {
        private readonly IPaperRepository _paperRepository;
        private readonly IUserClient _userClient;
        private readonly IKafkaProducer _kafkaProducer;

        public PaperService(IPaperRepository paperRepository, IUserClient userClient, IKafkaProducer kafkaProducer)
        {
            _paperRepository = paperRepository;
            _userClient = userClient;
            _kafkaProducer = kafkaProducer;
        }

        public Paper CreatePaper(Paper paper)
        {
            paper.Author = _userClient.GetName(paper.AuthorId);
            return _paperRepository.InsertOne(paper);
        }

        public void DeletePaper(string id)
        {
            ReadPaper(id);
            _paperRepository.DeleteById(id);
        }

        public IEnumerable<Paper> GetAllByAuthor(string authorId)
        {
            return _paperRepository.FindAllByAuthorId(authorId);
        }

        public void Publish(string paperId, string authorId)
        {
            var paper = ReadPaper(paperId);
            if (paper.AuthorId != authorId) throw new ForbiddenException("author cannot publish someone elses paper");
            PaperPublished message = new()
            {
                PaperId = paperId,
                Author = paper.Author,
                Title = paper.Title,
            };
            _kafkaProducer.Send("PUBLISH_PAPER", message);
        }

        public Paper ReadPaper(string id)
        {
            var paper = _paperRepository.FindById(id);
            if (paper == null) throw new NotFoundException($"{id} paper not found");
            return paper;
        }

        public Paper UpdatePaper(Paper paperUpdate)
        {
            paperUpdate.Author = _userClient.GetName(paperUpdate.AuthorId);
            var paper = ReadPaper(paperUpdate.Id);
            if (paper.AuthorId != paperUpdate.AuthorId) throw new ForbiddenException("author cannot publish someone elses paper");
            return _paperRepository.ReplaceOne(paperUpdate);
        }
    }
}
