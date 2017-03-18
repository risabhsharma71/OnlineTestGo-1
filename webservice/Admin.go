package webservice

import (
	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	dlist := manager.FetchData()
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status": "sucess",
		"score":  dlist,
	})

}
