package interfaces

//tokenDao interface to be implemented by tokenimpl
type LogoutDao interface {
	DeleteToken(token string) string
}
