package worker

import (
	"go-ecommerce-service/infrastructure/rabbitmq"
	"log"
	"time"
)

type OrderWorker struct {
	client *rabbitmq.RabbitMQClient
}

func NewOrderWorker(client *rabbitmq.RabbitMQClient) *OrderWorker {
	return &OrderWorker{
		client: client,
	}
}

func (w *OrderWorker) Start() {
	msgs, err := w.client.Channel.Consume(
		"order_created_queue", // dinlenecek kuyruk
		"",
		true,
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
			log.Printf("📩 Yeni İş : Mesaj alındı : %s", d.Body)

			// Simulasyon - Mail atma işlemi 3 sn sürsün
			time.Sleep(3 * time.Second)

			log.Println("✅ Mail Gönderildi ve Stok Güncellendi")
		}
	}()
}
