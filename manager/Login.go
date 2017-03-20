package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/models"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func AuthenticateUser(user models.User) tos.Tokento {

	var tokenTo tos.Tokento
	utility.GetLogger()
	log.Println("calling Login manager")
	token := GenerateToken()
	tokenDao := daoimpl.TokenImpl{}

	vsdd := tokenDao.GetToken(token, uid)

	//copy values fron user and generate token method to tokento and return it
	return tokenTo
}

func GenerateToken() string {
	str := RandStringRunes(20)
	{
		fmt.Println(str)

	}
	return str
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	token := make([]rune, n)
	for i := range token {
		token[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(token)
}
