package dating_service

import (
	"context"
	"golang_template/internal/db"
	"golang_template/pkg/helpers"

	"github.com/google/uuid"
)

func GetHistory(userID string) (any, *helpers.ResultError) {
	ctx := context.Background()
	uid := uuid.MustParse(userID)

	rows, err := db.Query.ListDatingHistory(ctx, uid)
	if err != nil {
		return nil, helpers.Error("unable to load history", err)
	}

	return rows, nil
}
