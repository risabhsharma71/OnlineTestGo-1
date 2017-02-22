package interfaces

import "OnlineTestGo/models"

//UserDao interface to be implemented by userdaoimpl
type UserDao interface {
	SaveNewUser(u models.User) (int64, error)
}
