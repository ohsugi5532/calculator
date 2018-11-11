package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"./rpn"
)

// CalculatorGET is the handler for GET
func CalculatorGET() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	}
}

// CalculatorPOST is the handler for POST
func CalculatorPOST() echo.HandlerFunc {
	return func(c echo.Context) error {
		formula := c.FormValue("formula")

		result := rpn.Calculate(&formula)

		return c.String(http.StatusOK, result)
	}
}
