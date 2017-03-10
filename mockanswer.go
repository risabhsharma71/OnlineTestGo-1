package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/itsjamie/gin-cors"
)

type Answer struct {
		ID     string `json:"id,omitempty"`
	//Uid int `json:"uid,omitempty"`
	//Qid int `json:"qid,omitempty"`
   // Correctness int `json:"correctness,omitempty"`
	Score int `json:"score,omitempty"`


}

var ans []Answer

func GetAnswer(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range ans {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Answer{})
 }

 func GetAnswers(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(ans)

//func PostName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var as Answer
	_ = json.NewDecoder(req.Body).Decode(&as)
	as.ID = params["id"]
	ans = append(ans, as)
	json.NewEncoder(w).Encode(ans)
}

/*   func DeleteName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range name {
		if item.ID == params["id"] {
			name = append(name[:index], name[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(name)
}
*/


  func main() {
	router := mux.NewRouter()
        router.use(cors.Middleware(cors.Config{
    Origins:        "*",
    Methods:        "GET, PUT, POST, DELETE",
    RequestHeaders: "Origin, Authorization, Content-Type,Access-Control-Allow-Origin",
    ExposedHeaders: "Access-Control-Allow-Origin",
    MaxAge: 50 * time.Second,
    Credentials: true,
    ValidateHeaders: false,
}))
	ans = append(ans, Answer{ID: "1",Score:15})
	ans = append(ans, Answer{ID: "2",Score:10})
	ans = append(ans, Answer{ID: "3",Score:12})
	router.HandleFunc("/ans", GetAnswer).Methods("GET")
	router.HandleFunc("/ans/{id}", GetAnswers).Methods("GET")
	//router.HandleFunc("/ans/{id}", PostAnswer).Methods("POST")
	//router.HandleFunc("/ans/{id}", DeleteAnswer).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}