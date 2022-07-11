using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using SciPaperService.Filters;
using SciPaperService.Models;
using SciPaperService.Services.Base;
using System.Threading.Tasks;

namespace SciPaperService.Controllers
{
    [Route("api/[controller]")]
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

        [HttpPost("")]
        [Authorize]
        public async Task<Paper> CreatePaper(Paper paper)
        {
            paper.AuthorId = HttpContext.User.Identity.Name;
            return await _paperService.CreatePaperAsync(paper);
        }

    }
}
