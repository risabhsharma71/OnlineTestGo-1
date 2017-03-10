package daoimpl

import (
	"fmt"
	//"log"
	"OnlineTestGo/models"
)

//QuestionImpl struct for Questions implementattion
type TotalQuestionImpl struct{}

//GetQuestions fetched the questions from quesions table takes testtype as input
func (dao TotalQuestionImpl) GetQuesions(testtype string) []models.TotalQuestion {
var totalquestions []models.TotalQuestion


      var totalquestion models.TotalQuestion
               

	db := connection()
	rows, err := db.Query("select Q.id,Q.question,O.choices as c1 from questions Q,Options O where Q.id=O.qid and type=?",testtype)
       
	fmt.Println(testtype)
	if err != nil {
		fmt.Print(err.Error())
		//return models.Question{}, err
	}

	for rows.Next() {
	    

			err = rows.Scan(&totalquestion.ID, &totalquestion.Question ,&totalquestion.Choices)
                        encountered := map[models.Question]bool{}

                       for v:= range totalquestion {
        encountered[totalquestion[v]] = true
    }
              result := []models.Question{}
                        for key, _ := range encountered {
        result = append(result, key)
    }
   
}
			totalquestions = append(totalquestions,result)
                        
			if err != nil {
				fmt.Print(err.Error())
			}
		}
          
	defer rows.Close()

	return totalquestions



}