package daoimpl

import (
	"OnlineTestGo/models"
	"fmt"
	"log"
)

type UserImpl struct{}

func (dao UserImpl) SaveNewUser(user models.User) (int64, error) {

	query := "insert into registration(fname, lname, phone, email,password) values(?,?,?,?,?)"
	db := connection()
	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Fname, user.Lname, user.Phone, user.Email,user.Password)

	if err != nil {
		log.Panic("Exec err:", err.Error())
	}

	id, err := res.LastInsertId()

	fmt.Printf("%T", id)
	if err != nil {
		log.Println("Exec err:", err.Error())
	}

	return id, err
}

func (dao UserImpl) CheckUser(user models.User) (int64, error) {

	var id int64
	phone := user.Phone
	query := "select phone from registration where phone = ?"
	db := connection()

	rows, err := db.Query(query, phone)
	if err != nil {
		log.Fatal(err)
		defer rows.Close()
	}

	for rows.Next() {

		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}

	return id, err
}