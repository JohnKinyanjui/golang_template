package dating_handler

import (
	api_helpers "golang_template/internal/api/helpers"
	dating_service "golang_template/services/dating"

	"github.com/labstack/echo/v4"
)

func (rt *DatingRouter) rejectUser(ctx echo.Context) error {
	id := ctx.Get("user_id").(string)

	type rejectRequest struct {
		ToUserID string `json:"to_user_id"`
	}

	var req rejectRequest
	if err := ctx.Bind(&req); err != nil {
		return api_helpers.ResultSimple(ctx, "invalid request data", err)
	}

	res, err := dating_service.RejectUser(id, req.ToUserID)
	if err != nil {
		return api_helpers.ResultSimple(ctx, err.Understandable, err.Error)
	}

	return api_helpers.ResultCustom(ctx, res, nil)
}
