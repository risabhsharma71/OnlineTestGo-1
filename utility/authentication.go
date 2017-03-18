package utility

import (
	"fmt"

)

type TokenImpl struct{}


//var tokens models.Token
func (dao TokenImpl)  GetToken(token string)   bool {

fmt.Println(token)
	return true

}

