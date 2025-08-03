package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category string
type Brand string

const (
	CategorySmartphone Category = "Smartphones"
	CategoryLaptop     Category = "Laptops"
	CategoryAccessory  Category = "Accessories"
	CategoryTV         Category = "TV"
)

const (
	BrandApple   Brand = "Apple"
	BrandSamsung Brand = "Samsung"
	BrandSony    Brand = "Sony"
	BrandXiaomi  Brand = "Xiaomi"
	BrandOther   Brand = "Other"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Available   bool      `json:"available"`
	Category    Category  `json:"category"`
	Brand       Brand     `json:"brand"`
	Rating      float64   `json:"rating"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
