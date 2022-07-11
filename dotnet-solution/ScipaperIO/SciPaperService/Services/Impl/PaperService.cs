using SciPaperService.Models;
using SciPaperService.Repository.Contracts;
using SciPaperService.Services.Base;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace SciPaperService.Services.Impl
{
    public class PaperService : IPaperService
    {
        private readonly IPaperRepository _paperRepository;
        private readonly IHttpClientFactory _httpClientFactory;

        public PaperService(IPaperRepository paperRepository, IHttpClientFactory httpClientFactory)
        {
            _paperRepository = paperRepository;
            _httpClientFactory = httpClientFactory;

        }

        public async Task<Paper> CreatePaperAsync(Paper paper)
        {
            var client = _httpClientFactory.CreateClient("user");

            string name = await client.GetStringAsync($"api/User/getname/{paper.AuthorId}");

            paper.Author = name.Trim(new[] { '"' });
            return _paperRepository.InsertOne(paper);
        }

        public void DeletePaper(string id)
        {
            throw new System.NotImplementedException();
        }

        public List<Paper> GetAllByAuthor(string authorId)
        {
            throw new System.NotImplementedException();
        }

        public void Publish(string paperId, string authorId)
        {
            throw new System.NotImplementedException();
        }

        public Paper ReadPaper(string id)
        {
            throw new System.NotImplementedException();
        }

        public Paper UpdatePaper(Paper paperUpdate)
        {
            throw new System.NotImplementedException();
        }
    }
}
