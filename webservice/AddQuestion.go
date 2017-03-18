package webservice

import (
	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"
 "OnlineTestGo/utility"
 "OnlineTestGo/authenticationfilter"
	"OnlineTestGo/models"
) 
func AddQuestions(c *gin.Context) {
        authenticationfilter.authenticateToken(c *gin.Context)

    var question models.Question
       
    c.BindJSON(&question)
       // authenticationfilter.authenticateToken(c *gin.Context)

    //function should be calling this manager class

        insertedid := manager.AddQuestion(question)
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")    
    c.JSON(200, gin.H{
        "status":  "success",
        "message":  insertedid,
    })
    

        
}