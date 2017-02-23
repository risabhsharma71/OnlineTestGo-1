package daoimpl

import (
	//"fmt"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connection() *sql.DB {
	db, err := sql.Open("mysql",
		"umashankar:Rpqb123@tcp(db4free.net:3306)/rpqbmysql")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
