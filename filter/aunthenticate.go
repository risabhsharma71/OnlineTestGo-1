package filter

import (
	"OnlineTestGo/daos/daoimpl"
	"log"
	//"github.com/gin-gonic/gin"
	"OnlineTestGo/utility"
	"time"
)

//var token string ="bsdhugcgdh"
func AuthenticateToken(token string) bool {
	utility.GetLogger()
	t1 := "Dec 29, 2014 at 6:00am (SGT)"
	layOut := "Jan 2, 2006 at 3:04am (MST)"
	timeStamp, err := time.Parse(layOut, t1)

	if err != nil {
		log.Println(err)

	}
	log.Println(timeStamp)
	hr, min, sec := timeStamp.Clock()
	log.Println(hr, min, sec)
	//return answer

	var tokenEncodeString string

	if len(token) > 0 {
		tokenEncodeString = token
	}
	log.Println(tokenEncodeString)

	tokenDao := daoimpl.TokenImpl{}

	log.Println("calling  daoimpl")
	tokenfromdb, lastaccesstime := tokenDao.AunthenticateToken(tokenEncodeString)
	log.Println(tokenfromdb, lastaccesstime)
	if len(tokenfromdb) == 0 {
		return false

	}
	currenttime := time.Now().Format("2006-01-02 15:04:05")
	log.Println(currenttime)
	const layOut1 = "2006-01-02 15:04:05"
	timeStamp1, err := time.Parse(layOut1, currenttime)
	if err != nil {
		log.Println(err)
		// os.Exit(1)
	}
	log.Println(timeStamp1)
	duration := timeStamp1.Sub(lastaccesstime)
	log.Println(duration)
	durationinhours := (float64)((duration / (1000 * 60 * 60) % 24))
	log.Println(durationinhours)
	if tokenEncodeString == tokenfromdb {
		//currenttime := time.Now().Format(time.RFC850)
		log.Println("calling modifyfunction")
		tokenDao.ModifyLastAccessTime(timeStamp1, tokenEncodeString)
		log.Println(timeStamp1)
	}
	if durationinhours > float64(hr) {

		log.Println("calling delete token function")
		bool := tokenDao.DeleteToken(tokenEncodeString)
		log.Println(bool)

	}
	return true
}
