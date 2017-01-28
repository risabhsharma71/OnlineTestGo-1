package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "rakesh:root@tcp(192.168.0.8:3306)/interview_test1")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	type Question struct {
		id     int
		t_type string
		que    string
		op1    string
		op2    string
		op3    string
		op4    string
		choice string
		ans    string
	}
	router := gin.Default()
	//Get one question
	router.GET("/question/:id", func(c *gin.Context) {
		var (
			question Question
			result   gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select * from question_table where id=?;", id)
		err = row.Scan(&question_table.id, &question_table.t_type, &question_table.que, &question_table.op1, &question_table.op2, &question_table.op3, &question_table.op4, &question_table.choice, &question_table.ans)
		if err != nil {
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": question,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})
	//Get all questions
	router.GET("/question", func(c *gin.Context) {
		var (
			question  Question
			questions []Question
		)
		rows, err := db.Query("select * from question_table")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&question_table.id, &question_table.t_type, &question_table.que, &question_table.op1, &question_table.op2, &question_table.op3, &question_table.op4, &question_table.choice, &question_table.ans)
			questions = append(questions, question)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": questions,
			"count":  len(questions),
		})
	})

	//Post New Question
	router.POST("/question", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.PostForm("id")
		t_type := c.PostForm("t_type")
		que := c.PostForm("que")
		op1 := c.PostForm("op1")
		op2 := c.PostForm("op2")
		op3 := c.PostForm("op3")
		op4 := c.PostForm("op4")
		choice := c.PostForm("choice")
		ans := c.PostForm("ans")
		stmt, err := db.Prepare("insert into question_table(id, t_type, que, op1, op2, op3, op4, choice, ans values(?,?,?,?,?,?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id, t_type, que, op1, op2, op3, op4, choice, ans)
		if err != nil {
			fmt.Print(err.Error())
		}
		// Fastest way to append strings
		buffer.WriteString(t_type)
		buffer.WriteString("que ")
		buffer.WriteString(que)
		defer stmt.Close()
		que := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", que),
		})
	})

	// PUT - update a question
	router.PUT("/question", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.Query("id")
		que := c.PostForm("que")
		op1 := c.PostForm("op1")
		op2 := c.PostForm("op2")
		op3 := c.PostForm("op3")
		op4 := c.PostForm("op4")
		stmt, err := db.Prepare("update question_table set que= ?, op1= ?, op2=?, op3=?, op4=? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(que, op1, op2, op3, op4, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(que)
		buffer.WriteString("op1")
		buffer.WriteString(op1)
		buffer.WriteString("op2")
		buffer.WriteString(op2)
		buffer.WriteString("op3")
		buffer.WriteString(op3)
		buffer.WriteString("op4")
		buffer.WriteString(op4)
		defer stmt.Close()
		que := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", que),
		})
	})

	// Delete resources
	router.DELETE("/question", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from question_table where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted question: %s", id),
		})
	})
	router.Run(":3000")
}
