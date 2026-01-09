package dating_handler

import (
	api_helpers "golang_template/internal/api/helpers"
	dating_service "golang_template/services/dating"

	"github.com/labstack/echo/v4"
)

func (rt *DatingRouter) getFeed(ctx echo.Context) error {
	id := ctx.Get("user_id").(string)

	items, err := dating_service.GetFeed(id)
	if err != nil {
		return api_helpers.ResultSimple(ctx, err.Understandable, err.Error)
	}

	return api_helpers.ResultCustom(ctx, items, nil)
}
