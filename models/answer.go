package answerModels
  
type (Answer struct {
		Id   			int		`json:"id"`
		Userid     		int		`json:"userid"`
		Questionid      int		`json:"questionid"`
		Answer      	string	`json:"answer"`
		Correctness 	string	`json:"correctness"`
	}
)

