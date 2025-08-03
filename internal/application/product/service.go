package product

import (
	"github.com/DaniilKalts/elasticsearch-training/internal/domain"
	"github.com/DaniilKalts/elasticsearch-training/internal/repository"
)

type service struct {
	repo repository.ProductRepository
}

func NewService(repo repository.ProductRepository) Service {
	return &service{repo: repo}
}

func (s *service) GetProducts() ([]*domain.Product, error) {
	return s.repo.GetProducts()
}
