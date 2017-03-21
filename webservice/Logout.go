package webservice

import (
	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	message := manager.DeleteToken(token)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "sucess",
		"message": message,
	})

}
