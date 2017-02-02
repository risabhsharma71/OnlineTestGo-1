package webservice

import (

	//"bytes"
	//"database/sql"
	"fmt"

	"github.com/MIghtykukulkan/OnlineTestGo/manager"
	"github.com/gin-gonic/gin"
	//"net/http"
	"github.com/MIghtykukulkan/OnlineTestGo/models"
)

func RegisterUser(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)
	//function should be calling this manager class

	questionlist := manager.Register(user)

	//println statements should be replaced with logs
	fmt.Println("questionlist", questionlist)

	//response should return theh question and answer list
	c.JSON(200, gin.H{
		"status":  "success",
		"message": questionlist,
	})
}
