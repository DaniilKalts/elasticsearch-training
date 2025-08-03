package product

import "github.com/DaniilKalts/elasticsearch-training/internal/domain"

type Service interface {
	GetProducts() ([]*domain.Product, error)
}
