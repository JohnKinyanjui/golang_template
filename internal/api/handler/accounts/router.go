package accounts_handler

import "github.com/labstack/echo/v4"

type AccountRouter struct {
	private *echo.Group
	public  *echo.Group
}

func Handler(private, public *echo.Group) *AccountRouter {
	return &AccountRouter{public: public, private: private}
}

func (rt *AccountRouter) Routes() {
	rt.private.GET("/account/my", rt.getMyAccount)

}
