package controller

import (
	"github.com/gin-gonic/gin"
	"pg-manager/utils"
)

func Test(c *gin.Context) {
	a := c.Query("name")
	c.JSON(200, utils.SuccessMess("success", a))

}
