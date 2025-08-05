package properties_service

import (
	"time"

	"github.com/google/uuid"
)

type Property struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	ZipCode      string    `json:"zip_code"`
	Bedrooms     int       `json:"bedrooms"`
	Bathrooms    int       `json:"bathrooms"`
	SquareFeet   int       `json:"square_feet"`
	PropertyType string    `json:"property_type"`
	Status       string    `json:"status"`
	Images       []string  `json:"images"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreatePropertyRequest struct {
	Title        string   `json:"title" validate:"required"`
	Description  string   `json:"description"`
	Price        float64  `json:"price" validate:"required,gt=0"`
	Address      string   `json:"address" validate:"required"`
	City         string   `json:"city" validate:"required"`
	State        string   `json:"state" validate:"required"`
	ZipCode      string   `json:"zip_code" validate:"required"`
	Bedrooms     int      `json:"bedrooms"`
	Bathrooms    int      `json:"bathrooms"`
	SquareFeet   int      `json:"square_feet"`
	PropertyType string   `json:"property_type" validate:"required"`
	Images       []string `json:"images"`
}

type UpdatePropertyRequest struct {
	Title        string   `json:"title" validate:"required"`
	Description  string   `json:"description"`
	Price        float64  `json:"price" validate:"required,gt=0"`
	Address      string   `json:"address" validate:"required"`
	City         string   `json:"city" validate:"required"`
	State        string   `json:"state" validate:"required"`
	ZipCode      string   `json:"zip_code" validate:"required"`
	Bedrooms     int      `json:"bedrooms"`
	Bathrooms    int      `json:"bathrooms"`
	SquareFeet   int      `json:"square_feet"`
	PropertyType string   `json:"property_type" validate:"required"`
	Status       string   `json:"status"`
	Images       []string `json:"images"`
}
