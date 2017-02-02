package userModels

type Person struct {
    	Id int `json:"id"`
	Fname string `json:"fname"`
	Lname  string `json:"lname"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}