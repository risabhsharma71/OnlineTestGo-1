package webservice

import (
	"OnlineTestGo/models"

	"github.com/gin-gonic/gin"
)

func Mocklogin(c *gin.Context) {

	var user models.User
	var login1 models.Login
	message := "failure"

	c.BindJSON(&user)

	if user.Email == "test@t.com" {
		login1 = models.Login{Uid: 10, Fname: "dj", Token: "ghctfdtf", UserType: "admin"}
		message = "success"
	}

	//function should be calling this manager class
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  message,
		"message": login1,
	})
}
