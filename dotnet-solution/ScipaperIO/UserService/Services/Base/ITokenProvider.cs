namespace UserService.Services.Base
{
    public interface ITokenProvider
    {
        string GenerateToken(string username);
    }
}
