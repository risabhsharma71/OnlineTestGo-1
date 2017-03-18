package daoimpl

import (
	"OnlineTestGo/models"
	"fmt"
	"log"
)

type LoginImpl struct{}

func (dao LoginImpl) chcktoken(token models.Token) (int64, error) {
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
