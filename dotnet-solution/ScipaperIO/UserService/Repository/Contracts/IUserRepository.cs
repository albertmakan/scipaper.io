using UserService.Models;

namespace UserService.Repository.Contracts
{
    public interface IUserRepository : IRepository<User>
    {
        User FindByUsername(string username);
    }
}
