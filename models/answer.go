package models

//Answer indicates answer table
type Answer struct {
	ID          int64  `json:"id"`
	Uid      int    `json:"uid"`
	Qid  int    `json:"qid"`
	Correctness int `json:"correctness"`
}
