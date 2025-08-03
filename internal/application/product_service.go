package application

import (
	"github.com/DaniilKalts/elasticsearch-training/internal/domain"
	"github.com/DaniilKalts/elasticsearch-training/internal/repository"
)

type productService struct {
	repo repository.ProductRepository
}

type ProductService interface {
	GetProducts() ([]*domain.Product, error)
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProducts() ([]*domain.Product, error) {
	return s.GetProducts()
}
