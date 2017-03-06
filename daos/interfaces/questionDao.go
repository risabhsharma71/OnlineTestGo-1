package interfaces

import "OnlineTestGo/models"

//QuestionDao interface to be implemented by userdaoimpl
type QuestionDao interface {
	GetQuesions(testtype string) []models.Question
	//other methods
}
