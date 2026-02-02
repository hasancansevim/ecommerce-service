package elasticsearch

import (
	"fmt"
	"go-ecommerce-service/config"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticSearchClient(cfg config.ElasticSearchConfig) *elasticsearch.Client {
	// http://localhost:9200
	address := fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port)
	esCfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}

	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		log.Fatal(err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	return es
}
