package manager

import (
	"log"
       
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
)

 var userDao interfaces.UserDao
 var typeDao interfaces.TypeDao
var totalquestionDao interfaces.TotalQuestionDao
  //var typeDao interfaces.TypeDao
  //var questionDao interfaces.QuestionDao


//Register manager takes care of business logic like calling daos
func Register(user models.User) int64 {
	log.Println("calling register manager")

	userDao = daoimpl.UserImpl{}
	
	//insert user info first
	insertedid, err := userDao.SaveNewUser(user)
	if err != nil {
		log.Println("error occured", err)
	}
	log.Println(insertedid)
	return insertedid
}

func Answer(answer models.Answer) string {
	log.Println("calling Answer manager")

	//var answers []models.Answer
	answerDao := daoimpl.AnswerImpl{}

	//insert Answers of user
	insertedid, err := answerDao.SaveAnswer(answer)
	if err != nil {
		log.Println("error occured", err)
	}
	log.Println(insertedid)
	 
        return ("succesful")       
	
}
/*

//func FetchQuestion(question models.Question,testtype string) []models.Question {

	log.Println("calling Question manager")
        totalquestionDao := daoimpl.TotalQuestionImpl{}
        questionlist := totalquestionDao.GetQuesions(testtype)

	//insert Question of user

	return questionlist

}
*/