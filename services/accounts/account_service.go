package accounts_service

import (
	"context"
	"golang_template/internal/db"
)

func GetAccount(id string) (interface{}, error) {
	var res MyAccount
	err := db.PgConn.QueryRow(context.Background(), getAccountQuery, id).Scan(&res.Picture, &res.FullName, &res.Email)
	if err != nil {
		return res, err
	}

	return res, nil
}
