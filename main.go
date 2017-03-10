package main

import (
	"OnlineTestGo/webservice"
         //"encoding/json"
	//"log"


	"github.com/gin-gonic/gin"
         //"net/http"

	//"github.com/gorilla/mux"

)

func main() {

	router := gin.Default()
        //router := mux.NewRouter()


	//router.POST("/registerUser", webservice.RegisterUser)
	//router.POST("/userAnswer", webservice.AnswerList)
	router.GET("/userQuestion", webservice.QuestionList)
        //router.HandleFunc("/questions",webservice.GetQuestion).Methods("GET")
	//router.HandleFunc("/questions/{type}", webservice.GetQuestions).Methods("GET")

	router.GET("/testService", webservice.TestService)
	//define other service here

	router.Run(":8081")
}
