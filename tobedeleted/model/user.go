
package models 
2 
 
3 type ( 
4 	User struct { 
5 		Registration_id     int   	 `json:"Registration_id" bson:"_Registration_id"` 
6 		First_Name   		string	 `json:"First_Name" bson:"First_Name"` 
		Last_Name           string 	 `json:"Last_Name" bson:"Last_Name"`
		Phone_no            int    	 `json:"Phone_no" bson:"Phone_no"`  
7 		Email_id            string 	 `json:"Email_id" bson:"Email_id"` 
8 		
9 	} 
10 ) 
 




