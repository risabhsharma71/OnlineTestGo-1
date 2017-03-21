package daoimpl

import (
	"database/sql"
	"log"
	"time"
	//  "OnlineTestGo/models"
	"OnlineTestGo/utility"
)

type TokenImpl struct{}

func (dao TokenImpl) ModifyLastAccessTime(currentime time.Time, tokenEncodeString string) error {
	db := connection()
	defer db.Close()

	utility.GetLogger()
	log.Println("entering In ModifyLastAccessTime()")
	log.Println("executing query updating lastaccesstime to currentime")
	query := "update  token set lastaccesstime=? where token=? "

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(currentime, tokenEncodeString)
	log.Println(res)
	val, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(val, err)
	return nil

}
func (dao TokenImpl) AunthenticateToken(tokenEncodeString string) (string, time.Time) {

	utility.GetLogger()
	db := connection()
	token := ""
	lastaccesstime := ""
	log.Println("entering in AunthenticateToken() function")
	log.Println("executing query fetching and returning token,lastaccesstime")
	err := db.QueryRow("select token,lastaccesstime from token where token=?", tokenEncodeString).Scan(&token, &lastaccesstime)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No uid with that token.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Type is %v\n", token)
	}

	const layOut1 = "2006-01-02 15:04:05"
	timeStamp1, err := time.Parse(layOut1, lastaccesstime)

	return token, timeStamp1
}
func (dao TokenImpl) DeleteToken(deletetoken string) bool {
	db := connection()
	defer db.Close()

	utility.GetLogger()
	log.Println("entering DeleteToken() function")
	log.Println("executing query and delete token ")
	query := "delete  from token where token=?"

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(deletetoken)
	log.Println(res)

	if err != nil {
		log.Println(err)
	}

	return true

}
