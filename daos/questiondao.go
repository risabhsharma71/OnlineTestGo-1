package daos

import "fmt"

func GetquestionList(testType string) []Question {

	var questions []Question
	db := Connection()

	rows, err := db.Query("select * from questions where Type = ? ", testType)
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		var question Question
		err = rows.Scan(&question.Id, &question.Question, &question.Option1, &question.Option2, &question.Option3, &question.Option4, &question.Option5, &question.Answer, &question.Type)
		questions = append(questions, Question{question.Id, question.Question, question.Option1, question.Option2, question.Option3, question.Option4, question.Option5, question.Answer, question.Type})

		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return questions

}
