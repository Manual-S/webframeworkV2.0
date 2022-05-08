package main

import (
	"net/http"

	"webframeworkV2.0/provider/demo"

	"webframeworkV2.0/framework/gin"
)

func main() {
	core := gin.New()
	core.Bind(&demo.DemoServiceProvider{})
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
