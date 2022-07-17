using SciPaperService.Models;
using SciPaperService.Repository.Contracts;
using SciPaperService.Settings;
using System.Collections.Generic;

namespace SciPaperService.Repository.Impl
{
    public class PaperRepository : Repository<Paper>, IPaperRepository
    {
        public PaperRepository(IMongoDbSettings settings) : base(settings) { }

        public IEnumerable<Paper> FindAllByAuthorId(string authorId) => FilterBy(paper => paper.AuthorId == authorId);
    }
}
