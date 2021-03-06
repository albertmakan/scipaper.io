using Confluent.Kafka;
using SciPaperService.Messaging.Messages;
using System;

namespace SciPaperService.Messaging
{
    public interface IKafkaProducer
    {
        object Send(string topic, IMessage message);
    }

    public class KafkaProducer : IKafkaProducer
    {
        private readonly ProducerConfig _config = new() { BootstrapServers = "localhost:9092" };

        public object Send(string topic, IMessage message)
        {
            using (var producer = new ProducerBuilder<Null, string>(_config).Build())
            {
                try
                {
                    return producer.ProduceAsync(topic, new Message<Null, string> { Value = message.ToJson() })
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
