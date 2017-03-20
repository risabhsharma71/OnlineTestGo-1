package interfaces


import (
	// /"OnlineTestGo/models"
       "time"
	// /"github.com/gin-gonic/gin"
)

//tokenDao interface to be implemented by tokenimpl
type TokenDao interface {
	GetToken(token string)
	AunthenticateToken(tokenEncodeString string) (string,time.Time)
	ModifyLastAccessTime(currentime time.Time,tokenEncodeString string) error
	DeleteToken(tokenEncodeString string) (bool)

}
