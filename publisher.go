package amqp

type Publisher interface {
	Publish(exchange, key string, message []byte) error
}

type stateLessPublisher struct {
}

type stateFullPublisher struct {
}

func NewStateLessPublisher(uri string) (Publisher, error) {
	return nil, nil
}

func NewStateFullPublisher(uri string) (Publisher, error) {
	return nil, nil
}

func (slp *stateLessPublisher) Publish(exchange, key string, message []byte) error {
	return nil
}

func (sfp *stateFullPublisher) Publish(exchange, key string, message []byte) error {
	return nil
}
