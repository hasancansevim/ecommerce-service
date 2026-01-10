package worker

import (
	"encoding/json"
	"go-ecommerce-service/infrastructure/rabbitmq"
	"go-ecommerce-service/persistence"
	"log"
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
		log.Printf("❌ Worker Kuyruğa Bağlanamadı: %v", err)
		return
	}

	go func() {
		log.Println("👷‍♂️ Worker İş Başında! Siparişler bekleniyor...")

		for d := range msgs {
			order := OrderMessage{}
			json.Unmarshal(d.Body, &order)

			log.Printf("📩 Yeni İş : Mesaj alındı : %d", order.OrderId)

			_, err := w.repository.UpdateOrderStatus(order.OrderId, "Shipped")

			if err != nil {
				log.Println(err)
			}

			log.Println("✅ Mail Gönderildi ve Stok Güncellendi")
			d.Ack(false)
		}
	}()
}

type OrderMessage struct {
	OrderId int64 `json:"order_id"`
}
