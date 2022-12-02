package mortgageapi

import "github.com/labstack/echo/v4"

func Load(e *echo.Group) *echo.Group {
	e.POST("/calculate", Calculate)

	return e
}
