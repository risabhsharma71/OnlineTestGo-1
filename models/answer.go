package models

//Answer indicates answer table
type Answer struct {
	ID          int64  `json:"id"`
	Userid      int    `json:"userid"`
	Questionid  int    `json:"questionid"`
	Answer      string `json:"answer"`
	Correctness string `json:"correctness"`
}
