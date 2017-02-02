package interfaces

//TypeDao interface to be implemented by userdaoimpl
type TypeDao interface {
	GetIdfromType(s string) int64
	//other method defnitions
}
