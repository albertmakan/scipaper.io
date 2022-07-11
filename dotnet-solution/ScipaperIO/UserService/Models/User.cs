using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace UserService.Models
{
    public class User : Document
    {
        public string Username { get; set; }
        [JsonIgnore]
        public string Password { get; set; }
        [Required]
        public string FirstName { get; set; }
        [Required]
        public string LastName { get; set; }
        public string Email { get; set; }
    }
}
