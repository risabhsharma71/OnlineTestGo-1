package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/models"
)

func FetchData() []models.Admin {
	adminDao := daoimpl.AdminImpl{}
	return adminDao.FetchData()

}
