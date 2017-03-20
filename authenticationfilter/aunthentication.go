package aunthenticationfilter

import (
	
	"log"
	"OnlineTestGo/daos/daoimpl"
	//"github.com/gin-gonic/gin"
	"time"
	//"OnlineTestGo/models"
)



//var token string ="bsdhugcgdh"
func  authenticateToken(token string) {
	//return answer

	var tokenEncodeString string 

	if len(token)<=0 {
		tokenEncodeString = token
	}
	// /utility.GetLogger()
    log.Println("calling register manager")
    tokenDao := daoimpl.TokenImpl{}
	
    
//var uid int64=1;
    
    tokens, lastaccesstime :=  tokenDao.AunthenticateToken(tokenEncodeString)
    log.Println(lastaccesstime)
   // return insertedid
  // testtime:=time.Hour().Minute().Second()
  lastaccesstime1 := lastaccesstime.Format(time.RFC850)
  log.Println(lastaccesstime1)
   currenttime := time.Now().Format(time.RFC850)
   log.Println(currenttime)
	if tokenEncodeString == tokens {
	//currenttime := time.Now().Format(time.RFC850)
		//tokenDao.ModifyLastAccessTime(currenttime,tokenEncodeString)
		log.Println(currenttime)
	} 
	//else if(currenttime-lastaccesstime)>(testtime.){

//tokenDao.DeleteToken(token models.Token,uid)


	}
    
	



