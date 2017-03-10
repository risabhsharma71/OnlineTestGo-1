package interfaces

import "OnlineTestGo/models"

//QuestionDao interface to be implemented by userdaoimpl
type TotalQuestionDao interface {
	GetQuesions(testtype string) []models.TotalQuestion
	//other methods
}
