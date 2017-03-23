package webservice

import (
	"OnlineTestGo/filter"
	"OnlineTestGo/manager"
	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func QuestionList(c *gin.Context) {
	//recieving query param "testtype"
	utility.GetLogger()
	testtype := c.Query("testtype")
	token := c.Request.Header.Get("Authorization")
	log.Println(testtype)
	log.Println(token)
	bool := filter.AuthenticateToken(token)
	//utility.GetToken(token)
	if bool == true {
		Qlist := manager.FetchQuestion(testtype)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

		c.JSON(200, gin.H{
			"status":  "success",
			"message": Qlist,
		})

	} else {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(401, gin.H{
			"status":  "failure",
			"message": "you dont have permission to acces",
		})

	}
}
