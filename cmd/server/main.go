package main

import (
	"log"

	"github.com/DaniilKalts/elasticsearch-training/internal/adapters/database"
	"github.com/DaniilKalts/elasticsearch-training/internal/adapters/http"
	httpProduct "github.com/DaniilKalts/elasticsearch-training/internal/adapters/http/product"
	"github.com/DaniilKalts/elasticsearch-training/internal/application"
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

	productRepo := repository.NewProductRepository(db)

	esClient := database.NewElasticClient(cfg)
	productESRepo := repository.NewProductElasticRepository(esClient)

	productSvc := application.NewProductService(productRepo, productESRepo)
	productHandler := httpProduct.NewHandler(productSvc)

	if err := http.StartServer(cfg, productHandler); err != nil {
		log.Fatal(err)
	}
}
