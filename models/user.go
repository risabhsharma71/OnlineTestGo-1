package models

//User indicates the person model
type User struct {
	ID       int    `json:"id"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Test     string `json:"test"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}
