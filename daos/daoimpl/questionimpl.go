package daoimpl

import (
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"database/sql"
	"log"
)

type QuestionImpl struct{}

func (dao QuestionImpl) FetchQuestionsByType(testtype string) []models.Question {
	var totalquestions []models.TotalQuestion
	var questionList []models.Question
	utility.GetLogger()
	log.Println("entering into FetchQuestionsByType()")
	log.Println("executing query and Fetching Questions By Type ")
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
	log.Println("entering GetAnswerById() ")
	db := connection()
	defer db.Close()
	answer := ""
	log.Println("executing query")
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

func addOption(id int64, options string, answer string) error {
	db := connection()
	defer db.Close()

	utility.GetLogger()
	log.Println("entering into addoption function")
	log.Println("executing query and storing options to corresponding questions into db")
	query := "INSERT INTO Options (qid,choices,answers) VALUES(?,?,?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id, options, answer)
	log.Println(res)
	val, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(val, err)
	return nil

}

func (dao QuestionImpl) AddQuestion(question models.Question) (int64, error) {
	db := connection()
	defer db.Close()

	utility.GetLogger()
	log.Println("entering in AddQuestion() function")
	log.Println("executing query and storing questions into database")
	query := "INSERT INTO questions (question,type) VALUES (?, ?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(question.Question, question.Type)
	log.Println(res)

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Exec err:", err.Error())
	}
	log.Println(id)

	Options := question.Options

	log.Println("calling addoption function")
	for i := 0; i < len(Options); i++ {

		addOption(id, Options[i], question.UserAnswer)
	}

	return id, err
}
