package router

import "github.com/gin-gonic/gin"

func GetEngine() *gin.Engine {
	engine := gin.Default() //创建一个gin引擎
	//UserRouter(engine)
	TestRouter(engine)
	return engine
}
