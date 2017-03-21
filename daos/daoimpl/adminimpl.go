package daoimpl

import (
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"
)

type AdminImpl struct{}

func (dao AdminImpl) FetchData() []models.Admin {
	utility.GetLogger()
	log.Println("entering FetchData()")
	var datas []models.Admin
	log.Println("executing query and Fetching data from db ")
	query := "select c.uid,a.fname,a.lname,b.type,c.score from rpqbmysql.registration a,rpqbmysql.questions b,rpqbmysql.answers c where b.id=c.uid"

	db := connection()
	defer db.Close()

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
