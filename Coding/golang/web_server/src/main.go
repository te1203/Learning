package main

import (
	"log"

	"github.com/AubSs/fasthttplogger"
	"github.com/valyala/fasthttp"
)

func main() {
	router := RouterInit()

	server := &fasthttp.Server{
		Handler: fasthttplogger.CombinedColored(router.Handler),
		Name:    "LearningWebServerWithGo",
	}

	log.Fatal(server.ListenAndServe(":8080"))
}
