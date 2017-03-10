package daoimpl

import (
	"fmt"
	//"log"
	"OnlineTestGo/models"
)

//QuestionImpl struct for Questions implementattion
type QuestionImpl struct{}

//GetQuestions fetched the questions from quesions table takes testId as input
func (dao QuestionImpl) GetQuesions(testtype string) []models.Question {
var questions []models.Question
       var question models.Question
                var option models.Options
       var questionlist []string
       var totalquestions []models.TotalQuestion
       var choice  []models.Options

	db := connection()
	rows, err := db.Query("select Q.id,Q.question,Q.type,O.choices from questions Q,Options O where Q.id=O.qid and type=?",testtype)
       
	fmt.Println(testtype)
	if err != nil {
		fmt.Print(err.Error())
		//return models.Question{}, err
	}

	for rows.Next() {
		
	        err := rows.Scan(&question.ID,&question.Question)
                if(&question.Question==&questionlist.Question){
                   for rows.Next(){
                         rows.Scan(&option.Choices)
                     rows = append(,&option)
                          

}
}
                 

                   
