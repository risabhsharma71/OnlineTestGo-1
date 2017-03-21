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
	log.Println("exntering in SaveToken() function")
	log.Println("executing query and storing token in database for the user ")
	query := "insert  into Token(uid,token,lastaccesstime) values (?,?,?)"
	db, conn := connectaws()
	fmt.Println("token inserted")
	fmt.Println(token)

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
	if err != nil {
		log.Println("Exec err:", err.Error())
	}

	return id, err
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
