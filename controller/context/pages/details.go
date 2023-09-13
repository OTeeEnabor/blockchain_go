package pages

import (
	"fmt"
	"net/http"

	"github.com/OTeeEnabor/blockchain_go/db/product"
	"github.com/labstack/echo/v4"
)

func DetailsContext(c echo.Context) error {
	getProductId := c.Param("productId")
	fmt.Println(getProductId)
	// get a db record
	dbRecord, _ := product.GetProductDetails(getProductId)
	return c.Render(http.StatusOK, "detail.html", dbRecord)
}