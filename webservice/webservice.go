	package webservice

import (

	//"bytes"
	//"database/sql"
	"fmt"
        "log"

	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"
	//"net/http"
	"OnlineTestGo/models"
)

func RegisterUser(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)
        log.Println("Fname: " ,user.Fname)
	  log.Println("Lname: " ,user.Lname)
  		log.Println("phone: " ,user.Phone)
 	 log.Println("email: " ,user.Email)
   

	//function should be calling this manager class

	insertedid := manager.Register(user)
        c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
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

/*func QuestionList(c *gin.Context) {

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
*/