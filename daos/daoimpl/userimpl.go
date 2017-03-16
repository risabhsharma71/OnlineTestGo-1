package daoimpl

import (
	"log"

	"OnlineTestGo/models"
)

type UserImpl struct{}

func (dao UserImpl) SaveNewUser(user models.User) (int64, error) {

	query := "insert into registration(fname, lname, phone, email) values(?,?,?,?)"
	db := connection()
	defer db.Close()
	log.Println("Fname: ", user.Fname)
	log.Println("Lname: ", user.Lname)
	log.Println("phone: ", user.Phone)
	log.Println("email: ", user.Email)

	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Fname, user.Lname, user.Phone, user.Email)

	if err != nil {
		log.Panic("Exec err:", err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Exec err:", err.Error())
	}

	return id, err
}
