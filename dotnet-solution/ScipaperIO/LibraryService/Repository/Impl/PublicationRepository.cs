using LibraryService.Models;
using LibraryService.Repository.Contracts;
using LibraryService.Settings;

namespace LibraryService.Repository.Impl
{
    public class PublicationRepository : Repository<Publication>, IPublicationRepository
    {
        public PublicationRepository(IMongoDbSettings settings) : base(settings) { }
    }
}
