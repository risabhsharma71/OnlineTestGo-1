
package daoimpl

import (
	"log"
	"OnlineTestGo/models"
      
         
)


    var (

        sum int=0
	 Tanswer []models.Options
         answers models.Options
	 tanswer models.Options
         Tuser_answers =[]string {"java","0","0","0","java","0","0","0","java","0","0","0","java","0","0","0"}
         

)
type AnswerImpl struct{}

 func (dao AnswerImpl) SaveAnswer(answer models.Answer) (int64, error){


  

// fetching correctness//
db := connection()
rows, err := db.Query("select answer from Options")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
        err := rows.Scan(&tanswer)
        Tanswer=append(Tanswer,tanswer)
	if (err != nil) {
		log.Fatal(err)
	}
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}


//removeing redundant zeros in correctness array//

remove := []string{"0"}

loop:
for a := 0; a < len(Tanswer); a++ {
    url :=  Tanswer 
    for _, rem := range remove {
        if url == rem {
            Tanswer = append(Tanswer[:a], Tanswer[a+1:]...)
            a-- // Important: decrease index
            continue loop
        }
    }
}

log.Println(Tanswer)
return 1,err
}






	func ArrayCompare(Tuser_answers[] string , Tanswer []models.Options ){
    var score int=0

    // Loop over  ints and print them.
    for i := 0; i < len(Tuser_answers); i++ {
        if(Tuser_answers[i]==Tanswer[i]){
        score=score+1
    }
        
    }
    log.Println(score)

}


