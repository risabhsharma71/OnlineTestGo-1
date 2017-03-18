package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"	
     "log"
	"OnlineTestGo/utility"
)

var questionDao interfaces.QuestionDao

func CalculateScore(answerList []models.Answer) int {
	utility.GetLogger()
	log.Println("calling Answer manager")
	questionDao := daoimpl.QuestionImpl{}
	answerDao := daoimpl.AnswerImpl{}

	score := 0

	log.Println("aanswerList", answerList)
	for _, answer := range answerList {
		//checck if given answer is right
		log.Println("Fetching the right answer...")
		correctAnswer := questionDao.GetAnswerById(answer.Q_type)

		log.Println("actual vs correct", answer.Selected, correctAnswer)
		if answer.Selected == correctAnswer {
			//if right increment the score
			score++

			answerDao.SaveAnswer(answer)
		}

		//return the score

	}

	return score

}
