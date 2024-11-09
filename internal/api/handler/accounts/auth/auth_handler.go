package auth_handler

import (
	api_helpers "golang_template/internal/api/helpers"
	auth_service "golang_template/services/accounts/auth"

	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
)

var (
	count = 0
)

func (rt *AuthRouter) continueWithGoogle(ctx echo.Context) error {
	token, err := auth_service.ContinueWithGoogle(api_helpers.String(ctx.QueryParam("token")))
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to complete authentication", err)
	}

	api_helpers.StoreCookie(ctx, token)

	return api_helpers.ResultCustom(ctx, map[string]interface{}{"token": token}, nil)
}

func (rt *AuthRouter) continueWithEmail(ctx echo.Context) error {
	return nil
}

func (rt *AuthRouter) continueWithGithub(ctx echo.Context) error {
	res, err := auth_service.ContinueWithGithub(ctx.Param("code"))
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to complete authentication", err)
	}

	api_helpers.StoreCookie(ctx, res["access_token"])

	return api_helpers.ResultCustom(ctx, res, nil)
}

func (rt *AuthRouter) validateToken(ctx echo.Context) error {
	db := ctx.Get("db").(*surrealdb.DB)

	res, err := auth_service.ValidateToken(db)
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to complete authentication", err)
	}

	return api_helpers.ResultCustom(ctx, res, nil)
}
