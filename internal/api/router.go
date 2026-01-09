package handler

import (
	accounts_handler "golang_template/internal/api/handler/accounts"
	auth_handler "golang_template/internal/api/handler/auth"
	dating_handler "golang_template/internal/api/handler/dating"
	"golang_template/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	private := e.Group("/api/v1")
	public := e.Group("/api/v1")
	private.Use(middlewares.AuthMiddleware)

	// routes
	{
		accounts_handler.Handler(private, public).Routes()
		auth_handler.Handler(private, public).Routes()
		dating_handler.Handler(private, public).Routes()
	}

}
