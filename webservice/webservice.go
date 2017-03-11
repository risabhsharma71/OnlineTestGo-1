package webservice

import (
	"fmt"
	"log"

	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"

	"OnlineTestGo/models"
)

func RegisterUser(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)
	log.Println("Fname: ", user.Fname)
	log.Println("Lname: ", user.Lname)
	log.Println("phone: ", user.Phone)
	log.Println("email: ", user.Email)

	//function should be calling this manager class

	insertedid := manager.Register(user)
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": insertedid,
	})
}

func AnswerList(c *gin.Context) {

	//Answer list recieves a list of answers
	var answer []models.Answer
	c.BindJSON(&answer)
	log.Println(answer)
	score := manager.CalculateScore(answer)

	fmt.Println("answerlist", score)

	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status": "success",
		"score":  score,
	})
}

func QuestionList(c *gin.Context) {
	//recieving query param "testtype"
	testtype := c.Query("testtype")

	Qlist := manager.FetchQuestion(testtype)
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": Qlist,
	})

}

func TestService(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "your webserivce is reachable",
	})
}