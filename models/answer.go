package models

//Answer indicates answer table
type Answer struct {
	Uid      int64  `json:"uid"`
	Qid      int64  `json:"qid"`
	Selected string `json:"selected"`
}
