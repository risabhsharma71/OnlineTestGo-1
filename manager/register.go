package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
	"log"
	"fmt"
       )

var userDao interfaces.UserDao


//var typeDao interfaces.TypeDao
var questionDao interfaces.QuestionDao
//var typeDao interfaces.TypeDao
//var questionDao interfaces.QuestionDao

//Register manager takes care of business logic like calling daos
func Register(user models.User) int64 {
	log.Println("calling register manager")
	 userDao = daoimpl.UserImpl{}
     		id:=userDao.checkUser(user)
         
               if id!=0{
		return id
                fmt.Println("checked phone number")
                       }
         
	//insert user info first
	insertedid, err := userDao.SaveNewUser(user)
          if err != nil {
		log.Println("error occured", err)
	}
	log.Println(insertedid)
	
	return insertedid
       }



func CalculateScore(answerList []models.Answer) int {
	log.Println("calling Answer manager")
	questionDao := daoimpl.QuestionImpl{}
	answerDao := daoimpl.AnswerImpl{}

	score := 0

	log.Println("aanswerList", answerList)
	for _, answer := range answerList {
		//checck if given answer is right
		log.Println("Fetching the right answer...")
		correctAnswer := questionDao.GetAnswerById(answer.Q_type)

		log.Println("actual vs correct", answer.Selected, correctAnswer)
		if answer.Selected == correctAnswer {
			//if right increment the score
			score++

			answerDao.SaveAnswer(answer)
		}

		//return the score

	}

	return score

}

func FetchQuestion(testtype string) []models.Question {
	questionDao := daoimpl.QuestionImpl{}
	return questionDao.FetchQuestionsByType(testtype)

}
func FetchData() []models.Admin {
	adminDao := daoimpl.AdminImpl{}
	return adminDao.FetchData()

}