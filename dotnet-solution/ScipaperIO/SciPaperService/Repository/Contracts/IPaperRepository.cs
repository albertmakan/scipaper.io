using SciPaperService.Models;
using System.Collections.Generic;

namespace SciPaperService.Repository.Contracts
{
    public interface IPaperRepository : IRepository<Paper>
    {
        IEnumerable<Paper> FindAllByAuthorId(string authorId);
    }
}
