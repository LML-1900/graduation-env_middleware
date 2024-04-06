package service

import (
	"encoding/json"
	"env_middleware/data"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"log"
)

type RabbitMq struct {
	Conn         *amqp.Connection
	Ch           *amqp.Channel
	exchangeName string
	qname        string
}

func RunRabbitMqConsumer(exchangeName string) (*RabbitMq, error) {
	url := viper.GetString("rabbitmq.url")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

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
		return nil, err
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
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,                  // queue name
		data.CRATER_ROUTING_KEY, // routing key
		exchangeName,            // exchange
		false,
		nil)
	if err != nil {
		return nil, err
	}
	mq := RabbitMq{
		Conn:         conn,
		Ch:           ch,
		exchangeName: exchangeName,
		qname:        q.Name,
	}
	return &mq, nil
}

func (mq *RabbitMq) ConsumeMsgs() {
	msgs, err := mq.Ch.Consume(
		mq.qname, // queue
		"",       // consumer
		true,     // auto ack
		false,    // exclusive
		false,    // no local
		false,    // no wait
		nil,      // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer, err:%s\n", err)
		return
	}

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

}
