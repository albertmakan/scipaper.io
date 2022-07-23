using Steeltoe.CircuitBreaker.Hystrix;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading.Tasks;

namespace SciPaperService.Services.Dependencies
{
    public interface IUserClient
    {
        string GetName(string username);
    }


    public class UserClient : IUserClient
    {
        private readonly GetNameCommand _getNameCommand;

        public UserClient(GetNameCommand getNameCommand)
        {
            _getNameCommand = getNameCommand;
        }

        public string GetName(string username)
        {
            _getNameCommand.Username = username;
            return _getNameCommand.Execute();
        }

    }

    public class GetNameCommand : HystrixCommand<string>
    {
        public string Username { get; set; }
        private readonly HttpClient _httpClient;

        public GetNameCommand(IHystrixCommandOptions options, IHttpClientFactory httpClientFactory) : base(options)
        {
            _httpClient = httpClientFactory.CreateClient("user");
        }

        protected override async Task<string> RunAsync()
        {
            return await _httpClient.GetFromJsonAsync<string>($"getname/{Username}");
        }

        protected override string RunFallback()
        {
            return "";
        }
    }
}
