package handler

import (
	auth_handler "golang_template/internal/api/handler/accounts/auth"
	"golang_template/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	private := e.Group("/api/v1")
	public := e.Group("/api/v1")
	private.Use(middlewares.AuthMiddleware)

	// routes
	{
		auth_handler.Handler(private, public).Routes()

	}

}
