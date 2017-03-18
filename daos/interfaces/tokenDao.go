package interfaces

import "OnlineTestGo/models"

//QuestionDao interface to be implemented by Questionimpl
type TokenDao interface {
	GetToken(c *gin.Context,token models.Token)
    
           //AddOption(ID ,Options models.Question,userAnswer models,Question)(int64,error)
}
