using UserService.Models;

namespace UserService.Services.Base
{
    public interface IUserService
    {
        User CreateUser(User user);
        User GetUser(string username);

        string Authenticate(string username, string password);
        bool IsAuthenticated(string token);
        string GetUserName(string token);
    }
}
