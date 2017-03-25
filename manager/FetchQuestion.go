package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"log"
	"math/rand"
	"time"
)

func FetchQuestion(testtype string) []tos.Question {

	utility.GetLogger()
	log.Println("entering into manager.FetchQuestion")

	log.Println("calling Question manager")

	questionDao1 := daoimpl.QuestionImpl{}

	log.Println("calling FetchQuestionsByType()")

	questionlist := questionDao1.FetchQuestionsByType(testtype)
	shuffle(questionlist)
	return questionlist

}
func shuffle(arr []tos.Question) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond())) // no shuffling without this line

	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
