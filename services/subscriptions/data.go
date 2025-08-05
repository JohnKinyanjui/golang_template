package subscriptions_service

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Tier          string    `json:"tier"`
	Status        string    `json:"status"`
	MaxProperties int       `json:"max_properties"`
	Features      []string  `json:"features"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type SubscriptionTier struct {
	Tier          string   `json:"tier"`
	Name          string   `json:"name"`
	MaxProperties int      `json:"max_properties"`
	Features      []string `json:"features"`
	Price         float64  `json:"price"`
}

type UpdateSubscriptionRequest struct {
	Tier string `json:"tier" validate:"required,oneof=free basic pro"`
}

var SubscriptionTiers = map[string]SubscriptionTier{
	"free": {
		Tier:          "free",
		Name:          "Free",
		MaxProperties: 1,
		Features:      []string{"Basic property listing", "Standard support"},
		Price:         0,
	},
	"basic": {
		Tier:          "basic",
		Name:          "Basic",
		MaxProperties: 5,
		Features:      []string{"Up to 5 properties", "Priority support", "Property analytics"},
		Price:         9.99,
	},
	"pro": {
		Tier:          "pro",
		Name:          "Professional",
		MaxProperties: 50,
		Features:      []string{"Up to 50 properties", "Premium support", "Advanced analytics", "Featured listings"},
		Price:         29.99,
	},
}
