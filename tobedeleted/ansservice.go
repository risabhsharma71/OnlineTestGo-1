package main
import ( 
   //"bytes" 
   "database/sql" 
   "fmt" 
   "net/http" 
   "github.com/gin-gonic/gin" 
  _ "github.com/go-sql-driver/mysql" 
 ) 

func main() { 
  db, err := sql.Open("mysql", "rakesh:root@tcp(192.168.0.8:3306)/interview_test1") 
  if err != nil { 
      fmt.Print(err.Error()) 
  } 

  defer db.Close()
err = db.Ping() 
  if err != nil { 
      fmt.Print(err.Error()) 
  } 
  type Answer struct { 
    result_id        int 
    test_id          int 
    id               int 
    answer           string
    correctness      string
 } 
 router := gin.Default()

// GET a answer of one person 
 router.GET("/answer/:id", func(c *gin.Context) { 
     var ( 
          answer Answer 
          result gin.H 
         )
         id := c.Param("id") 
      row := db.QueryRow("select * from answer_table where id=?;",id) 
      err = row.Scan(&answer.result_id, &answer.test_id, &answer.id ,&answer.answer, &answer.correctness) 
      if err != nil { 
        result = gin.H{ 
              "result": nil, 
              "count":  0, 
          } 
      } else { 
        result = gin.H{ 
              "result": answer, 
              "count":  1, 
          } 
      } 
      c.JSON(http.StatusOK, result) 
  }) 
 //get all the answers
router.GET("/answers", func(c *gin.Context) { 
      var ( 
          answer  Answer 
          answers []Answer 
      ) 
rows, err := db.Query("select * from answer_table;") 
     if err != nil { 
          fmt.Print(err.Error()) 
      } 
      for rows.Next() { 
          err = rows.Scan(&answer.result_id, &answer.test_id, &answer.id,&answer.answer,&answer.correctness)
          answers = append(answers, answer) 
          if err != nil { 
              fmt.Print(err.Error()) 
          } 
      } 
        defer rows.Close() 
        c.JSON(http.StatusOK, gin.H{ 
          "result": answers, 
          "count":  len(answers), 
     }) 
  }) 

//Inserting new answers to database
 
 router.POST("/answer", func(c *gin.Context) { 
      //var buffer bytes.Buffer 
       result_id := c.PostForm("result_id") 
       test_id := c.PostForm("test_id") 
       id := c.PostForm("id)") 
       answer := c.PostForm("answer")
       correctness := c.PostForm("correctness")
       stmt, err := db.Prepare("insert into answer_table(result_id, test_id, id, answer, correctness) values(?,?,?,?,?);") 
      if err != nil { 
          fmt.Print(err.Error()) 
      }
  _, err = stmt.Exec(result_id,test_id, id, answer, correctness)
    
      /*if err != nil { 
          fmt.Print(err.Error()) 
      } 
      buffer.WriteString(ans) 
      buffer.WriteString(" ") 
      buffer.WriteString(id) 
      defer stmt.Close() 
      name := buffer.String() 
         c.JSON(http.StatusOK, gin.H{ 
             "message": fmt.Sprintf(" %s successfully created", name), 
         })*/
    })

         // Delete resources 
     router.DELETE("/answer", func(c *gin.Context) { 
         result_id := c.Query("result_id") 
         stmt, err := db.Prepare("delete from answer_table where result_id= ?;") 
         if err != nil { 
             fmt.Print(err.Error()) 
         } 
         _, err = stmt.Exec(result_id) 
         if err != nil { 
             fmt.Print(err.Error()) 
         } 
         c.JSON(http.StatusOK, gin.H{ 
             "message": fmt.Sprintf("Successfully deleted user: %s", result_id), 
         }) 
     }) 

router.Run(":3000")
}




 



