package manager

import (
	"OnlineTestGo/daos/daoimpl"
)

func DeleteToken(delt string) string {

	logoutdao := daoimpl.LogoutImpl{}
	a := logoutdao.DeleteToken(delt)

	return a
}
