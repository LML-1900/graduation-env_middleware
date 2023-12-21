package service

import (
	"encoding/json"
	"env_middleware/data"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func RunRabbitMqConsumer(exchangeName string) {
	conn, err := amqp.Dial("amqp://guest:guest@172.17.0.2:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ, err:%s\n", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel, err:%s\n", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Printf("Failed to declare an exchange, err:%s\n", err)
		return
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue, err:%s\n", err)
		return
	}

	err = ch.QueueBind(
		q.Name,                  // queue name
		data.CRATER_ROUTING_KEY, // routing key
		exchangeName,            // exchange
		false,
		nil)
	if err != nil {
		log.Fatalf("Failed to bind a queue, err:%s\n", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer, err:%s\n", err)
		return
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			crater := data.Crater{}
			err := json.Unmarshal(d.Body, &crater)
			if err != nil {
				log.Printf("fail to unmarshal received crater, err:%s\n", err)
			}
			log.Printf("received a new crater: lon(%v)-lat(%v), width(%v), depth(%v)\n",
				crater.Position.Longitude, crater.Position.Latitude, crater.Width, crater.Depth)
		}
	}()

	<-forever
}
