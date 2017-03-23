package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"
)

func FetchData() []models.Admin {
	utility.GetLogger()
	log.Println("entering manager.FetchData()")
	adminDao := daoimpl.AdminImpl{}
	log.Println("calling adminDao.FetchData()")

	return adminDao.FetchData()

}
