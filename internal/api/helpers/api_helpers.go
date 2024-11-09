package api_helpers

import (
	blogger "golang_template/pkg/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ResultSimple(ctx echo.Context, msg string, err error) error {
	if err != nil {
		blogger.Logger(blogger.AUTH_SERVICE_LOG, err, msg).Log()
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": msg,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": msg,
	})
}

func ResultCustom(ctx echo.Context, result interface{}, err error) error {
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)
	}

	return ctx.JSON(http.StatusOK, result)
}

func Int(i string) int {
	if s, err := strconv.Atoi(i); err == nil {
		return s
	}

	return 0
}

func String(i string, alt ...string) string {
	if len(i) == 0 {
		if len(alt) > 0 {
			return alt[0]
		}

		return "none"
	}

	return i
}

func Bool(i string) bool {
	if r, err := strconv.ParseBool(i); err == nil {
		return r
	}

	return false
}
