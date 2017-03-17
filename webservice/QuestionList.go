package webservice

import (
	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"
)

func QuestionList(c *gin.Context) {
	//recieving query param "testtype"
	testtype := c.Query("testtype")

	Qlist := manager.FetchQuestion(testtype)
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": Qlist,
	})

}
