package mortgageapi

import (
	"net/http"

	"github.com/highercomve/mortgage_calculator/modules/insurace/insuranceservices"
	"github.com/highercomve/mortgage_calculator/modules/mortgage/mortgagemodels"
	"github.com/highercomve/mortgage_calculator/modules/mortgage/mortgageservices"
	"github.com/labstack/echo/v4"
)

// Calculate Calculate mortgage
// @Summary Calculate mortgage
// @Description Calculate mortgage
// @Accept  json
// @Produce  json
// @Param body body mortgagemodels.InputP true "Input mortgage calculator"
// @Success 200 {object} mortgagemodels.OutputP "Mortgage response"
// @Failure 400 {object} echo.HTTPError "Input is not valid"
// @Failure 500 {object} echo.HTTPError "Error processing request"
// @Router /calculate [get]
func Calculate(c echo.Context) (err error) {
	payload := new(mortgagemodels.InputP)

	if err = c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	service := mortgageservices.NewDefaultService(insuranceservices.NewCMHC())
	resp, err := service.Calculate(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
