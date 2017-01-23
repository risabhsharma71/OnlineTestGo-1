package main 
import(
//"fmt"
"log"
"database/sql"
 _"github.com/go-sql-driver/mysql"
)

 
func main(){
 	db, err := sql.Open("mysql",
			"rakesh:root@tcp(192.168.0.8:3306)/interview_test1")
	if err != nil {		
			log.Fatal(err)
	}
	insertlist(db)
	getlist(db)
	defer db.Close()

}
func insertlist(db *sql.DB)	{
rows, err := db.Query("insert into registration_table(Registration_id,First_Name,Last_name,Phone_no,Email_id) values(129,'sweta1','vahia1','9892098920','s@rapidqube.com')")
if err != nil {
	log.Fatal(err)
	}
defer rows.Close()
err = rows.Err()
}
func getlist(db *sql.DB){
var(
		Registration_id int
		First_Name string
		Last_name string
		Phone_no int
		Email_id string
	)
rows, err := db.Query("select * from registration_table")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&Registration_id, &First_Name, &Last_name, &Phone_no, &Email_id)
	if err != nil { 
		log.Fatal(err)
	}
	log.Println(Registration_id, First_Name, Last_name, Phone_no, Email_id)
}
err = rows.Err()
if err != nil {
		log.Fatal(err)
}
}



