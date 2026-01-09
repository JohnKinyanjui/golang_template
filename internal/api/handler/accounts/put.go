package accounts_handler

import (
	"errors"
	api_helpers "golang_template/internal/api/helpers"
	accounts_service "golang_template/services/accounts"

	"github.com/labstack/echo/v4"
)

func (rt *AccountRouter) updateMyAccount(ctx echo.Context) error {
	id := ctx.Get("user_id").(string)

	var params accounts_service.UpdateAccountParams
	if err := ctx.Bind(&params); err != nil {
		return api_helpers.ResultSimple(ctx, "invalid request data", err)
	}

	res, err := accounts_service.UpdateAccount(id, params)
	if err != nil {
		if errors.Is(err, accounts_service.ErrUsernameAlreadyExists) {
			return api_helpers.ResultSimple(ctx, "username already exists", err)
		}
		return api_helpers.ResultSimple(ctx, "unable to update profile", err)
	}

	return api_helpers.ResultCustom(ctx, res, nil)
}
