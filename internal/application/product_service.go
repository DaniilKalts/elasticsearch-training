package application

import (
	"github.com/DaniilKalts/elasticsearch-training/internal/domain"
	"github.com/DaniilKalts/elasticsearch-training/internal/repository"
)

type productService struct {
	repo   repository.ProductRepository
	esRepo repository.ProductElasticRepository
}

type ProductService interface {
	GetProducts() ([]*domain.Product, error)
	SearchProducts(query string) ([]*domain.Product, error)
}

func NewProductService(
	repo repository.ProductRepository,
	esRepo repository.ProductElasticRepository,
) ProductService {
	return &productService{repo: repo, esRepo: esRepo}
}

func (s *productService) GetProducts() ([]*domain.Product, error) {
	return s.repo.GetProducts()
}

func (s *productService) SearchProducts(query string) (
	[]*domain.Product, error,
) {
	return s.esRepo.SearchProducts(query)
}
