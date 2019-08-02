package amqp

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type Publisher interface {
	PublishJson(exchange, key string, message interface{}) error
}

type stateFullPublisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewStateFullPublisher(uri string) (Publisher, error) {
	connection, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &stateFullPublisher{
		connection: connection,
		channel:    channel,
	}, nil
}

func (sfp *stateFullPublisher) PublishJson(exchange, key string, message interface{}) error {
	// serialize interface into json
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// publish the message
	return sfp.channel.Publish(exchange, key, false, false, amqp.Publishing{
		Headers:         amqp.Table{},
		ContentType:     "application/json",
		ContentEncoding: "",
		Body:            bytes,
		DeliveryMode:    amqp.Persistent,
		Priority:        0,
	})
}
