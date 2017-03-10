package models

//Question indicates the question table
type Question struct {
	ID         int      `json:"id"`
	Question   string   `json:"question"`
	Options    []string `json:"choices"`
	UserAnswer string   `json:"userAnswer"`
	Type       string   `json:"type"`
}
