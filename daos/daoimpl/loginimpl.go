package daoimpl

import (
	"OnlineTestGo/models"
	"fmt"
	"log"
	"time"
)

type LoginImpl struct{}

func (dao LoginImpl) SaveToken(token models.Token) (int64, error) {
	query := "insert  into Token(uid,token,lastaccesstime) values (?,?,?)"
	db := connection()
	defer db.Close()

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
	db := connection()
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
