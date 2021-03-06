using Kafka.Public;
using Kafka.Public.Loggers;
using LibraryService.Messaging.Messages;
using System;
using System.Text;
using System.Text.Json;

namespace LibraryService.Messaging
{
    public interface IKafkaConsumer<T> where T : IMessage
    {
        void AddListener(Action<T> listener);
    }

    public class KafkaConsumer<T> : IKafkaConsumer<T> where T : IMessage
    {
        private readonly string _topic;
        private readonly ClusterClient _clusterClient;

        public KafkaConsumer(string topic)
        {
            _topic = topic;
            _clusterClient = new(new() { Seeds = "localhost:9092" }, new ConsoleLogger());
            _clusterClient.ConsumeFromLatest(_topic);
        }

        public void AddListener(Action<T> action)
        {
            _clusterClient.MessageReceived += record =>
            {
                string msg = Encoding.UTF8.GetString(record.Value as byte[]);
                try
                {
                    T obj = JsonSerializer.Deserialize<T>(msg);
                    action(obj);
                }
                catch (Exception)
                {
                    Console.WriteLine($"Invalid message received: {msg}");
                }
            };
        }
    }
}
