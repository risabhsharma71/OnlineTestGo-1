package interfaces

import "net"

//TypeDao interface to be implemented by userdaoimpl
type SqlConnDao interface {
	Dial(addr string) (net.Conn, error)
	//other method defnitions
}
