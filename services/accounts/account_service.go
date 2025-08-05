package accounts_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"
	"golang_template/pkg/helpers"

	"github.com/google/uuid"
)

func GetAccount(id string) (interface{}, error) {
	res, err := db.Query.GetUser(context.Background(), uuid.MustParse(id))
	if err != nil {
		return res, err
	}

	return res, nil
}

func CreateAgent(id string, params CreateUserAgentParams) *helpers.ResultError {
	ctx := context.Background()
	if err := params.validate(); err != nil {
		reason := "invalid user agent parameters"
		return helpers.Error(reason, errors.New(reason))
	}

	if _, err := db.Query.CreateUserAgent(ctx, query.CreateUserAgentParams{
		UserID:      uuid.MustParse(id),
		Picture:     params.Picture,
		Name:        params.Name,
		Description: params.Description,
		PhoneNumber: params.PhoneNumber,
		Email:       params.Email,
	}); err != nil {
		return helpers.Error("invalid user agent creation", err)
	}

	return nil
}
