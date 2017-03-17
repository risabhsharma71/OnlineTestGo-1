package webservice

import (
	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"

	"OnlineTestGo/models"
)

func RegisterUser(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)

	//function should be calling this manager class

	insertedid := manager.Register(user)
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": insertedid,
	})
}
