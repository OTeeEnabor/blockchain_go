package http

import (
	"github.com/OTeeEnabor/blockchain_go/controller/context/pages"
	"github.com/labstack/echo/v4"
)
//  define a function that routes to the index page
func DetailRouter(app *echo.Echo) {
	app.GET("/:productId", pages.DetailsContext)
}