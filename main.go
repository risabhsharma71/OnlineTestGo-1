package main

import (
	"OnlineTestGo/webservice"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//utility.GetLogger()
	/*tracelog.Start(tracelog.LevelTrace)
	tracelog.Trace("main", "main", " Trace")
	tracelog.Info("main", "main", " Info")
	tracelog.Warning("main", "main", " Warn")
	//tracelog.Errorf(fmt.Errorf("Exception At..."), "main", "main", " Error")*/

	router := gin.Default()

	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/userAnswer", webservice.AnswerList)
	router.GET("/questions", webservice.QuestionList)
	//router.HandleFunc("/questions", webservice.GetQuestion).Methods("GET")
	//router.GET("/questions", webservice.GetQuestions)

	router.GET("/testService", webservice.TestService)
	//define other service here
	//log.Println("First log message!")

	router.Run(":8080")

}

//Getport() method gets the OS env variable for the port, if not available, it will se the default port
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
