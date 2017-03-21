package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var user models.User // request obj

	c.BindJSON(&user)
	//token := c.Request.Header.Get("Authorization")
	message := manager.Login(user)
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "sucess",
		"message": message,
	})
}
