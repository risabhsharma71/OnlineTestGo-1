package webservice

import (
	"OnlineTestGo/models"

	"github.com/gin-gonic/gin"
)

func Mocklogin(c *gin.Context) {

	var login models.Login
	var login1 models.Login
	
	if user.Email == "test@t.com" {
	login1 = models.Login{Uid: 10, Fname: "dj", Token: "ghctfdtf", UserType: "user"}
		
	}

	c.BindJSON(&login)

	//if user.Email == "test@t.com" {
	//	login1 = models.Login{Uid: 10, Fname: "dj", Token: "ghctfdtf", UserType: "user"}
		
	//}

	//function should be calling this manager class
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  Success,
		"message": login1,
	})
}
