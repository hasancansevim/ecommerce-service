package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQClient(user, password, host, port string) (*RabbitMQClient, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RabbitMq: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Failed to open a channel: %v", err)
	}

	_, err = ch.QueueDeclare(
		"order_created_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to declare a queue: %v", err)
	}
	return &RabbitMQClient{conn, ch}, nil
}

func (rc *RabbitMQClient) Close() {
	if rc.Channel != nil {
		rc.Channel.Close()
	}
	if rc.Conn != nil {
		rc.Conn.Close()
	}
}
