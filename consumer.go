package amqp

import "github.com/streadway/amqp"

type Consumer interface {
	Consume(queue string, autoAck bool, callback func(deliveries <-chan amqp.Delivery, done chan error)) error
	Shutdown() error
}

type consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	done       chan error
}

func NewConsumer(uri string, prefetch int) (Consumer, error) {
	connection, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	if err := channel.Qos(prefetch, 0, false); err != nil {
		return nil, err
	}

	return &consumer{
		connection: connection,
		channel:    channel,
	}, nil
}

func (c *consumer) Consume(queue string, autoAck bool, callback func(deliveries <-chan amqp.Delivery, done chan error)) error {
	deliveries, err := c.channel.Consume(queue, "", autoAck, false, false, false, nil)
	if err != nil {
		return err
	}

	go callback(deliveries, c.done)

	return nil
}

func (c *consumer) Shutdown() error {
	_ = c.channel.Close()
	_ = c.connection.Close()
	return <-c.done
}
