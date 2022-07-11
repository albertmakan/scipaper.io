using SciPaperService.Models;
using SciPaperService.Repository.Contracts;
using SciPaperService.Settings;

namespace SciPaperService.Repository.Impl
{
    public class PaperRepository : Repository<Paper>, IPaperRepository
    {
        public PaperRepository(IMongoDbSettings settings) : base(settings) { }


    }
}
