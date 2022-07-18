using System.Text.Json;

namespace LibraryService.Messaging.Messages
{
    public class PaperPublished : IMessage
    {
        public string PaperId { get; set; }
        public string Title { get; set; }
        public string Author { get; set; }

        public override string ToString() => JsonSerializer.Serialize(this);
    }
}
