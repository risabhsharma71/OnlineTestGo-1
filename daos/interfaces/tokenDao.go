package interfaces

//tokenDao interface to be implemented by tokenimpl
type TokenDao interface {
	GetToken(token string)
}
