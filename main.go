package main

import (
	"github.com/MIghtykukulkan/OnlineTestGo/webservice"
	"github.com/MIghtykukulkan/OnlineTestGo/models"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

func main() {

	router := gin.Default()

	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/userAnswer",webservice.AnswerList)
	router.GET("/userQuestion",webservice.QuestionList)
	//define other service here

	router.Run(":9090")
}
