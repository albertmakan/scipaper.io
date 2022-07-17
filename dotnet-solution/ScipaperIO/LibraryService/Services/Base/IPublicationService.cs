using LibraryService.Messaging.Messages;
using LibraryService.Models;
using System.Collections.Generic;

namespace LibraryService.Services.Base
{
    public interface IPublicationService
    {
        IEnumerable<Publication> Search();
        void PublishListener(PaperPublished paperPublished);
    }
}
