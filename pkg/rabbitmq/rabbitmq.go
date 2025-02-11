package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func Consume(queue string, ch *amqp.Channel, out chan amqp.Delivery) error {
	messages, err := ch.Consume(queue, "go-consumer", false, false, false, false, nil)
	if err != nil {
		return err
	}
	for message := range messages {
		out <- message
	}
	return nil
}
