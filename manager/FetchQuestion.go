package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
	"log"
	//	"OnlineTestGo/utility"
)

var questionDao1 interfaces.QuestionDao

func FetchQuestion(testtype string) []models.Question {
	//	utility.GetLogger()
	log.Println("calling Question manager")
	questionDao1 := daoimpl.QuestionImpl{}
	return questionDao1.FetchQuestionsByType(testtype)

}
