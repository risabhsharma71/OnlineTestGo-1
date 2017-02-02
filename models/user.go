package models

//User indicates the person model
type User struct {
	ID    int64  `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Test  string `json:"test"`
}
