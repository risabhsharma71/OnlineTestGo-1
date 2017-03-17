package models

type Admin struct {
	
	Uid string `json:"uid"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Score int64 `json:"score"`
}
