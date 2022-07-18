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

        public PublicationController(IPublicationService publicationService)
        {
            _publicationService = publicationService;
        }

        [HttpGet("search")]
        public IEnumerable<Publication> Search()
        {
            return _publicationService.Search();
        }
    }
}
