package webservice

import (
	"OnlineTestGo/manager"
    "OnlineTestGo/utility"
	"github.com/gin-gonic/gin"

	"OnlineTestGo/models"
) 
func AddQuestions(c *gin.Context) {
utility. GetToken(c *gin.Context,token models.Token)
    var question models.Question
       
    c.BindJSON(&question)
        

    //function should be calling this manager class

        insertedid := manager.AddQuestion(question)
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")    
    c.JSON(200, gin.H{
        "status":  "success",
        "message":  insertedid,
    })
    

        
}