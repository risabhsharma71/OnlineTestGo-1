	package webservice

import (

	//"bytes"
	//"database/sql"
	"fmt"

       // "net/http"

        "log"


	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"
	//"net/http"
	"OnlineTestGo/models"
//"github.com/gorilla/mux"
//"encoding/json"
	

)
/*
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

*/



func QuestionList(c *gin.Context) {


	var totalquestion models.TotalQuestion
	c.BindJSON(&totalquestion)
        testtype :=c.Query("type")
	questionlist := manager.FetchQuestion(totalquestion,testtype) 

	//fmt.Println("questionlist", questionlist)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": questionlist,
	})
}


/*
func GetQuestion(w http.ResponseWriter, req *http.Request,c *gin.Context) {
var questions []models.TotalQuestion
//c.BindJSON(&questions)
questions = append(questions,models.TotalQuestion{ID: "1", Question: "which of these is necesary condn for automatic type conversion?", Option1: "a", Option2: "b",Option3: "c", Option4: "d",Option5: "e",Type :"java"})
//questions = append(questions,models.TotalQuestion{ID: "2", Question: "what is the output of this program?", Option1: "a", Option2: "b",Option3: "c", Option4: "d",Option5: "e",Type :"java"})


	params := mux.Vars(req)
	for _, item := range questions {
		if item.Type == params["type"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(models.TotalQuestion{})
}

func GetQuestions(w http.ResponseWriter, req *http.Request,c *gin.Context) {
	json.NewEncoder(w).Encode(questions)
}
*/


func TestService(c *gin.Context) {
				
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "your webserivce is reachable",
	})
}
