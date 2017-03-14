package interfaces

import "OnlineTestGo/models"

//QuestionDao interface to be implemented by Questionimpl
type QuestionDao interface {
	FetchQuestionsByType(testtype string) ([]models.Question, error)
}
