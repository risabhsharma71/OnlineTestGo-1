package models

//Question indicates the question table
type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
        Choices  []*Options `json:"choices"`
	
        Type     string  `json:"type"`
       

        


}
