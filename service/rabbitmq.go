package service

import (
	"context"
	"encoding/json"
	"env_middleware/data"
	"env_middleware/store"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type RabbitMq struct {
	Conn         *amqp.Connection
	Ch           *amqp.Channel
	exchangeName string
	qname        string
	redisClient  *store.RedisClient
}

func RunRabbitMqConsumer(exchangeName string, redis *redis.Client, url string) (*RabbitMq, error) {
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
		redisClient:  store.NewRedisClient(redis),
	}
	return &mq, nil
}

func (mq *RabbitMq) ConsumeMsgs(ctx context.Context) {
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

	// 消费消息并检查取消信号
	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer stopping gracefully...")
			return
		case d, ok := <-msgs:
			if !ok {
				// 消息通道关闭
				log.Println("Message channel closed.")
				return
			}
			crater := data.Crater{}
			err := json.Unmarshal(d.Body, &crater)
			if err != nil {
				log.Printf("Failed to unmarshal received crater, err:%s\n", err)
				continue
			}
			log.Printf("Received a new crater: lon(%v)-lat(%v), width(%v), depth(%v)\n",
				crater.Position.Longitude, crater.Position.Latitude, crater.Width, crater.Depth)

			err = mq.redisClient.InsertCrater(context.Background(), crater)
			if err != nil {
				log.Printf("Failed to insert crater: craterID: %s, err: %s", crater.CraterID, err)
			}
		}
	}
}
