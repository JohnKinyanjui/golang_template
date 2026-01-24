package accounts_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

var ErrFullNameAlreadyExists = errors.New("full_name already exists")

type UpdateAccountParams struct {
	FullName *string `json:"full_name"`
	Phone    *string `json:"phone"`
}

func UpdateAccount(id string, params UpdateAccountParams) (interface{}, error) {
	ctx := context.Background()
	userID := uuid.MustParse(id)

	queries := query.New(db.PgConn)

	arg := query.UpdateMyAccountParams{
		ID:       userID,
		FullName: pgtype.Text{Valid: false},
		Phone:    pgtype.Text{Valid: false},
	}

	if params.FullName != nil {
		arg.FullName = pgtype.Text{String: *params.FullName, Valid: true}
	}
	if params.Phone != nil {
		arg.Phone = pgtype.Text{String: *params.Phone, Valid: true}
	}

	row, err := queries.UpdateMyAccount(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, ErrFullNameAlreadyExists
			}
		}
		return nil, err
	}

	res := MyAccount{
		Picture:  row.Picture,
		FullName: row.FullName,
		Email:    row.Email,
		Phone:    row.Phone.String,
	}

	return res, nil
}
