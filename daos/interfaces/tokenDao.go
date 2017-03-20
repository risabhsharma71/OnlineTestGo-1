package interfaces

import (
	"OnlineTestGo/models"
	"time"
	// /"github.com/gin-gonic/gin"
)

//tokenDao interface to be implemented by tokenimpl
type TokenDao interface {
	GetToken(token string, uid int64)
	AunthenticateToken(tokenEncodeString string) (string, time.Time)
	ModifyLastAccessTime(currentime time.Time, tokenEncodeString string) error
	DeleteToken(token models.Token, uid int64) error
}
