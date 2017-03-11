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
                         
			totalquestions = append(totalquestions,totalquestion)
                        
			if err != nil {
				fmt.Print(err.Error())
			}
		}
          for_, v:= range totalquestions{
              if (v.id==totalquestions.id){
totalquestions=append(totalquestions,totalquestions.choices)

}

         defer rows.Close()

	return totalquestions
}

}

