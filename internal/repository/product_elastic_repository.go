package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/DaniilKalts/elasticsearch-training/internal/domain"
)

type ProductElasticRepository interface {
	SearchProducts(query string) ([]*domain.Product, error)
}

type productElasticRepo struct {
	es *elasticsearch.Client
}

func NewProductElasticRepository(es *elasticsearch.Client) ProductElasticRepository {
	return &productElasticRepo{es: es}
}

func (r *productElasticRepo) SearchProducts(query string) (
	[]*domain.Product,
	error,
) {
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"name", "description"},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(esQuery); err != nil {
		return nil, fmt.Errorf("encode query: %w", err)
	}

	res, err := r.es.Search(
		r.es.Search.WithContext(context.Background()),
		r.es.Search.WithIndex("products"),
		r.es.Search.WithBody(&buf),
		r.es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("es search error: %w", err)
	}
	defer res.Body.Close()

	var result struct {
		Hits struct {
			Hits []struct {
				Source domain.Product `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	products := make([]*domain.Product, 0, len(result.Hits.Hits))
	for _, hit := range result.Hits.Hits {
		p := hit.Source
		products = append(products, &p)
	}

	return products, nil
}
