package main

import (
	"OnlineTestGo/webservice"
       // "github.com/itsjamie/gin-cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()


	
        
	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/userAnswer", webservice.AnswerList)
	//router.GET("/userQuestion", webservice.QuestionList)
	//router.GET("/testService", webservice.TestService)
	

	router.Run(":8082")
}
