using Microsoft.AspNetCore.Mvc;
using UserService.DTO;
using UserService.Filters;
using UserService.Models;
using UserService.Services.Base;

namespace UserService.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    [Produces("application/json")]
    [ValidateModel]
    public class UserController : ControllerBase
    {
        private readonly IUserService _userService;

        public UserController(IUserService userService)
        {
            _userService = userService;
        }

        [HttpPost("register")]
        public User CreateUser(User user)
        {
            return _userService.CreateUser(user);
        }

        [HttpGet("user/{id}")]
        public User ReadUser(string id)
        {
            return _userService.GetUser(id);
        }

        [HttpPost("auth")]
        public string Authenticate([FromBody] LoginRequest loginRequest)
        {
            return _userService.Authenticate(loginRequest.Username, loginRequest.Password);
        }

        [HttpGet("getname/{username}")]
        public string GetName(string username)
        {
            return _userService.GetUserName(username);
        }
    }
}
