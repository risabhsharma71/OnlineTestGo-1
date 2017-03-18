package interfaces

import (
	"OnlineTestGo/models"

	"github.com/gin-gonic/gin"
)

//tokenDao interface to be implemented by tokenimpl
type TokenDao interface {
	GetToken(c *gin.Context, token models.Token)
}
