package main
import ("bytes"
"database/sql"
"fmt"
"net/http"
"github.com/gin-gonic/gin"
_"github.com/go-sql-driver/mysql"
)
func main() {
	db, err := sql.Open("mysql",
<<<<<<< HEAD
			"rakesh:root@tcp(192.168.0.8:3306)/interview_test1")
=======
			"Rakesh:Root12345$@tcp(rpqb.centralindia.cloudapp.azure.com:3306)/interview_test")
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
if err != nil {
	fmt.Print(err.Error())
}
defer db.Close()
// make sure connection is available
err = db.Ping()
if err != nil {
	fmt.Print(err.Error())
	}
type Person struct {
<<<<<<< HEAD
    id int
	first_name string
	last_name  string
	phone_no string
	email_id string
=======
    	Id int `json:"id"`
	Fname string `json:"fname"`
	Lname  string `json:"lname"`
	Phone string `json:"phone"`
	Email string `json:"email"`
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
}
router := gin.Default()
// GET a person detail
router.GET("/person/:id", func(c *gin.Context) {
	var (
		person Person
		result gin.H
		)
id := c.Param("id")
<<<<<<< HEAD
row := db.QueryRow("select * from registration_table where id =?;",id)
err = row.Scan(&person.id, &person.first_name, &person.last_name, &person.phone_no, &person.email_id)
=======
row := db.QueryRow("select * from registration where id =?;",id)
err = row.Scan(&person.Id, &person.Fname, &person.Lname, &person.Phone, &person.Email)
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
			}
} else {
	result = gin.H{
		"result": person,
		"count":  1,
	}
}
c.JSON(http.StatusOK, result)
})	

// GET all persons
router.GET("/persons", func(c *gin.Context) {
<<<<<<< HEAD
	var (
		person  Person
		persons []Person
	)
rows, err := db.Query("select * from registration_table;")
=======
	var persons []Person
rows, err := db.Query("select * from registration;")
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
if err != nil {
	fmt.Print(err.Error())
}
for rows.Next() {
<<<<<<< HEAD
	err = rows.Scan(&person.id, &person.first_name, &person.last_name, &person.phone_no, &person.email_id)
	persons = append(persons, person)
=======
	var person Person
	err = 	rows.Scan(&person.Id, &person.Fname, &person.Lname, &person.Phone, &person.Email)
	//persons = append(persons, Person{id: person.id, fname: person.Fname, lname: person.lname, phone: person.phone, email: person.email})
	persons = append(persons, Person{person.Id, person.Fname, person.Lname, person.Phone, person.Email})
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
	if err != nil {
		fmt.Print(err.Error())
	}
}
defer rows.Close()
c.JSON(http.StatusOK, gin.H{
	"result": persons,
	"count":  len(persons),
})	
})

// POST new person details
router.POST("/person", func(c *gin.Context) {
	var buffer bytes.Buffer
<<<<<<< HEAD
	id := c.PostForm("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	phone_no := c.PostForm("phone_no")
	email_id := c.PostForm("email_id")
	stmt, err := db.Prepare("insert into registration_table(id, first_name, last_name, phone_no, email_id) values(?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id, first_name, last_name, phone_no, email_id)
=======
	Id := c.PostForm("id")
	Fname := c.PostForm("fname")
	Lname := c.PostForm("lname")
	Phone := c.PostForm("phone")
	Email := c.PostForm("email")
	stmt, err := db.Prepare("insert into registration(id, fname, lname, phone, email) values(?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(Id, Fname, Lname, Phone, Email)
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae

if err != nil {
	fmt.Print(err.Error())
}
// Fastest way to append strings
<<<<<<< HEAD
buffer.WriteString(first_name)
buffer.WriteString(" ")
buffer.WriteString(last_name)
=======
buffer.WriteString(Fname)
buffer.WriteString(" ")
buffer.WriteString(Lname)
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
defer stmt.Close()
name := buffer.String()
c.JSON(http.StatusOK, gin.H{
	"message": fmt.Sprintf(" %s successfully created", name),
	})
})

// PUT - update a person details
router.PUT("/person", func(c *gin.Context) {
	var buffer bytes.Buffer
<<<<<<< HEAD
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
    /*phone_no := c.PostForm("phone_no")
	email_id := c.PostForm("email_id")*/
	stmt, err := db.Prepare("update registration_table set first_name= ?, last_name= ? where id= ?;")
if err != nil {
	fmt.Print(err.Error())
	}
_, err = stmt.Exec(first_name,last_name, id)
=======
	Id := c.Query("id")
	Fname := c.PostForm("fname")
	Lname := c.PostForm("lname")
   // phone := c.PostForm("phone")
	// email := c.PostForm("email")
	stmt, err := db.Prepare("update registration set fname= ?, lname= ? where id= ?;")
if err != nil {
	fmt.Print(err.Error())
	}
_, err = stmt.Exec(Fname,Lname,Id)
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
if err != nil {
	fmt.Print(err.Error())
}

// Fastest way to append strings
<<<<<<< HEAD
buffer.WriteString(first_name)
buffer.WriteString(" ")
buffer.WriteString(last_name)
=======
buffer.WriteString(Fname)
buffer.WriteString(" ")
buffer.WriteString(Lname)
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
defer stmt.Close()
name := buffer.String()
c.JSON(http.StatusOK, gin.H{
	"message": fmt.Sprintf("Successfully updated to %s", name),
	})
})

// Delete resources
router.DELETE("/person", func(c *gin.Context) {
	id := c.Query("id")
<<<<<<< HEAD
	stmt, err := db.Prepare("delete from registration_table where id= ?;")
=======
	stmt, err := db.Prepare("delete from registration where id= ?;")
>>>>>>> 05d3d8968880bc019f3911c983208a2bf0ca93ae
	if err != nil {
		fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
			}
c.JSON(http.StatusOK, gin.H{
	"message": fmt.Sprintf("Successfully deleted user: %s", id),
	})
})
router.Run(":9090")
}