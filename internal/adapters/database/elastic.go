package database

import (
	"fmt"
	"log"

	"github.com/DaniilKalts/elasticsearch-training/internal/application/config"
	es "github.com/elastic/go-elasticsearch/v8"
)

func NewElasticClient(cfg *config.Config) *es.Client {
	esCfg := es.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%d", cfg.Elastic.Host, cfg.Elastic.Port),
		},
	}

	client, err := es.NewClient(esCfg)
	if err != nil {
		log.Fatalf("Error creating ES client: %s", err)
	}

	return client
}
