package broker

import (
	"context"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RMQConnection interface {
	Channel() (*amqp.Channel, error)
	Close() error
}

type Broker struct {
	name       string
	connection RMQConnection
}

type Message struct {
	Ctx  context.Context
	Data []byte
}

func New(name string, uri string) (*Broker, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	return &Broker{
		name:       name,
		connection: conn,
	}, nil
}

func (b *Broker) Connect(
	ctx context.Context,
	exchangeName string,
	exchangeType string,
	queueName string,
) error {
	ch, err := b.connection.Channel()
	if err != nil {
		return fmt.Errorf("open channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("exchange declare: %w", err)
	}

	queue, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("queue declare: %w", err)
	}

	err = ch.QueueBind(
		queue.Name,   // name of the queue
		"",           // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("queue bind: %w", err)
	}

	return nil
}

func (b *Broker) Close(
	ctx context.Context,
) error {
	return b.connection.Close()
}

func (b *Broker) Consume(ctx context.Context, queueName string) (<-chan Message, error) {
	messages := make(chan Message)
	ch, _ := b.connection.Channel()

	go func() {
		<-ctx.Done()
		if err := ch.Close(); err != nil {
			log.Println(err)
		}
	}()

	deliveries, err := ch.Consume(
		queueName,
		b.name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("start consuming: %w", err)
	}

	go func() {
		defer func() {
			close(messages)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case del := <-deliveries:
				if err := del.Ack(false); err != nil {
					log.Println(err)
				}

				msg := Message{
					Ctx:  context.TODO(),
					Data: del.Body,
				}

				select {
				case <-ctx.Done():
					return
				case messages <- msg:
				}
			}
		}
	}()

	return messages, nil
}

func (b *Broker) Publish(ctx context.Context, exchange string, routingKey string, body []byte) error {
	channel, err := b.connection.Channel()
	if err != nil {
		return fmt.Errorf("channel: %w", err)
	}
	if err = channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		return fmt.Errorf("exchange publish: %w", err)
	}
	return nil
}
