package accounts_handler

import (
	api_helpers "golang_template/internal/api/helpers"
	accounts_service "golang_template/services/accounts"

	"github.com/labstack/echo/v4"
)

func (rt *AccountRouter) getMyAccount(ctx echo.Context) error {
	id := ctx.Get("user_id").(string)
	t, err := accounts_service.GetAccount(id)
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to complete authentication", err)
	}

	return api_helpers.ResultCustom(ctx, t, nil)
}
