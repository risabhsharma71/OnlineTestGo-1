package main

import (
	"encoding/json"
	"log"
	"net/http"
        "github.com/itsjamie/gin-cors"
	"github.com/gorilla/mux"
        "time"
        
 //"gopkg.in/gin-gonic/gin.v1"
)

type Question struct {
	ID       string    `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
    Option1  string  `json:"option1,omitempty"`
    Option2  string  `json:"option2,omitempty"`
    Option3  string  `json:"option3,omitempty"`
    Option4  string  `json:"option4,omitempty"`
    Option5  string  `json:"option5,omitempty"`
	Type string  `json:"type,omitempty"`

}

var questions []Question

func GetQuestion(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range questions {
		if item.Type == params["type"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Question{})
}

func GetQuestions(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(questions)
}


func main() {
	router := mux.NewRouter()
       router.Use(cors.Middleware(cors.Config{
    Origins:        "*",
    Methods:        "GET, PUT, POST, DELETE",
    RequestHeaders: "Origin, Authorization, Content-Type,Access-Control-Allow-Origin",
    ExposedHeaders: "Access-Control-Allow-Origin",
    MaxAge: 50 * time.Second,
    Credentials: true,
    ValidateHeaders: false,
}))	
        questions = append(questions, Question{ID: "1", Question: "which of these is necesary condn for automatic type conversion?", Option1: "a", Option2: "b",Option3: "c", Option4: "d",Option5: "e",Type :"java"})
	questions = append(questions, Question{ID: "2", Question: "what is the output of this program?", Option1: "a", Option2: "b",Option3: "c", Option4: "d",Option5: "e",Type :"java"})
	
	router.HandleFunc("/questions", GetQuestion).Methods("GET")
	router.HandleFunc("/questions/{type}", GetQuestions).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":9090", router))
}
