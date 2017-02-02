package models

//Person indicates the person model
type Person struct {
	ID    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Test  string `json:"test"`
}
