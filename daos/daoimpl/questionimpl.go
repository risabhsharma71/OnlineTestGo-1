package daoimpl

import (
	"github.com/MIghtykukulkan/OnlineTestGo/models"
)

//QuestionImpl struct for Questions implementattion
type QuestionImpl struct{}

//GetQuesions fetched the questions from quesions table takes testId as input
func (dao QuestionImpl) GetQuesions(testid int64) []models.Question {

	var questions []models.Question

	//make DB calls and return quesion list

	return questions
}
