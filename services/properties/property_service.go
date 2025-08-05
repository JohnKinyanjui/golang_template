package properties_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	subscriptions_service "golang_template/services/subscriptions"

	"github.com/google/uuid"
)

func CreateProperty(userID string, req CreatePropertyRequest) (*Property, error) {
	// Check subscription limits
	userUUID, _ := uuid.Parse(userID)
	if err := checkPropertyLimit(userUUID); err != nil {
		return nil, err
	}

	propertyUUID := uuid.New()

	var property Property
	err := db.PgConn.QueryRow(context.Background(), createPropertyQuery,
		propertyUUID, userUUID, req.Title, req.Description, req.Price,
		req.Address, req.City, req.State, req.ZipCode, req.Bedrooms,
		req.Bathrooms, req.SquareFeet, req.PropertyType, req.Images,
	).Scan(
		&property.ID, &property.UserID, &property.Title, &property.Description,
		&property.Price, &property.Address, &property.City, &property.State,
		&property.ZipCode, &property.Bedrooms, &property.Bathrooms,
		&property.SquareFeet, &property.PropertyType, &property.Status,
		&property.Images, &property.CreatedAt, &property.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &property, nil
}

func GetUserProperties(userID string) ([]Property, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	rows, err := db.PgConn.Query(context.Background(), getUserPropertiesQuery, userUUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []Property
	for rows.Next() {
		var property Property
		err := rows.Scan(
			&property.ID, &property.UserID, &property.Title, &property.Description,
			&property.Price, &property.Address, &property.City, &property.State,
			&property.ZipCode, &property.Bedrooms, &property.Bathrooms,
			&property.SquareFeet, &property.PropertyType, &property.Status,
			&property.Images, &property.CreatedAt, &property.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}

	return properties, nil
}

func UpdateProperty(propertyID, userID string, req UpdatePropertyRequest) (*Property, error) {
	propertyUUID, err := uuid.Parse(propertyID)
	if err != nil {
		return nil, errors.New("invalid property ID")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var property Property
	err = db.PgConn.QueryRow(context.Background(), updatePropertyQuery,
		propertyUUID, req.Title, req.Description, req.Price, req.Address,
		req.City, req.State, req.ZipCode, req.Bedrooms, req.Bathrooms,
		req.SquareFeet, req.PropertyType, req.Status, req.Images, userUUID,
	).Scan(
		&property.ID, &property.UserID, &property.Title, &property.Description,
		&property.Price, &property.Address, &property.City, &property.State,
		&property.ZipCode, &property.Bedrooms, &property.Bathrooms,
		&property.SquareFeet, &property.PropertyType, &property.Status,
		&property.Images, &property.CreatedAt, &property.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &property, nil
}

func DeleteProperty(propertyID, userID string) error {
	propertyUUID, err := uuid.Parse(propertyID)
	if err != nil {
		return errors.New("invalid property ID")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	_, err = db.PgConn.Exec(context.Background(), deletePropertyQuery, propertyUUID, userUUID)
	return err
}

func checkPropertyLimit(userID uuid.UUID) error {
	subscription, err := subscriptions_service.GetUserSubscription(userID.String())
	if err != nil {
		return errors.New("subscription not found")
	}

	count, err := getUserPropertyCount(userID)
	if err != nil {
		return err
	}

	if count >= subscription.MaxProperties {
		return errors.New("property limit reached for your subscription tier")
	}

	return nil
}

func getUserPropertyCount(userID uuid.UUID) (int, error) {
	var count int
	err := db.PgConn.QueryRow(context.Background(), getUserPropertyCountQuery, userID).Scan(&count)
	return count, err
}
