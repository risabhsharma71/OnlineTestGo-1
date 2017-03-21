package main

import (
	"OnlineTestGo/webservice"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/goinggo/tracelog"
)

func main() {
	//	utility.GetLogger()
	tracelog.Start(tracelog.LevelTrace)
	tracelog.Trace("main", "main", " Trace")
	tracelog.Info("main", "main", " Info")
	tracelog.Warning("main", "main", " Warn")
	tracelog.Errorf(fmt.Errorf("Exception At..."), "main", "main", " Error")

	router := gin.Default()

	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/userAnswer", webservice.AnswerList)
	router.GET("/questions", webservice.QuestionList)

	router.POST("/addquestions", webservice.AddQuestions)
	router.GET("/testService", webservice.TestService)

	router.GET("/admin", webservice.Admin)
	router.POST("/mocklogin", webservice.Mocklogin)
	router.GET("/mocklogout", webservice.Mocklogout)

	//	router.POST("/login", webservice.Login)
	router.GET("/logout", webservice.Logout)

	router.Run(GetPort())

}

//Getport() method gets the OS env variable for the port, if not available, it will se the default port
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8083"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
