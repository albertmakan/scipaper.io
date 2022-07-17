using Kafka.Public;
using Kafka.Public.Loggers;
using System;
using System.Text;
using System.Text.Json;

namespace LibraryService.Messaging
{
    public interface IKafkaConsumer
    {
        void AddListener<T>(Action<T> listener);
    }

    public class KafkaConsumer : IKafkaConsumer
    {
        private readonly string _topic;
        private readonly ClusterClient _clusterClient;

        public KafkaConsumer(string topic)
        {
            _topic = topic;
            _clusterClient = new(new() { Seeds = "localhost:9092" }, new ConsoleLogger());
            _clusterClient.ConsumeFromLatest(_topic);
            AddListener<string>(msg => Console.WriteLine($"Message: {msg}"));
        }

        public void AddListener<T>(Action<T> action)
        {
            _clusterClient.MessageReceived += record =>
            {
                string msg = Encoding.UTF8.GetString(record.Value as byte[]);
                T obj = JsonSerializer.Deserialize<T>(msg);
                action(obj);
            };
        }
    }
}
