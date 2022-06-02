package http

import "webframeworkV2.0/framework/gin"

// NewHttpEngine 创建一个绑定了路由的web引擎
func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 绑定路由

	return r, nil
}
