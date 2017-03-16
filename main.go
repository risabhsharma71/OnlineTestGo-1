package main

import (
	"OnlineTestGo/webservice"
	//"encoding/json"
	//"log"

	// "github.com/itsjamie/gin-cors"

	"github.com/gin-gonic/gin"
	//"net/http"
	//"github.com/gorilla/mux"
)

func main() {

	router := gin.Default()

	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/userAnswer", webservice.AnswerList)
	router.GET("/questions", webservice.QuestionList)
	//router.HandleFunc("/questions",webservice.GetQuestion).Methods("GET")
	//router.GET("/questions", webservice.GetQuestions)
	router.GET("/testService", webservice.TestService)
     router.GET("/admin",webservice.Admin)


	router.Run(":8080")

}