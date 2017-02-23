package models

//Question indicates the question table
type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Type     int64  `json:"type"`
}
