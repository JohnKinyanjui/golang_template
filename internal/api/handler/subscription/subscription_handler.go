package subscription_handler

import (
	api_helpers "golang_template/internal/api/helpers"
	subscriptions_service "golang_template/services/subscriptions"

	"github.com/labstack/echo/v4"
)

func (rt *SubscriptionRouter) getMySubscription(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)

	subscription, err := subscriptions_service.GetUserSubscription(userID)
	if err != nil {
		return api_helpers.ResultSimple(ctx, "subscription not found", err)
	}

	return api_helpers.ResultCustom(ctx, subscription, nil)
}

func (rt *SubscriptionRouter) updateSubscription(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)

	var req subscriptions_service.UpdateSubscriptionRequest
	if err := ctx.Bind(&req); err != nil {
		return api_helpers.ResultSimple(ctx, "invalid request data", err)
	}

	subscription, err := subscriptions_service.UpdateSubscription(userID, req)
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to update subscription", err)
	}

	return api_helpers.ResultCustom(ctx, subscription, nil)
}

func (rt *SubscriptionRouter) getSubscriptionTiers(ctx echo.Context) error {
	tiers := subscriptions_service.GetSubscriptionTiers()
	return api_helpers.ResultCustom(ctx, tiers, nil)
}
