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
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	Available   bool
	Category    Category
	Brand       Brand
	Rating      float64

	CreatedAt time.Time
	UpdatedAt time.Time
}
