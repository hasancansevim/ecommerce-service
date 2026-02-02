package elasticsearch

import (
	"fmt"
	"go-ecommerce-service/config"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog/log"
)

func NewElasticSearchClient(cfg config.ElasticSearchConfig) *elasticsearch.Client {
	address := fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port)

	esCfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}

	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Elasticsearch İstemcisi Oluşturulamadı")
	}

	for i := 0; i < 20; i++ {
		res, err := es.Info()
		if err == nil && !res.IsError() {
			res.Body.Close()
			log.Info().Str("address", address).Msg("✅ Elasticsearch Bağlantısı Başarılı! 🚀")
			return es
		}

		log.Warn().
			Err(err).
			Str("host", cfg.Host).
			Int("deneme", i+1).
			Msg("Elasticsearch'e bağlanılamadı, tekrar deneniyor...")

		time.Sleep(3 * time.Second)
	}

	log.Fatal().Msg("❌ Elasticsearch bağlantısı kurulamadı, pes ediliyor.")
	return nil
}
