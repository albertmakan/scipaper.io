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
        private readonly IHttpClientFactory _httpClientFactory;

        public UserClient(IHttpClientFactory httpClientFactory)
        {
            _httpClientFactory = httpClientFactory;
        }

        public string GetName(string username)
        {
            return new GetNameCommand(_httpClientFactory)
            {
                Username = username
            }
            .Execute();
        }

    }

    public class GetNameCommand : HystrixCommand<string>
    {
        private readonly HttpClient _httpClient;
        public string Username { get; set; }

        public GetNameCommand(IHttpClientFactory httpClientFactory)
            : base(HystrixCommandGroupKeyDefault.AsKey("UserGroup"))
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
