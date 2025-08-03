package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/DaniilKalts/elasticsearch-training/internal/adapters/database"
	"github.com/DaniilKalts/elasticsearch-training/internal/application/config"
	"github.com/DaniilKalts/elasticsearch-training/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}
	defer db.Close()

	esClient := database.NewElasticClient(cfg)

	productRepo := repository.NewProductRepository(db)

	products, err := productRepo.GetProducts()
	if err != nil {
		log.Fatalf("failed to get products: %v", err)
	}

	log.Printf("Found %d products. Start indexing...\n", len(products))

	for _, p := range products {
		data, _ := json.Marshal(p)

		_, err := esClient.Index(
			"products",
			bytes.NewReader(data),
			esClient.Index.WithDocumentID(p.ID.String()),
			esClient.Index.WithContext(context.Background()),
		)

		if err != nil {
			log.Printf("Failed to index product %s: %v", p.ID, err)
		}
	}

	log.Println("All products indexed successfully.")
}
