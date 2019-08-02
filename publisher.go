package amqp

type Publisher interface {
	Publish(exchange, key string, message []byte) error
}
