package utility

import (
	"OnlineTestGo/models"
	"database/sql"
	"log"
	"OnlineTestGo/daos/daoimpl"
	"github.com/gin-gonic/gin"
)

type TokenImpl struct{}

var tokenEncodeString string = "xyzabc"

func (dao TokenImpl) GetToken(c *gin.Context, token models.Token) {

	db := connection()
	defer db.Close()
	tokens :=token.Token
	err := db.QueryRow("select token from rpqbmysql.Token where uid=1").Scan(&token)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No rows with that token.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Token", token)
	}

	log.Println("tokenid", token)
	//return answer
	if tokenEncodeString == tokens {
		return
	} else {

	  c.JSON(401, gin.H{
        "status":  "failure",
        "message":  "you dont have the acces ",
    })
    
	}

}

