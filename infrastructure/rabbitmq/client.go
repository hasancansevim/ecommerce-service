package rabbitmq

import (
	"fmt"
	"go-ecommerce-service/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type IRabbitMQClient interface {
	Publish(exchange, routingKey string, mandatory, immediate bool, msg amqp.Publishing) error
}

type RabbitMQClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQClient(cfg config.RabbitMQConfig) (*RabbitMQClient, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
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

func (rc *RabbitMQClient) Publish(exchange, routingKey string, mandatory, immediate bool, msg amqp.Publishing) error {
	return rc.Channel.Publish(exchange, routingKey, mandatory, immediate, msg)
}
