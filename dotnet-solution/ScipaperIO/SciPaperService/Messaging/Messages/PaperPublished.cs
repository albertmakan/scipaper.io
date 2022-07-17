namespace SciPaperService.Messaging.Messages
{
    public class PaperPublished
    {
        public string PaperId { get; set; }
        public string Title { get; set; }
        public string Author { get; set; }
    }
}
