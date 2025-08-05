package properties_handler

import "github.com/labstack/echo/v4"

type PropertyRouter struct {
	private *echo.Group
	public  *echo.Group
}

func Handler(private, public *echo.Group) *PropertyRouter {
	return &PropertyRouter{public: public, private: private}
}

func (rt *PropertyRouter) Routes() {
	rt.public.GET("/properties", rt.getMyProperties)
	rt.public.GET("/properties/info", rt.getPropertyInfo)

	// rt.private.GET("/properties", rt.getMyProperties)
	rt.private.PUT("/properties/:id", rt.updateProperty)
	rt.private.DELETE("/properties/:id", rt.deleteProperty)
}
