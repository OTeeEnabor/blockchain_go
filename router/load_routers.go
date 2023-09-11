package router

import (
	"github.com/OTeeEnabor/blockchain_go/router/http"
	"github.com/labstack/echo/v4"
)

func LoadAllRouters(app *echo.Echo) {
	// index router
	http.IndexRouter(app)
	// form/post router
	http.FormRouter(app)
}