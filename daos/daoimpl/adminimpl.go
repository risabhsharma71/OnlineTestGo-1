package daoimpl

import (
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"
)

type AdminImpl struct{}

func (dao AdminImpl) FetchData() []models.Admin {

	var datas []models.Admin

	utility.GetLogger()
	log.Println("entering FetchData()")

	log.Println("executing query and Fetching data from db ")

	query := " select a.uid,r.fname,r.lname,a.testtype,a.score from onlinetestdb.answers a, onlinetestdb.registration r where r.id=a.uid"

	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var data models.Admin
		err = rows.Scan(&data.Uid, &data.Fname, &data.Lname, &data.Type, &data.Score)

		if err != nil {
			log.Fatal(err)
		}

		datas = append(datas, data)

	}

	return datas
}
