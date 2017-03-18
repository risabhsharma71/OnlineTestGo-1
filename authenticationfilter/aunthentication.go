package aunthenticationfilter

import (
	
	//"log"
	"OnlineTestGo/daos/daoimpl"
	"github.com/gin-gonic/gin"
	"time"
)

type TokenImpl struct{}

var tokenEncodeString string = "xyzabc"
//var token string ="bsdhugcgdh"
func  authenticateToken(c *gin.Context) {
	//return answer
	utility.GetLogger()
    log.Println("calling register manager")
    tokenDao := daoimpl.TokenImpl{}
    
var uid int64=1;
    
    token, lastaccesstime := tokenDao.AunthenticateToken(token models.Token,uid)
    log.Println(lastaccesstime)
   // return insertedid
	if tokenEncodeString == token {
	currenttime := time.Now().Format(time.RFC850)
		tokenDao.ModifyLastAccessTime(token models.Token,currenttime)
		log.Println(currenttime)
	} else if(currenttime-lastaccesstime)>{

tokenDao.DeleteToken(token models.Token,uid)


	}else{

	  c.JSON(401, gin.H{
        "status":  "failure",
        "message":  "you dont have the acces ",
    })
	
    
	}

}

