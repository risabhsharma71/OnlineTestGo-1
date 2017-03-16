 package daoimpl

import(
	"OnlineTestGo/models"
	"log"
      )

type AdminImpl struct{}

func (dao AdminImpl) FetchData() []models.Admin {
	var datas []models.Admin
	
	query := "select r.fname,r.lname,a.uid,a.score from rpqbmysql.registration r join rpqbmysql.answers a on r.id=a.uid"

	db := connection()
	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var data models.Admin
		err = rows.Scan(&data.Fname, &data.Lname, &data.Uid,&data.Score)

		if err != nil {
			log.Fatal(err)
		}

		datas = append(datas, data)

	}
	//log.Println("data:",data)

      return datas
}