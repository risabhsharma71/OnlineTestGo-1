package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"
)

func FetchQuestion(testtype string) []models.Question {

	utility.GetLogger()
	log.Println("entering into manager.FetchQuestion")

	log.Println("calling Question manager")

	questionDao1 := daoimpl.QuestionImpl{}
	log.Println("calling FetchQuestionsByType()")
	return questionDao1.FetchQuestionsByType(testtype)

}
