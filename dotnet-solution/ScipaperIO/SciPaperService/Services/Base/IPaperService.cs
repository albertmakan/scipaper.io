using SciPaperService.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace SciPaperService.Services.Base
{
    public interface IPaperService
    {
        Task<Paper> CreatePaperAsync(Paper paper);
        Paper UpdatePaper(Paper paperUpdate);
        Paper ReadPaper(string id);
        void DeletePaper(string id);
        List<Paper> GetAllByAuthor(string authorId);
        void Publish(string paperId, string authorId);
    }
}
