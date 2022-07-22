using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using SciPaperService.DTO;
using SciPaperService.Filters;
using SciPaperService.Models;
using SciPaperService.Services.Base;
using System.Collections.Generic;

namespace SciPaperService.Controllers
{
    [ApiController]
    [Produces("application/json")]
    [ValidateModel]
    public class PaperController : ControllerBase
    {
        private readonly IPaperService _paperService;

        public PaperController(IPaperService paperService)
        {
            _paperService = paperService;
        }

        [HttpPost("paper")]
        [Authorize]
        public Paper CreatePaper(Paper paper)
        {
            paper.AuthorId = HttpContext.User.Identity.Name;
            return _paperService.CreatePaper(paper);
        }

        [HttpPost("publish")]
        [Authorize]
        public void PublishPaper(PublishRequest request)
        {
            _paperService.Publish(request.PaperId, HttpContext.User.Identity.Name);
        }

        [HttpGet("paper/{id}")]
        public Paper ReadPaper(string id)
        {
            return _paperService.ReadPaper(id);
        }

        [HttpPut("paper/{id}")]
        [Authorize]
        public Paper UpdatePaper(string id, Paper paper)
        {
            return _paperService.UpdatePaper(paper);
        }

        [HttpDelete("paper/{id}")]
        [Authorize]
        public void DeletePaper(string id)
        {
            _paperService.DeletePaper(id);
        }

        [HttpGet("my-papers")]
        [Authorize]
        public IEnumerable<Paper> GetMyPapers()
        {
            return _paperService.GetAllByAuthor(HttpContext.User.Identity.Name);
        }
    }
}
