package http

import (
	"webframeworkV2.0/framework/gin"
	"webframeworkV2.0/provider/demo"
)

// Routes 绑定具体的路由
func Routes(r *gin.Engine) {
	r.Static()

	demo.Register(r)
}
