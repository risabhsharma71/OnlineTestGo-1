package daoimpl

import (
	"log"
)

type LogoutImpl struct{}

func (dao LogoutImpl) DeleteToken(token string) string {
	deltoken := "notfoud"
	query := "DELETE FROM token WHERE token=?"

	db := connection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return deltoken
	}

	defer stmt.Close()
	log.Println(token)
	res, err1 := stmt.Exec(token)

	if err1 != nil {
		log.Panic("Exec err:", err1.Error())
	}
	q, err2 := res.RowsAffected()
	if err2 != nil {
		log.Panic("Exec err:", err2.Error())
	}
	log.Println(q)
	if q > 0 {
		deltoken = "deleted"
	}

	return deltoken

}
