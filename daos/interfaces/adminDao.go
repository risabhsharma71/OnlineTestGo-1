package interfaces

import "OnlineTestGo/models"

//adminDao interface to be implemented adminimpl
type AdminDao interface {
	Fetchdata() ([]models.Admin,error)
}