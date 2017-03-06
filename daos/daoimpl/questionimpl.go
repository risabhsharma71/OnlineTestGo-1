package daoimpl

import (
	"fmt"
	"log"
	"OnlineTestGo/models"
)

//QuestionImpl struct for Questions implementattion
type QuestionImpl struct{}

//GetQuestions fetched the questions from quesions table takes testId as input
func (dao QuestionImpl) GetQuesions(testtype string) []models.Question {
        var totalquestions []models.TotalQuestion
	var questions []models.Question
        var choice  []models.Options
	db := connection()
	rows, err := db.Query("select * from questions where type=?",testtype)
       // rows1,err:db.Query("select choices from Options where qid=?",&question.ID)
	fmt.Println(testtype)
	if err != nil {
		fmt.Print(err.Error())
		//return models.Question{}, err
	}

	for rows.Next() {
		var question models.Question
		err := rows.Scan(&question.ID, &question.Question,&question.Type)
                 rows1:db.Query("select choices from Options where qid=?",&question.ID)
                for rows.Next(){
                 var option models.Options
                 rows.Scan(&option.Choices)
                 choice= append(choice, option)
                   
                 




}
		questions = append(questions, question)
		totalquestions=append(questions, choice)
		if err != nil {
			fmt.Print(err.Error())
			//return models.questions{},err
		}
		log.Println(question.ID,question.Question)
	}
	defer rows.Close()

	return questions
}
