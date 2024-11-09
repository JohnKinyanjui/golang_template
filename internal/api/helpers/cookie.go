package api_helpers

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func StoreCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "ksession_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(9999 * time.Hour)
	cookie.HttpOnly = true // Not accessible via JavaScript
	cookie.Secure = func() bool {
		return os.Getenv("ENV") == "production"
	}()

	cookie.Path = "/"
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}
