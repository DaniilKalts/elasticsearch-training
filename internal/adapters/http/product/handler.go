package product

import (
	"encoding/json"
	"net/http"

	"github.com/DaniilKalts/elasticsearch-training/internal/application"
)

type Handler struct {
	svc application.ProductService
}

func NewHandler(svc application.ProductService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
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
