package constant

import "github.com/labstack/echo/v4"

func LoadStatic(app *echo.Echo) {
	// load static path
	app.Static("static", "repository/assets")

}