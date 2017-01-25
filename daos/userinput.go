package daos
import(
"log"
"database/sql"
 _"github.com/go-sql-driver/mysql"
)

 func insertlist(db *sql.DB)	{
rows, err := db.Query("insert into registration_table(id,first_name,last_name,phone_no,email_id) values(120,'Rakesh1','Bharati','9681817926','rakesh.bharati@rapidqube.com')")
if err != nil {
	log.Fatal(err)
	}
defer rows.Close()
err = rows.Err()
}
func getlist(db *sql.DB){
var(
		id int
		first_name string
		last_name string
		phone_no int
		email_id string
	)
rows, err := db.Query("select * from registration_table")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&id, &first_name, &last_name, &phone_no, &email_id)
	if err != nil { 
		log.Fatal(err)
	}
	log.Println(id, first_name, last_name, phone_no, email_id)
}
err = rows.Err()
if err != nil {
		log.Fatal(err)
}
}