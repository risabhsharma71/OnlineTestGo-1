package daoimpl

import (
	"fmt"

	"github.com/MIghtykukulkan/OnlineTestGo/models"
)

//QuestionImpl struct for Questions implementattion
type QuestionImpl struct{}

//GetQuesions fetched the questions from quesions table takes testId as input
func (dao QuestionImpl) GetQuesions(testid int64) []models.Question {
	var questions []models.Question
	db := connection()
	rows, err := db.Query("select * from questions where type = ?", testid)

	if err != nil {
		fmt.Print(err.Error())
		//return models.Question{}, err
	}

	for rows.Next() {
		var question models.Question
		err := rows.Scan(&question.ID, &question.Question, &question.Option1, &question.Option2, &question.Option3, &question.Option4, &question.Option5, &question.Answer, &question.Type)
		questions = append(questions, question)

		if err != nil {
			fmt.Print(err.Error())
			//return models.questions{},err
		}
	}
	defer rows.Close()

	return questions
}
