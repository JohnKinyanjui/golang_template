package properties_handler

import (
	api_helpers "golang_template/internal/api/helpers"
	query "golang_template/internal/db/generated"
	properties_service "golang_template/services/properties"

	"github.com/labstack/echo/v4"
)

func (rt *PropertyRouter) createProperty(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)

	var req properties_service.CreatePropertyRequest
	if err := ctx.Bind(&req); err != nil {
		return api_helpers.ResultSimple(ctx, "invalid request data", err)
	}

	property, err := properties_service.CreateProperty(userID, req)
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to create property", err)
	}

	return api_helpers.ResultCustom(ctx, property, nil)
}

func (rt *PropertyRouter) getPropertyInfo(ctx echo.Context) error {
	property, err := properties_service.GetPropertyInfo(ctx.QueryParam("id"))
	if err != nil {
		return api_helpers.ResultSimple(ctx, "property not found", err)
	}

	return api_helpers.ResultCustom(ctx, property, nil)
}

func (rt *PropertyRouter) getMyProperties(ctx echo.Context) error {
	properties, err := properties_service.GetProperties(query.GetPropertiesParams{
		Name:         api_helpers.String(ctx.QueryParam("name")),
		City:         api_helpers.String(ctx.QueryParam("city")),
		State:        api_helpers.String(ctx.QueryParam("state")),
		MinPrice:     float64(api_helpers.Int(ctx.QueryParam("min_price"))),
		MaxPrice:     float64(api_helpers.Int(ctx.QueryParam("max_price"))),
		PropertyType: api_helpers.String(ctx.QueryParam("property_type")),
		SaleType:     api_helpers.String(ctx.QueryParam("sale_type")),
		Skip:         int32(api_helpers.Int(ctx.QueryParam("skip"))),
	})
	if err != nil {
		return api_helpers.ResultSimple(ctx, err.Understandable, err.Error)
	}

	return api_helpers.ResultCustom(ctx, properties, nil)
}

func (rt *PropertyRouter) updateProperty(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	propertyID := ctx.Param("id")

	var req properties_service.UpdatePropertyRequest
	if err := ctx.Bind(&req); err != nil {
		return api_helpers.ResultSimple(ctx, "invalid request data", err)
	}

	property, err := properties_service.UpdateProperty(propertyID, userID, req)
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to update property", err)
	}

	return api_helpers.ResultCustom(ctx, property, nil)
}

func (rt *PropertyRouter) deleteProperty(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	propertyID := ctx.Param("id")

	err := properties_service.DeleteProperty(propertyID, userID)
	if err != nil {
		return api_helpers.ResultSimple(ctx, "unable to delete property", err)
	}

	return api_helpers.ResultSimple(ctx, "property deleted successfully", nil)
}
