using RabbitMQ.Client;
using System;

namespace SciPaperService.Messaging
{
    public abstract class RabbitMqClientBase : IDisposable
    {
        protected const string QueueName = "PUBLISH_PAPER";

        protected IModel Channel { get; private set; }
        private IConnection _connection;
        private readonly ConnectionFactory _connectionFactory;

        protected RabbitMqClientBase(ConnectionFactory connectionFactory)
        {
            _connectionFactory = connectionFactory;
            ConnectToRabbitMq();
        }

        private void ConnectToRabbitMq()
        {
            if (_connection == null || !_connection.IsOpen)
            {
                _connection = _connectionFactory.CreateConnection();
            }

            if (Channel == null || !Channel.IsOpen)
            {
                Channel = _connection.CreateModel();
                Channel.QueueDeclare(QueueName, durable: false, exclusive: false, autoDelete: false);
            }
        }

        public void Dispose()
        {
            try
            {
                Channel?.Close();
                Channel?.Dispose();
                Channel = null;

                _connection?.Close();
                _connection?.Dispose();
                _connection = null;
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex);
            }
        }
    }
}
