using UserService.Models;
using UserService.Repository.Contracts;
using UserService.Settings;

namespace UserService.Repository.Impl
{
    public class UserRepository : Repository<User>, IUserRepository
    {
        public UserRepository(IMongoDbSettings settings) : base(settings) { }

        public User FindByUsername(string username)
        {
            return FindOne(user => user.Username == username);
        }
    }
}
