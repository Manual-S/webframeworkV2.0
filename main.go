package main

import (
	"fmt"

	"webframeworkV2.0/framework/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		res, ok := ctx.DefaultQueryBool("id", false)
		if ok {
			fmt.Println(res)
		}
	})
	r.Run()
}
