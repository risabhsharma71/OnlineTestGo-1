package daoimpl

import (
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"fmt"
	"log"
)

type UserImpl struct{}

func (dao UserImpl) SaveNewUser(user models.User) (int64, error) {

	utility.GetLogger()

	log.Println("entering in SaveNewUser() function")
	log.Println("executing query and storing registration details")
	query := "insert into registration(fname, lname, phone, email,password,usertype) values(?,?,?,?,?,?)"
	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Fname, user.Lname, user.Phone, user.Email, user.Password, user.UserType)

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
	utility.GetLogger()
	log.Println("entering in userDao.CheckUser()...")
	var id int64
	phone := user.Phone
	log.Println("executing query and checking user exists")
	query := "select phone from registration where phone = ?"
	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()

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

func (dao UserImpl) AuthenticateUser(user models.User) models.User {

	utility.GetLogger()

	log.Println("entering in AuthenticateUser()")

	log.Println("Executing query and authenticating user exist")
	query := "select id, fname, lname, phone, usertype from registration where email=? and password = ?"
	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()

	rows, err := db.Query(query, user.Email, user.Password)

	if err != nil {
		log.Fatal(err)
		defer rows.Close()
	}

	var newuser models.User
	for rows.Next() {

		err := rows.Scan(&newuser.ID, &newuser.Fname, &newuser.Lname, &newuser.Phone, &newuser.UserType)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(newuser)
	}

	log.Println("Response user Obj : ", user)

	return newuser
}
