package utility

import (
    "database/sql"
    "log"
    
     "github.com/gin-gonic/gin"
)
type TokenImpl struct{}
var tokenEncodeString string = "xyzabc"
//var tokens models.Token
func (dao TokenImpl)  GetToken(c *gin.Context,token models.Token)   {
	db := connection()
	defer db.Close()
	token := ""
	//lastaccesstime:=""
	err := db.QueryRow("select token from rpqbmysql.Token where uid=1").Scan(&token)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No rows with that token.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Token", token)
	}

	log.Println("tokenid",token)
	//return answer
 if (tokenEncodeString == token) {
             /* query := "INSERT INTO token (lastaccesstime) VALUES (?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()
var currenttime = time.Now()
currenttime.Format(time.RFC3339)
	res, err := stmt.Exec(currenttime)
	log.Println(res)
    
	//id, err := res.LastInsertId()
	if err != nil {
		log.Println("Exec err:", err.Error())
	}*/
	
				  return
}else{
		

			c.JSON(code, gin.H{
				"code":    401,
				"message": "you dnt have permisssion to access",
			})
		}
	
}