package models

//Question indicates the question table
type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2"`
	Option3  string `json:"option3"`
	Option4  string `json:"option4"`
	Option5  string `json:"option5"`
	Answer   string `json:"answer"`
	Type     int64  `json:"type"`
}
