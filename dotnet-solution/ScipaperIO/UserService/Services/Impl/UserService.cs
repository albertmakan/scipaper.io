using UserService.Exceptions;
using UserService.Models;
using UserService.Repository.Contracts;
using UserService.Services.Base;
using BCryptNet = BCrypt.Net.BCrypt;

namespace UserService.Services.Impl
{
    public class UserService : IUserService
    {
        private readonly IUserRepository _userRepository;
        private readonly ITokenProvider _tokenProvider;

        public UserService(IUserRepository repository, ITokenProvider tokenProvider)
        {
            _userRepository = repository;
            _tokenProvider = tokenProvider;
        }

        public string Authenticate(string username, string password)
        {
            var user = _userRepository.FindByUsername(username);
            if (user == null)
                throw new UnauthorizedException("invalid username or password");
            if (!BCryptNet.Verify(password, user.Password))
                throw new UnauthorizedException("invalid username or password");
            return _tokenProvider.GenerateToken(username);
        }

        public User CreateUser(User user)
        {
            if (_userRepository.FindByUsername(user.Username) != null)
                throw new BadLogicException("Username already exists");
            user.Password = BCryptNet.HashPassword(user.Password);
            return _userRepository.InsertOne(user);
        }

        public User GetUser(string id)
        {
            return _userRepository.FindById(id);
        }

        public string GetUserName(string token)
        {
            string username = token;
            var user = _userRepository.FindByUsername(username);
            if (user == null) return "";
            return $"{username} {user.FirstName} {user.LastName}";
        }

        public bool IsAuthenticated(string token)
        {
            return true;
        }
    }
}
