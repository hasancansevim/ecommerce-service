package elasticsearch

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticSearchClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{ // Elasticsearch'ün ayar objesi
		// Elasticsearch ün çalıştığı yer
		Addresses: []string{"http://localhost:9200"},
	}

	es, err := elasticsearch.NewClient(cfg) // Go ile Elasticsearch bunun üzerinden konuşacak
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
