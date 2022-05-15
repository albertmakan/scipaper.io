package messaging

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

type AMQPConsumerFunction func([]byte)

type AMQPReceiver struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func (receiver *AMQPReceiver) Initialize(queueName string) {
	conn, err := amqp.Dial(os.Getenv("AMPQ_CONNECTION"))
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}
	receiver.connection = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %s", err)
	}
	receiver.channel = ch

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
	receiver.queue = &q
}

func (receiver *AMQPReceiver) Consume(consumerFunc AMQPConsumerFunction) {
	msgs, err := receiver.channel.Consume(
		receiver.queue.Name, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	if err != nil {
		log.Panicf("Failed to register a consumer: %s", err)
	}

	go func() {
		for d := range msgs {
			consumerFunc(d.Body)
		}
	}()
}

func (receiver *AMQPReceiver) Deinitialize() {
	defer receiver.connection.Close()
	defer receiver.channel.Close()
}