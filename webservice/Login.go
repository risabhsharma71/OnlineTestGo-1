package webservice

import (
	"OnlineTestGo/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var login models.Login
	var login1 models.Login

	c.BindJSON(&login)
login1:= c.Query("email")
	login1:=c.Query("password")

	login1 := manager.FetcLogin(Token)

	c.BindJSON(&login)
	if (login1.Email==login1.Password){
		return Token
	}
	login1 := manager.FetchUser(Token)

	// if login.Email == "testuser" {
	// 	login1 = models.Login{Uid: 10, Fname: "dj", Token: "ghctfdtf", UserType: "admin"}
	// }
    //     login1 := manager.Login(email)
	//function should be calling this manager class
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": login1,
	})
}
