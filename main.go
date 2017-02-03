package main

import (
	"github.com/MIghtykukulkan/OnlineTestGo/webservice"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/userAnswer",webservice.AnswerList)
	//define other service here

	router.Run(":9090")
}
