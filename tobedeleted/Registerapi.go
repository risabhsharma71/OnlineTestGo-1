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
			"rakesh:root@tcp(192.168.0.8:3306)/interview_test1")
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
    id int
	first_name string
	last_name  string
	phone_no string
	email_id string
}
router := gin.Default()
// GET a person detail
router.GET("/person/:id", func(c *gin.Context) {
	var (
		person Person
		result gin.H
		)
id := c.Param("id")
row := db.QueryRow("select * from registration_table where id =?;",id)
err = row.Scan(&person.id, &person.first_name, &person.last_name, &person.phone_no, &person.email_id)
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
	var persons []Person
rows, err := db.Query("select * from registration_table;")
if err != nil {
	fmt.Print(err.Error())
}
for rows.Next() {
	var person Person
	err = rows.Scan(&person.id, &person.first_name, &person.last_name, &person.phone_no, &person.email_id)
	persons = append(persons, person)
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

if err != nil {
	fmt.Print(err.Error())
}
// Fastest way to append strings
buffer.WriteString(first_name)
buffer.WriteString(" ")
buffer.WriteString(last_name)
defer stmt.Close()
name := buffer.String()
c.JSON(http.StatusOK, gin.H{
	"message": fmt.Sprintf(" %s successfully created", name),
	})
})

// PUT - update a person details
router.PUT("/person", func(c *gin.Context) {
	var buffer bytes.Buffer
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
if err != nil {
	fmt.Print(err.Error())
}

// Fastest way to append strings
buffer.WriteString(first_name)
buffer.WriteString(" ")
buffer.WriteString(last_name)
defer stmt.Close()
name := buffer.String()
c.JSON(http.StatusOK, gin.H{
	"message": fmt.Sprintf("Successfully updated to %s", name),
	})
})

// Delete resources
router.DELETE("/person", func(c *gin.Context) {
	id := c.Query("id")
	stmt, err := db.Prepare("delete from registration_table where id= ?;")
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