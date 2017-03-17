package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
	"fmt"
	"log"

	"OnlineTestGo/utility"
)

var userDao interfaces.UserDao

//Register manager takes care of business logic like calling daos

func Register(user models.User) int64 {
	utility.GetLogger()
	log.Println("calling register manager")
	userDao = daoimpl.UserImpl{}
	id, err := userDao.CheckUser(user)

	if id != 0 {
		return id
		fmt.Println("checked phone number")
	}

	//insert user info first
	insertedid, err := userDao.SaveNewUser(user)
	if err != nil {
		log.Println("error occured", err)
	}
	log.Println(insertedid)

	return insertedid
}
