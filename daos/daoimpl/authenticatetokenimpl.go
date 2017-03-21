package daoimpl

import (
	"database/sql"
	"log"
	"time"
	//  "OnlineTestGo/models"
	"OnlineTestGo/utility"
)

//TypeImpl is a implementation for type interface
type TokenImpl struct{}

//GetIdfromType returns the id from Type table for input value
func (dao TokenImpl) ModifyLastAccessTime(currentime time.Time, tokenEncodeString string) error {
	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()
	utility.GetLogger()
	log.Println("In modifyfunction")
	query := "update  token set lastaccesstime=? where token=? "
	log.Println("after query execution")

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()
	log.Println("reached to modifyexecution")
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
	//var token string
	//var lastaccesstime time.Time
	utility.GetLogger()
	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()
	token := ""
	lastaccesstime := ""
	log.Println("in AunthenticateToken function")
	err := db.QueryRow("select token,lastaccesstime from token where token=?", tokenEncodeString).Scan(&token, &lastaccesstime)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No uid with that token.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Type is %v\n", token)
	}
	log.Println(lastaccesstime)
	const layOut1 = "2006-01-02 15:04:05"
	timeStamp1, err := time.Parse(layOut1, lastaccesstime)
	log.Println(timeStamp1)
	return token, timeStamp1
}
func (dao TokenImpl) DeleteToken(deletetoken string) bool {
	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()
	//var token models.Token
	utility.GetLogger()
	log.Println("calling DeleteToken function")
	query := "delete  from token where token=?"
	log.Println("after query execution")

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()
	log.Println("reached to execution")
	res, err := stmt.Exec(deletetoken)
	log.Println(res)
	//val, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	//log.Println(val)
	return true

}
