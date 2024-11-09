package middlewares

import (
	"net/http"

	"strings"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware checks for session token in the cookie or Authorization header
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var foundToken string
		// Check for session token in the cookie
		cookie, err := c.Cookie("access_token")
		if err == nil && cookie != nil {
			sessionToken := cookie.Value
			if sessionToken != "" {
				foundToken = sessionToken
			}
		}

		// If no valid session token found in cookie, check for Bearer token
		if len(foundToken) == 0 {
			token := c.Request().Header.Get("Authorization")
			if token != "" {
				tokenString := strings.Split(token, " ")
				if len(tokenString) == 2 && tokenString[0] == "Bearer" {
					// Get DB connection using Bearer token
					foundToken = tokenString[1]
				}
			}
		}

		if len(foundToken) > 0 {
			claim, err := verifyJWT(foundToken)
			if err == nil {
				c.Set("user_id", claim["sub"].(string))
				return next(c)
			}
		}

		// If neither token is valid, return unauthorized error
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
	}
}
