package manager
import (
    "OnlineTestGo/daos/daoimpl"
    //"OnlineTestGo/daos/interfaces"
    "OnlineTestGo/models"
    //"fmt"
    "log"

    "OnlineTestGo/utility"
)
 func AddQuestion(question models.Question) int64 {
    utility.GetLogger()
    log.Println("calling register manager")
    questionDao := daoimpl.QuestionImpl{}
    

    
    insertedid, err := questionDao.AddQuestion(question)
    if err != nil {
        log.Println("error occured", err)
    }
    log.Println(insertedid)

    return insertedid
}