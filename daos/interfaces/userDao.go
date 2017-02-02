package interfaces

import "github.com/MIghtykukulkan/OnlineTestGo/models"

//UserDao interface to be implemented by userdaoimpl
type UserDao interface {
	SaveNewUser(u *models.User) int
}
