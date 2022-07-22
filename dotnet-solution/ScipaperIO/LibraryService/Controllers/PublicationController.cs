using LibraryService.Filters;
using LibraryService.Models;
using LibraryService.Services.Base;
using Microsoft.AspNetCore.Mvc;
using System.Collections.Generic;


namespace LibraryService.Controllers
{
    [ApiController]
    [Produces("application/json")]
    [ValidateModel]
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
