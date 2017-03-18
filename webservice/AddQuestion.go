package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/models"
	// /"OnlineTestGo/utility"
    //"OnlineTestGo/authenticationfilter"
	"github.com/gin-gonic/gin"
)

func AddQuestions(c *gin.Context) {

	//token := c.Request.Header.Get("Authorization")
	// aunthenticationfilter.authenticateToken(token)
	//utility.GetToken(token)
    
	var question models.Question

	c.BindJSON(&question)

	//function should be calling this manager class

	insertedid := manager.AddQuestion(question)
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": insertedid,
	})

}