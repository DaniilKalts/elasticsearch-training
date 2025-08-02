package repository

import (
	"database/sql"

	"github.com/DaniilKalts/elasticsearch-training/internal/domain"
)

type productRepo struct {
	db *sql.DB
}

type ProductRepository interface {
	GetProducts() ([]*domain.Product, error)
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) GetProducts() ([]*domain.Product, error) {
	rows, err := r.db.Query(
		`SELECT
			id, 
			name, 
			description, 
			price,
			available,
			category,
			brand,
			rating,
			created_at,
			updated_at 
		FROM 
			products
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*domain.Product, 0)

	for rows.Next() {
		p := &domain.Product{}
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Available,
			&p.Category,
			&p.Brand,
			&p.Rating,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
