package daoimpl

import (
	"fmt"
	"log"
)

type LoginImpl struct{}

func (dao LoginImpl) SaveToken(token string, uid int64) (int64, error) {
	query := "insert str into Token where uid= ?"
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

func (dao LoginImpl) GetToken(token string, uid int64) (int64, error) {
	query := "select token from token where uid= ?"
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
