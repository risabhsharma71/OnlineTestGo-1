package utility

/*import (
	"OnlineTestGo/models"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type TokenImpl struct{}

var tokenEncodeString string = "xyzabc"

func (dao TokenImpl) GetToken(c *gin.Context, token models.Token) {

	db := connection()
	defer db.Close()
	token := ""
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
	if tokenEncodeString == token {
		return
	} else {

		c.JSON(code, gin.H{
			"code":    401,
			"message": "you dnt have permisssion to access",
		})
	}

}
*/
