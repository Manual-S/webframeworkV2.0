package main

import (
	"net/http"

	"webframeworkV2.0/framework/gin"
)

func UserLoginController(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
