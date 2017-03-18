package interfaces


import (
	"OnlineTestGo/models"
       "time"
	// /"github.com/gin-gonic/gin"
)

//tokenDao interface to be implemented by tokenimpl
type TokenDao interface {
	GetToken(token string)
	AunthenticateToken( token models.Token,uid int64) (string,time.Time)
	ModifyLastAccessTime(token models.Token,currentime time.Time) error
	DeleteToken(token models.Token,uid int64) error
}
