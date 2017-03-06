package models

//Question indicates the question table
type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Type     string  `json:"type"`
}
