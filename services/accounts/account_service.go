package accounts_service

import (
	"context"
	"golang_template/internal/db"

	"github.com/google/uuid"
)

func GetAccount(id string) (interface{}, error) {
	res, err := db.Query.GetUser(context.Background(), uuid.MustParse(id))
	if err != nil {
		return res, err
	}

	return res, nil
}
