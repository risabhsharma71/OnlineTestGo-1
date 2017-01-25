package daos
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