package daoimpl

import (
	"log"

	"OnlineTestGo/models"
)

type AnswerImpl struct{}

func (dao AnswerImpl) SaveAnswer(answer models.Answer) (int64, error) {

	query := "insert into answers(userid,questionid,answer) values(?,?,?)"
	db := connection()
	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(answer.Userid, answer.Questionid, answer.Answer)

	if err != nil {
		log.Panic("Exec err:", err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Exec err:", err.Error())
	}

	return id, err
}
