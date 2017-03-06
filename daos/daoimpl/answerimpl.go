package daoimpl

import (
	"log"
	"OnlineTestGo/models"
         
)


    var (
    sum int=0
	 correctness[]int
	 answer[]int
        

)
type AnswerImpl struct{}

func (dao AnswerImpl) SaveAnswer(answer models.Answer) (int64, error){

	query := "insert into answers(uid,qid,correctness) values(?,?,?)"
	db := connection()
	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(answer.Uid, answer.Qid, answer.Correctness)

	if err != nil {
		log.Panic("Exec err:", err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Exec err:", err.Error())
	}

// fetching correctness//
rows, err := db.Query("select correctness from answers where qid = ?", 1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(correctness)
	if (err != nil) {
		log.Fatal(err)
	}
	log.Println(correctness)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}

//fetching user answer//
rows, err = db.Query("select answers from options where qid = ?", 1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(answer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(answer)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}
  
return id ,err

}

func IntArrayEquals(correctness[] int , answer[] int ) bool {
    if len(correctness) != len(answer) {
        return false
    }
    for i, v := range correctness{
        if v != answer[i] {
            return false
        }
    }
	sum++
    return true
}

