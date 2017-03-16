package daoimpl

import (
	"log"
        "fmt"
	"OnlineTestGo/models"
)
       
type UserImpl struct{}

func (dao UserImpl) SaveNewUser(user models.User) (int, error) {

	query := "insert into registration(fname, lname, phone, email) values(?,?,?,?)"
	db := connection()
	defer db.Close()
          
	stmt, err := db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Fname, user.Lname, user.Phone, user.Email)

	if err != nil {
		log.Panic("Exec err:", err.Error())
	}

	_, err := res.LastInsertId()

	fmt.Printf("%T",id);
	if err != nil {
		log.Println("Exec err:", err.Error())
	}
	
	return 0, err
}
func (dao UserImpl) CheckUser(user models.User) (int, error) {
                
	        id:=0
query := "select phone from registration where phone = 9176061985"
                   db:= connection()
	     
 	  rows, err := db.Query(query)	
          if err != nil {
	  log.Fatal(err)
              defer rows.Close()
}

for rows.Next() {
         var  no[] models.User
	err := rows.Scan(&user.Phone)
	if err != nil {
		log.Fatal(err)
	}
           no= append(no,user) 


           if len(no)!=0{
          return id,err
         }
	fmt.Println(id)
}

     return 0,err;
}