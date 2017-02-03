package manager

import (
	"log"

	"github.com/MIghtykukulkan/OnlineTestGo/daos/daoimpl"
	"github.com/MIghtykukulkan/OnlineTestGo/models"
)

//Register manager takes care of business logic like calling daos
func Register(user models.User) []models.Question {
	log.Println("calling register manager")

	var questions []models.Question
	userDao := daoimpl.UserImpl{}
	//typeDao := daoimpl.TypeImpl{}
	//questionDao := daoimpl.QuestionImpl{}

	//insert user info first
	insertedid, err := userDao.SaveNewUser(user)
	if err != nil {
		log.Println("error occured", err)
	}
	log.Println(insertedid)

	//get the type ID
	//testID := typeDao.GetIdfromType(user.Test)

	//now get the questions from Questions table
	//questions := questionDao.GetQuesions(testID)

	return questions
}

func Answer(answer models.Answer) {
	log.Println("calling Answer manager")

	var answers []models.Answer
	answerDao := daoimpl.AnswerImpl{}
	
	//insert Answers of user
	insertedid, err := answerDao.SaveAnswer(answer)
	if err != nil {
		log.Println("error occured", err)
	}
	log.Println(insertedid)
	return "successful"
}
