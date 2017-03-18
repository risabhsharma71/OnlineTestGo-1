package daoimpl

import (
	"database/sql"
	"log"
    "time"
)

//TypeImpl is a implementation for type interface
type TokenImpl struct{}

//GetIdfromType returns the id from Type table for input value
func (dao TokenImpl) ModifyLastAccessTime(token models.Token,currentime time.Time) error{
	db := connection()
	defer db.Close()

	utility.GetLogger()
	log.Println("calling  function")
	query := "INSERT INTO token (lastaccesstime) VALUES(?)"
	log.Println("after query execution")

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()
	log.Println("reached to execution")
	res, err := stmt.Exec(currentime )
	log.Println(res)
	val, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(val, err)
	return nil

}
func (dao TokenImpl) AunthenticateToken(token models.Token)(uid int64) (string,string) {
	var token strin(g
	var lastaccesstime time.Time
	db := connection()

	err := db.QueryRow("select token,lastaccesstime from token where uid=?", uid).Scan(&token,&lastaccesstime)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No uid with that token.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Type is %v\n", token)
	}

	
	return token,lastaccesstime
}
func (dao TokenImpl) DeleteToken(token models.Token) error{
	db := connection()
	defer db.Close()

	utility.GetLogger()
	log.Println("calling  function")
	query := "INSERT INTO token (lastaccesstime) VALUES(?)"
	log.Println("after query execution")

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()
	log.Println("reached to execution")
	res, err := stmt.Exec(currentime )
	log.Println(res)
	val, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(val, err)
	return nil

}