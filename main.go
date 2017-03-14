package main

import (
	   "OnlineTestGo/webservice"
	   "OnlineTestGo/utility"
           "log"
          
	   "github.com/goinggo/tracelog"
           "fmt"
	   "github.com/gin-gonic/gin"

	
)


func main() {
utility.GetLogger()
   tracelog.Start(tracelog.LevelTrace)
    tracelog.Trace("main", "main", " Trace")
    tracelog.Info("main", "main", " Info")
    tracelog.Warning("main", "main", " Warn")
    tracelog.Errorf(fmt.Errorf("Exception At..."), "main", "main", " Error")

	router := gin.Default()
        

	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/userAnswer", webservice.AnswerList)
	router.GET("/questions", webservice.QuestionList)
	//router.HandleFunc("/questions",webservice.GetQuestion).Methods("GET")
	//router.GET("/questions", webservice.GetQuestions)

	router.GET("/testService", webservice.TestService)
	//define other service here
log.Println("First log message!")
	router.Run(":8081")




 


}
