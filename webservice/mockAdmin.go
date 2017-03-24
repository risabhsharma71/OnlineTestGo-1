package webservice

import (
	"OnlineTestGo/tos"

	"github.com/gin-gonic/gin"
)

func MockAdmin(c *gin.Context) {

	//var user models.User // request obj
	var dlist tos.Admin
	//c.BindJSON(&user)

	//	responsemessage := "failure"
	//	if user.Email == "test@t.com" {
	dlist = tos.Admin{Uid: 1, Fname: "Sweta", Lname: "Vahia", Java: 20, Fundamental: 30, Language: 25}
	//		responsemessage = "success"

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "successs",
		"message": dlist,
	})
}
