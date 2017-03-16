package webservice

import (
	"fmt"
	//"io"
        //"io/ioutil"
        "log"
      //  "os"

	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"

	"OnlineTestGo/models"
)
func RegisterUser(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)
	

	//function should be calling this manager class

	insertedid := manager.Register(user)
	c.Header("Access-Control-Allow-Origin", "*")
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

	c.Header("Access-Control-Allow-Origin", "*")
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
	c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": Qlist,
	})

}
func AddQuestions(c *gin.Context) {

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

func TestService(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "your webserivce is reachable",
	})
}