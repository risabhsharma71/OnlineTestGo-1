package daoimpl

import "OnlineTestGo/models"
import "fmt"

type LoginImpl struct{}

func (dao LoginImpl) authenticateUser(user models.User) (int64, error) {

	fmt.Println(user)

	return 0, nil

}
