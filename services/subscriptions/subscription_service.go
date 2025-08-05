package subscriptions_service

import (
	"context"
	"errors"
	"golang_template/internal/db"

	"github.com/google/uuid"
)

func GetUserSubscription(userID string) (*Subscription, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var subscription Subscription
	err = db.PgConn.QueryRow(context.Background(), getUserSubscriptionQuery, userUUID).Scan(
		&subscription.ID, &subscription.UserID, &subscription.Tier, &subscription.Status,
		&subscription.MaxProperties, &subscription.Features, &subscription.CreatedAt, &subscription.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func UpdateSubscription(userID string, req UpdateSubscriptionRequest) (*Subscription, error) {
	// Validate user ID format
	if _, err := uuid.Parse(userID); err != nil {
		return nil, errors.New("invalid user ID")
	}

	// Get current subscription
	currentSub, err := GetUserSubscription(userID)
	if err != nil {
		return nil, err
	}

	// Get tier details
	tier, exists := SubscriptionTiers[req.Tier]
	if !exists {
		return nil, errors.New("invalid subscription tier")
	}

	// Update subscription
	var subscription Subscription
	err = db.PgConn.QueryRow(context.Background(), updateSubscriptionQuery,
		currentSub.ID, req.Tier, tier.MaxProperties, tier.Features,
	).Scan(
		&subscription.ID, &subscription.UserID, &subscription.Tier, &subscription.Status,
		&subscription.MaxProperties, &subscription.Features, &subscription.CreatedAt, &subscription.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func GetSubscriptionTiers() map[string]SubscriptionTier {
	return SubscriptionTiers
}

func CreateDefaultSubscription(userID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	freeTier := SubscriptionTiers["free"]
	_, err = db.PgConn.Exec(context.Background(), createSubscriptionQuery,
		userUUID, freeTier.Tier, freeTier.MaxProperties, freeTier.Features,
	)

	return err
}
