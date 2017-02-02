package daoimpl

import (
	//"fmt"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connection() *sql.DB {
	db, err := sql.Open("mysql",
		"Rakesh:Root12345$@tcp(rpqb.centralindia.cloudapp.azure.com:3306)/interview_test")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
