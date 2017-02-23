package daoimpl

import (
	"fmt"
	"log"
	"OnlineTestGo/models"
)

//QuestionImpl struct for Questions implementattion
type QuestionImpl struct{}

//GetQuestions fetched the questions from quesions table takes testId as input
func (dao QuestionImpl) GetQuesions(testtypes string) []models.Question {
	var questions []models.Question
	db := connection()
	rows, err := db.Query("select * from questions where type='java'")
	fmt.Println(testtypes)
	if err != nil {
		fmt.Print(err.Error())
		//return models.Question{}, err
	}

	for rows.Next() {
		var question models.Question
		err := rows.Scan(&question.ID, &question.Question,&question.Type)
		questions = append(questions, question)
		
		if err != nil {
			fmt.Print(err.Error())
			//return models.questions{},err
		}
		log.Println(question.ID,question.Question)
	}
	defer rows.Close()

	return questions
}
