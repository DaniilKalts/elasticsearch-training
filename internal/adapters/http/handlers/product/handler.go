package product

import (
	"encoding/json"
	"net/http"

	"github.com/DaniilKalts/elasticsearch-training/internal/application/product"
)

type handler struct {
	svc product.Service
}

func NewHandler(svc product.Service) *handler {
	return &handler{svc: svc}
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
