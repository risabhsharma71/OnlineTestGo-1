package main

import
//"bytes"
//"database/sql"
//"fmt"
//"net/http"

(
	"fmt"

	"github.com/MIghtykukulkan/OnlineTestGo/manager"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/registerUser", registerUser)

	router.Run(":9090")
}

func registerUser(c *gin.Context) {

	//function should be calling this manager class
	questionlist := manager.Register()

	//println statements should be replaced with logs
	fmt.Println("questionlist", questionlist)

	//response should return theh question and answer list
	c.JSON(200, gin.H{
		"status":  "success",
		"message": questionlist,
	})
}
