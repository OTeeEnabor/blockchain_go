package main

import (
	"github.com/OTeeEnabor/blockchain_go/constant"
	"github.com/OTeeEnabor/blockchain_go/router"
	"github.com/OTeeEnabor/blockchain_go/server"
	"github.com/labstack/echo/v4"
)

func main() {
	//  initialize echo app
	app := echo.New()

	// render the templates
	app.Renderer = constant.LoadTemplate()

	//  load static  files
	constant.LoadStatic(app)
	
	// define URL
	router.LoadAllRouters(app)

	// start the server 
	server.RunServer(app)

}