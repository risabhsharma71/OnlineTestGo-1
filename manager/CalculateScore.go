package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"
)

var questionDao interfaces.QuestionDao

func CalculateScore(answerList []models.Answer) int {

	utility.GetLogger()
	log.Println("entering manager.CalculateScore()")

	log.Println("calling Answer manager")

	questionDao := daoimpl.QuestionImpl{}
	answerDao := daoimpl.AnswerImpl{}

	score := 0

	log.Println("answerList", answerList)
	for _, answer := range answerList {

		log.Println("calling questionDao.GetAnswerById() and Fetching the right answer...")
		correctAnswer := questionDao.GetAnswerById(answer.Q_type)

		log.Println("actual vs correct", answer.Selected, correctAnswer)
		if answer.Selected == correctAnswer {

			score++

			answerDao.SaveAnswer(answer)
		}

	}

	return score

}
