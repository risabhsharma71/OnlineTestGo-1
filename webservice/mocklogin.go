package webservice

import (
	"OnlineTestGo/models"

	"github.com/gin-gonic/gin"
)

func Mocklogin(c *gin.Context) {

	var login models.Login
	var login1 models.Login

	c.BindJSON(&login)

	if login.Email == "test@t.com" {
		login1 = models.Login{Uid: 10, Fname: "dj", Token: "ghctfdtf", UserType: "admin"}
	}

	//function should be calling this manager class
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": login1,
	})
}
