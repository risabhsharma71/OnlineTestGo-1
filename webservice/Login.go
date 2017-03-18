package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/models"
	"OnlineTestGo/tos"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var user models.User // request obj
	var token tos.Tokento
	c.BindJSON(&user)
	responsemessage := "failure"
	if user.Email == "test@t.com" {
		token = manager.AuthenticateUser(user)
		responsemessage = "success"
	}

	token = manager.AuthenticateUser(user)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  responsemessage,
		"message": token,
	})
}
