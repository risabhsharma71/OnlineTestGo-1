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

	questionlist := manager.Register(user)

	//println statements should be replaced with logs
	fmt.Println("questionlist", questionlist)

	//response should return then question and answer list
	c.JSON(200, gin.H{
		"status":  "success",
		"message": questionlist,
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

	questionlist := manager.Question(question)

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
