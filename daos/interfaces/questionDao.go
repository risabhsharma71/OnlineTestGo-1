package interfaces

import "github.com/MIghtykukulkan/OnlineTestGo/models"

//QuestionDao interface to be implemented by userdaoimpl
type QuestionDao interface {
	GetQuesions(testid int64) []models.Question
	//other methods
}
