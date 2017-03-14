package daoimpl

import (
	"OnlineTestGo/models"
	"database/sql"
	"log"
        "OnlineTestGo/utility"
)



type QuestionImpl struct{}

// FetchQuestionsByType fetchs the questions from DB by merging quesion and options table
func (dao QuestionImpl) FetchQuestionsByType(testtype string) []models.Question {
	var totalquestions []models.TotalQuestion
	var questionList []models.Question
        utility.GetLogger()
	query := "SELECT A.id, A.question, B.choices FROM rpqbmysql.questions as A right join rpqbmysql.Options as B on A.id = B.qid where type = ?"

	db := connection()
	defer db.Close()

	rows, err := db.Query(query, testtype)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var totalquestion models.TotalQuestion
		err = rows.Scan(&totalquestion.ID, &totalquestion.Question, &totalquestion.Choices)

		if err != nil {
			log.Fatal(err)
		}

		totalquestions = append(totalquestions, totalquestion)

	}

	log.Println("totalquestions:", totalquestions)

	intitialID := 0
	var options []string
	var question models.Question
	for _, questionRow := range totalquestions {

		if intitialID != questionRow.ID {
			question.ID = questionRow.ID
			question.Question = questionRow.Question
			questionList = append(questionList, question)
			intitialID = questionRow.ID
		}

	}

	n := 0
	intitialID = totalquestions[0].ID
	for i := 0; i < len(totalquestions); i++ {

		if intitialID != totalquestions[i].ID {
			questionList[n].Options = options
			n++
			options = make([]string, 0)
			intitialID = totalquestions[i].ID
		}

		options = append(options, totalquestions[i].Choices)

	}
	questionList[n].Options = options

	return questionList

}

func (dao QuestionImpl) GetAnswerById(ID int64) string {
utility.GetLogger()
	db := connection()
	defer db.Close()
	answer := ""
	err := db.QueryRow("select answers from rpqbmysql.Options where answers != '0' && qid=?", ID).Scan(&answer)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No rows with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Type is %v\n", answer)
	}

	log.Println("correct Answer fir ID", ID, answer)
	return answer
}
