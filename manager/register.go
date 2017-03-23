package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"
)

var userDao interfaces.UserDao

func Register(user models.User) int64 {

	utility.GetLogger()
	log.Println("entering in  manager.Register() function")

	log.Println("calling register manager")
	userDao := daoimpl.UserImpl{}
	log.Println("calling userDao.CheckUser function")
	id, err := userDao.CheckUser(user)
	{
		if id != 0 {
			return id
			log.Println(id)
		}
	}

	user.UserType = "user"
	log.Println("calling userDao.SaveNewUser() function")
	insertedid, err := userDao.SaveNewUser(user)
	if err != nil {
		log.Println("error occured", err)
	}
	log.Println(insertedid)

	return insertedid
}
