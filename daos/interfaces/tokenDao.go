package interfaces

import (
	"OnlineTestGo/models"
       "time"
	// /"github.com/gin-gonic/gin"
)

//tokenDao interface to be implemented by tokenimpl
type TokenDao interface {
	AunthenticateToken( token models.Token,uid int64) (string,time.Time)
}
