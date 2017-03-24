package interfaces

import (
	"OnlineTestGo/models"
	"OnlineTestGo/tos"
)

//QuestionDao interface to be implemented by Questionimpl
type QuestionDao interface {
	FetchQuestionsByType(testtype string) ([]tos.Question, error)
	AddQuestion(question models.Question) (int64, error)
	//AddOption(ID ,Options models.Question,userAnswer models,Question)(int64,error)
}
