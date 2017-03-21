package daoimpl

import (
	"OnlineTestGo/models"
	"database/sql"
	"log"
	"time"
	//"OnlineTestGo/utility"
)

//TypeImpl is a implementation for type interface
type TokenImpl struct{}

//GetIdfromType returns the id from Type table for input value
func (dao TokenImpl) ModifyLastAccessTime(currentime time.Time, tokenEncodeString string) error {
	db := connection()
	defer db.Close()

	//	utility.GetLogger()
	log.Println("calling  function")
	query := "update  token set lastaccesstime=? where token=? "
	log.Println("after query execution")

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()
	log.Println("reached to execution")
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
	db := connection()
	var token models.Token
	err := db.QueryRow("select token,lastaccesstime from token where token=?", tokenEncodeString).Scan(&token.Token, &token.LastAccessTime)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No uid with that token.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Type is %v\n", token)
	}

	tokens := token.Token
	lastaccesstime := token.LastAccessTime
	return tokens, lastaccesstime
}
func (dao TokenImpl) DeleteToken(deletetoken string) error {
	db := connection()
	defer db.Close()
	//var token models.Token
	//	utility.GetLogger()
	log.Println("calling  function")
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
	val, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(val, err)
	return nil

}
