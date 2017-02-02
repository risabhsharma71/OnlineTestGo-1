package manager

import (
	"fmt"

	"github.com/MIghtykukulkan/OnlineTestGo/models"
)

//Register manager takes care of business logic like calling daos
func Register(user models.Person) []models.Question {
	fmt.Println("calling regiter manager")

	var questionList []models.Question

	testType = user.TestType
	//make connection to

	return questionList
}
