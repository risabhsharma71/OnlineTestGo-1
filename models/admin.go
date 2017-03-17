package models

//Admin indicates answer table
type Admin struct {
	Uid        int64   `json:"uid"`
	Fname      string  `json:"fname"`
	Lname      string  `json:"lname"`
	Type       string  `json:"type"`
	Score      int64   `json:"score"`
	
}
