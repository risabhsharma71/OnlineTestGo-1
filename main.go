package main

import (
	"OnlineTestGo/webservice"
<<<<<<< HEAD
         //"encoding/json"
	//"log"


=======
       // "github.com/itsjamie/gin-cors"
>>>>>>> 97c3febc037c9c7f2a2f3a7b89708240fa055d19
	"github.com/gin-gonic/gin"
         //"net/http"

	//"github.com/gorilla/mux"

)

func main() {

	router := gin.Default()
        //router := mux.NewRouter()


<<<<<<< HEAD
	//router.POST("/registerUser", webservice.RegisterUser)
	//router.POST("/userAnswer", webservice.AnswerList)
	router.GET("/userQuestion", webservice.QuestionList)
        //router.HandleFunc("/questions",webservice.GetQuestion).Methods("GET")
	//router.HandleFunc("/questions/{type}", webservice.GetQuestions).Methods("GET")

	router.GET("/testService", webservice.TestService)
	//define other service here

	router.Run(":8081")
=======
	
        
	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/userAnswer", webservice.AnswerList)
	//router.GET("/userQuestion", webservice.QuestionList)
	//router.GET("/testService", webservice.TestService)
	

	router.Run(":8082")
>>>>>>> 97c3febc037c9c7f2a2f3a7b89708240fa055d19
}
