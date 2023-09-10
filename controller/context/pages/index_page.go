package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexContext(c echo.Context) error {
	//  create index context function
	// input parameter is an echo interface
	// returns a HTTP status code with a string
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Oselu",
	})
}