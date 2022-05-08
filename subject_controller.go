package main

import (
	"net/http"

	"webframeworkV2.0/framework/gin"
	"webframeworkV2.0/provider/demo"
)

func SubjectListController(c *gin.Context) {
	demoService := c.MustMake(demo.Key).(demo.Service)

	foo := demoService.GetFoo()

	c.JSON(http.StatusOK, foo)
}
