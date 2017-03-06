package webservice

import (

	//"bytes"
	//"database/sql"
	"fmt"

	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"
	//"net/http"
	"OnlineTestGo/models"
)

func RegisterUser(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)
	//function should be calling this manager class

	insertedid := manager.Register(user)
	c.JSON(200, gin.H{
		"status":  "success",
		"message": insertedid,
	})
}

func AnswerList(c *gin.Context) {

	var answer models.Answer
	c.BindJSON(&answer)

	answerlist := manager.Answer(answer)

	fmt.Println("answerlist", answerlist)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": answerlist,
	})
}

func QuestionList(c *gin.Context) {

	var question models.Question
	c.BindJSON(&question)
        testtype :=c.Query("type")
	questionlist := manager.FetchQuestion(question,testtype) 

	fmt.Println("questionlist", questionlist)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": questionlist,
	})
}

func TestService(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "your webserivce is reachable",
	})
}
