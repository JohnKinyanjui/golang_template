package properties_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"
	"golang_template/pkg/helpers"
	"log"

	"github.com/google/uuid"
)

func GetProperties(params query.GetPropertiesParams) (any, *helpers.ResultError) {
	log.Print(params)
	properties, err := db.Query.GetProperties(context.Background(), params)
	if err != nil {
		return nil, helpers.Error("failed to fetch properties", err)
	}

	return properties, nil
}

func GetPropertyInfo(propertyID string) (any, error) {
	propertyUUID, err := uuid.Parse(propertyID)
	if err != nil {
		return nil, errors.New("invalid property ID")
	}

	property, err := db.Query.GetPropertyInfo(context.Background(), propertyUUID)
	if err != nil {
		return nil, err
	}

	return &property, nil
}
