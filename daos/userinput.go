package daos
import(
"log"
"database/sql"
 _"github.com/go-sql-driver/mysql"
)

 func insertlist(db *sql.DB) {
	rows, err := db.Query("insert into registration(id,fname,lname,phone,email) values(129,'Rakesh','kumar','9681817926','r@rapidqube.com')")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	err = rows.Err()
}
func getlist(db *sql.DB) {
	var (
		id         int
		fname string
		lname  string
		phone   int
		email   string
	)
	rows, err := db.Query("select * from registration")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &fname, &lname, &phone, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, fname, lname, phone, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}