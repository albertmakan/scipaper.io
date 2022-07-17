using SciPaperService.Models;
using System.Collections.Generic;

namespace SciPaperService.Services.Base
{
    public interface IPaperService
    {
        Paper CreatePaper(Paper paper);
        Paper UpdatePaper(Paper paperUpdate);
        Paper ReadPaper(string id);
        void DeletePaper(string id);
        IEnumerable<Paper> GetAllByAuthor(string authorId);
        void Publish(string paperId, string authorId);
    }
}
