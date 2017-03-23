package daoimpl

import (
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"fmt"
	"log"
	"time"
)

type LoginImpl struct{}

func (dao LoginImpl) SaveToken(token models.Token) (int64, error) {
	utility.GetLogger()
	log.Println("exntering in SaveToken() function", token)
	log.Println("executing query and storing token in database for the user ")
	query := "insert into token(uid,token,lastaccesstime) values (?,?,?)"
	db, conn := connectaws()
	log.Println("token inserted")
	log.Println(token)

	defer db.Close()
	defer conn.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(token.Uid, token.Token, token.LastAccessTime)

	if err != nil {
		log.Panic("Exec err:", err.Error())
	}

	id, err := res.LastInsertId()
	log.Println("last inserted id:", id)
	if err != nil {
		log.Println("Exec err:", err.Error())
	}

	return id, err
}

func (dao LoginImpl) DeleteDuplicateUid(token models.Token) error {
	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()
	utility.GetLogger()
	log.Println("entering In DeleteDuplicateUid()")
	log.Println("executing query.....")
	query := "DELETE FROM token WHERE uid=?"

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(token.Uid)
	log.Println(res)
	val, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(val, err)
	return nil

}
func (dao LoginImpl) GetToken(uid int64, fname string, token string, lastaccestime time.Time) (int64, error) {
	query := "select  uid, token, lastaccesstime from token where  id= ?"

	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()

	rows, err := db.Query(query, token)
	if err != nil {
		log.Fatal(err)
		defer rows.Close()
	}

	for rows.Next() {

		err := rows.Scan(&token)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(token)
	}
	return 0, nil
}
