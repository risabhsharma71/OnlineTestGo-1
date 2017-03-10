package models

//Answer indicates answer table
type Answer struct {
	ID          int64  `json:"id"`
	Uid      int    `json:"uid"`
	Q_type string    `json:"q_type"`
	Score int `json:"score"`
}
