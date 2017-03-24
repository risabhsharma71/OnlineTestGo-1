package interfaces

import "OnlineTestGo/models"

type UserDao interface {
	SaveNewUser(u models.User) (int64, error)
	CheckUser(u models.User) (int64, error)
	AuthenticateUser(user models.User) models.User
}
