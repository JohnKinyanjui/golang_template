package auth_handler

import "github.com/labstack/echo/v4"

type AuthRouter struct {
	private *echo.Group
	public  *echo.Group
}

func Handler(private, public *echo.Group) *AuthRouter {
	return &AuthRouter{public: public, private: private}
}

func (rt *AuthRouter) Routes() {
	rt.public.POST("/auth/google", rt.continueWithGoogle)
	rt.public.POST("/auth/email", rt.continueWithEmail)

}
