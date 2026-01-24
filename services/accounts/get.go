package accounts_service

import (
	"context"
	"golang_template/internal/db"

	"github.com/google/uuid"
)

func GetAccount(id string) (any, error) {
	ctx := context.Background()
	userID := uuid.MustParse(id)

	row, err := db.Query.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return row, nil
}
