namespace LibraryService.Models
{
    public class Publication : Document
    {
        public string PaperId { get; set; }
        public string Title { get; set; }
        public string Author { get; set; }
    }
}
