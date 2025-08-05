package auth_handler

import (
	api_helpers "golang_template/internal/api/helpers"
	auth_service "golang_template/services/auth"

	"github.com/labstack/echo/v4"
)

func (rt *AuthRouter) continueWithGoogle(ctx echo.Context) error {
	token, err := auth_service.ContinueWithGoogle(api_helpers.String(ctx.QueryParam("token")), auth_service.AuthParams{
		Option: auth_service.SignIn,
	})
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to complete authentication", err)
	}

	api_helpers.StoreCookie(ctx, token)

	return api_helpers.ResultCustom(ctx, map[string]any{"token": token}, nil)
}

func (rt *AuthRouter) continueWithEmail(ctx echo.Context) error {
	var params auth_service.EmailParams
	if err := ctx.Bind(&params); err != nil {
		return api_helpers.ResultSimple(ctx, "invalid request data", err)
	}

	result, err := auth_service.ContinueWithEmail(params)
	if err != nil {
		return api_helpers.ResultSimple(ctx, err.Understandable, err.Error)
	}

	// api_helpers.StoreCookie(ctx, token)

	return api_helpers.ResultCustom(ctx, result, nil)
}
