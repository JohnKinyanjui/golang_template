package subscription_handler

import "github.com/labstack/echo/v4"

type SubscriptionRouter struct {
	private *echo.Group
	public  *echo.Group
}

func Handler(private, public *echo.Group) *SubscriptionRouter {
	return &SubscriptionRouter{public: public, private: private}
}

func (rt *SubscriptionRouter) Routes() {
	rt.private.GET("/subscription", rt.getMySubscription)
	rt.private.PUT("/subscription", rt.updateSubscription)
	rt.public.GET("/subscription/tiers", rt.getSubscriptionTiers)
}
