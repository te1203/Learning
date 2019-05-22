package main

import (
	"handler"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// RouterInit will init the router
func RouterInit() *fasthttprouter.Router {
	user := "testuser"
	pass := "testuser!!!"

	router := fasthttprouter.New()
	router.GET("/", handler.Index)
	router.GET("/protected/", handler.BasicAuthHandler(handler.Protected, user, pass))

	// Serve static files from the ./public directory
	router.NotFound = fasthttp.FSHandler("./public", 0)

	return router

}
