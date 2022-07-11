using System.Collections.Generic;

namespace SciPaperService.Models
{
    public class Paper : Document
    {
        public string Title { get; set; }
        public string Author { get; set; }
        public string AuthorId { get; set; }
        public List<Section> Sections { get; set; }
    }

    public class Section
    {
        public string Name { get; set; }
        public string Content { get; set; }
    }
}
