package server

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

func RunServer(app *echo.Echo) {
	/*
	This is a function to setup the server. 
	Here the following options are set:

	Max Concurrent Streams: the number of concurrent requests the client can make to the server is limited by this attribute.
	MaxReadFrameSize: sets the maximum number of frames a server will read. 
	IdleTimeOut: idle duration time for client

	
	*/
	// declare server configuration
	server := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	
	if  err := app.StartH2CServer(":3000", server); err != http.ErrServerClosed {
		log.Fatal(err)
	  }
}