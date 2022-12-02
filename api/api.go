package api

import (
	"github.com/highercomve/mortgage_calculator/modules/mortgage/mortgageapi"
	"github.com/labstack/echo/v4"
)

// LoadAPI load all API
func LoadAPI(e *echo.Echo) *echo.Group {
	api := e.Group("api/v1/")

	mortgageapi.Load(api)

	return api
}
