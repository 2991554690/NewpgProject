package router

import (
	"github.com/gin-gonic/gin"
	"pg-manager/controller"
)

func TestRouter(engine *gin.Engine) {
	test := engine.Group("test")
	{
		test.GET("/test/test", controller.Test)
	}
}
