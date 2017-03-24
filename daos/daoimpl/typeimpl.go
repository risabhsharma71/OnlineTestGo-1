package daoimpl

import (
	"database/sql"
	"log"
)

//TypeImpl is a implementation for type interface
type TypeImpl struct{}

//GetIdfromType returns the id from Type table for input value
func (dao TypeImpl) GetIdfromType(s string) int64 {
	var id int64
	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()
	err := db.QueryRow("select id from type where value=?", s).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No type with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Type is %v\n", id)
	}

	return id
}
