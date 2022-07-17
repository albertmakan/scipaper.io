using Confluent.Kafka;
using System;
using System.Text.Json;

namespace SciPaperService.Messaging
{
    public interface IKafkaProducer
    {
        object Send(string topic, object message);
    }

    public class KafkaProducer : IKafkaProducer
    {
        private readonly ProducerConfig _config = new() { BootstrapServers = "localhost:9092" };

        public object Send(string topic, object message)
        {
            using (var producer = new ProducerBuilder<Null, string>(_config).Build())
            {
                try
                {
                    return producer.ProduceAsync(topic, new Message<Null, string> { Value = JsonSerializer.Serialize(message) })
                        .GetAwaiter()
                        .GetResult();
                }
                catch (Exception e)
                {
                    Console.WriteLine($"Oops, something went wrong: {e}");
                }
            }
            return null;
        }
    }
}
