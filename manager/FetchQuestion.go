package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
	"log"

	"OnlineTestGo/utility"
)

var questionDao1 interfaces.QuestionDao

func FetchQuestion(testtype string) []models.Question {
	utility.GetLogger()
	log.Println("entering into manager.FetchQuestion")
	questionDao1 := daoimpl.QuestionImpl{}
	log.Println("calling FetchQuestionsByType()")
	return questionDao1.FetchQuestionsByType(testtype)

}
