
package models 
 
 
 type ( 
 	User struct { 
 		Registration_id     int   	 `json:"Registration_id" bson:"_Registration_id"` 
 		First_Name   		string	 `json:"First_Name" bson:"First_Name"` 
		Last_Name           string 	 `json:"Last_Name" bson:"Last_Name"`
		Phone_no            int    	 `json:"Phone_no" bson:"Phone_no"`  
 		Email_id            string 	 `json:"Email_id" bson:"Email_id"` 
 		
 	} 
 ) 
 




