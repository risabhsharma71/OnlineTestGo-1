package webservice

import (
	"github.com/gin-gonic/gin"
)

func Mocklogin(c *gin.Context) {

	//function should be calling this manager class
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": true,
	})
}
