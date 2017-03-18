package tos

//User indicates the person model
type Tokento struct {
	Uid      int    `json:"uid"`
	Fname    string `json:"fname"`
	Token    string `json:"token"`
	UserType string `json:"usertype"`
}
