package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/models"
      "OnlineTestGo/utility"
    "OnlineTestGo/filter"
	"github.com/gin-gonic/gin"
	"log"
)

func AddQuestions(c *gin.Context) {
utility.GetLogger()
//fmt.println("ab")
	token := c.Request.Header.Get("Authorization")
	log.Println(token)
	 bool:= filter.AuthenticateToken(token)
	 log.Println(bool)
	//utility.GetToken(token)
    if bool==true{
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

}else{

c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(401, gin.H{
		"status":  "failure",
		"message": "you dont have permission to acces",
	})





}
}