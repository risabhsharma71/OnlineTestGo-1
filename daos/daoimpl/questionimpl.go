package daoimpl

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)	

//QuestionImpl struct for Questions implementattion
type QuestionImpl struct{}

//GetQuesions fetched the questions from quesions table takes testId as input
func (dao QuestionImpl) GetQuesions(testid int64) []models.Question {
	var questions []models.Question
	rows, err := db.Query("select * from questions")
	
	if err != nil {
		fmt.Print(err.Error())
		//return models.Question{}, err
	}

	for rows.Next(){
		var question Question
		err := rows.Scan(&question.Id, &question.Question, &question.Opton1, &question.Option2, &question.Option3, &question.Option4, &question.Option5, &question.Answer, &question.Type)
		questions = append(questions,Question{question.Id, question.Question, question.Opton1, question.Option2, question.Option3, question.Option4, question.Option5, question.Answer, question.Type})
		
		if err != nil {
			fmt.Print(err.Error())
			//return models.questions{},err
		}
	}
	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"result": questions,
		"count": len(questions),
	})
	//make DB calls and return quesion list

	return questions
}
