package accounts_service

import (
	"context"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"

	"github.com/google/uuid"
)

func GetAccount(id string) (interface{}, error) {
	ctx := context.Background()
	userID := uuid.MustParse(id)

	queries := query.New(db.PgConn)
	row, err := queries.GetMyAccount(ctx, userID)
	if err != nil {
		return nil, err
	}

	return row, nil
}
