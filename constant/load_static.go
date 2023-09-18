package constant

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func LoadStatic(app *echo.Echo) {
	//  get app path
	path, _ := os.Executable()
	// // // get file path
	filePath := filepath.Dir(path)
	// // // store  template folder path into variable (CLOUD server)
	staticFolder := fmt.Sprintf("%v/repository/assets",filePath)

	// store the favicon file path into variable (Cloud server)
	faviconFilePath := fmt.Sprintf("%v/images/favicon.ico",staticFolder)


	// load the favicon
	app.File("/favicon.ico", faviconFilePath)
	
	// load static path for assets like css, js, fonts, images
	app.Static("static", staticFolder)

}