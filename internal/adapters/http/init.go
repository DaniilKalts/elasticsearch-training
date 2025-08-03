package http

import (
	"fmt"
	"net/http"

	"github.com/DaniilKalts/elasticsearch-training/internal/adapters/http/product"
	"github.com/DaniilKalts/elasticsearch-training/internal/application/config"
)

func StartServer(cfg *config.Config, handler *product.Handler) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/products", handler.GetProducts)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: mux,
	}

	return server.ListenAndServe()
}
