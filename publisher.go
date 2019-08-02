package amqp

type Publisher interface {
	Publish(exchange, key string, message []byte) error
}

type StateLessPublisher struct {
}

type StateFullPublisher struct {
}

func NewStateLessPublisher(uri string) (Publisher, error) {
	return nil, nil
}

func NewStateFullPublisher(uri string) (Publisher, error) {
	return nil, nil
}

func (slp *StateLessPublisher) Publish(exchange, key string, message []byte) error {
	return nil
}

func (sfp *StateFullPublisher) Publish(exchange, key string, message []byte) error {
	return nil
}
