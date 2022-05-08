package main

import "webframeworkV2.0/framework/gin"

func registerRouter(core *gin.Engine) {
	core.GET("/subjest/list/all", SubjectListController)
	core.GET("/hello", UserLoginController)
}
