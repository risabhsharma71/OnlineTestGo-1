package webservice

import (
	"fmt"

	"log"

	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"

	"OnlineTestGo/models"
)

func AnswerList(c *gin.Context) {

	//Answer list recieves a list of answers
	var answer []models.Answer
	c.BindJSON(&answer)
	log.Println(answer)
	score := manager.CalculateScore(answer)

	fmt.Println("answerlist", score)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{

		"status": "success",

		"score": score,
	})
}
