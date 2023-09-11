package pages

import "github.com/labstack/echo/v4"

func FormContext(c echo.Context) error {
	productID := c.FormValue("product_id")
	quantity := c.FormValue("quantity")
	cornColour :=c.FormValue("corn_colour")
}