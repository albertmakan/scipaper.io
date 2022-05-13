package clients

import (
	"log"

	"github.com/streadway/amqp"
)

type AMQPSender struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func (sender *AMQPSender) Initialize(queueName string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}
	sender.connection = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %s", err)
	}
	sender.channel = ch

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Panicf("Failed to declare a queue: %s", err)
	}
	sender.queue = &q
}

func (sender *AMQPSender) Send(body []byte) {
	sender.channel.Publish(
		"",                // exchange
		sender.queue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func (sender *AMQPSender) Deinitialize() {
	defer sender.connection.Close()
	defer sender.channel.Close()
}