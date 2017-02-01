package daos
import(
//"fmt"
"log"
"database/sql"
 _"github.com/go-sql-driver/mysql"
)

 
func main(){
 	db, err := sql.Open("mysql",
			"Rakesh:Root12345$@tcp(rpqb.centralindia.cloudapp.azure.com:3306)/interview_test")
	if err != nil {		
			log.Fatal(err)
	}
	insertlist(db)
	getlist(db)
	defer db.Close()

}