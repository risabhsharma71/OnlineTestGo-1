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

func Login(user models.User) tos.Tokento {
	utility.GetLogger()
	log.Println("entering login manager...")
	var tokenObj tos.Tokento
	var tokenModel models.Token

	//check user esxists or not
	userDao := daoimpl.UserImpl{}
	tokenDao := daoimpl.LoginImpl{}

	userObj := userDao.AuthenticateUser(user)

	if userObj.ID != 0 {
		//genereate a token for the user and save it in DB
		token := GenerateToken()

		//insert token to table
		tokenModel.LastAccessTime = time.Now()
		tokenModel.Token = token
		tokenModel.Uid = userObj.ID

		id, err := tokenDao.SaveToken(tokenModel)
		if err != nil {
			log.Println(err)
		}

		//copy the valuuse to tokenObj if insertion intoken table is successful
		if id == 0 {
			tokenObj.Token = token
			tokenObj.Fname = userObj.Fname
			tokenObj.Uid = userObj.ID
			tokenObj.UserType = userObj.UserType
		}

		log.Println(tokenObj)

	}

	//return the tokenObj back
	return tokenObj

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
