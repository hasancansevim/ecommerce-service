package worker

import (
	"encoding/json"
	"go-ecommerce-service/infrastructure/rabbitmq"
	"go-ecommerce-service/persistence"

	"github.com/rs/zerolog/log"
)

type OrderWorker struct {
	client     *rabbitmq.RabbitMQClient
	repository persistence.IOrderRepository
}

func NewOrderWorker(client *rabbitmq.RabbitMQClient, repository persistence.IOrderRepository) *OrderWorker {
	return &OrderWorker{
		client:     client,
		repository: repository,
	}
}

func (w *OrderWorker) Start() {
	msgs, err := w.client.Channel.Consume(
		"order_created_queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error().Err(err).Msg("❌ Worker could not connect to queue")
		return
	}

	go func() {
		log.Info().Msg("👷‍♂️ Worker ready! Waiting for orders...")

		for d := range msgs {
			order := OrderMessage{}
			json.Unmarshal(d.Body, &order)

			log.Info().Int64("order_id", order.OrderId).Msg("📩 New job received")

			_, err := w.repository.UpdateOrderStatus(order.OrderId, "Shipped")

			if err != nil {
				log.Error().Err(err).Msg("Order update failed")
			}

			log.Info().Int64("order_id", order.OrderId).Msg("✅ Stock updated and status changed")
			d.Ack(false)
		}
	}()
}

type OrderMessage struct {
	OrderId int64 `json:"order_id"`
}
