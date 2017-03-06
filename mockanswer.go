package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Answer struct {
	ID     string `json:"id,omitempty"`
	Uid int `json:"uid,omitempty"`
	Qid int `json:"qid,omitempty"`
    Correctness int `json:"correctness,omitempty"`
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

/*func DeleteName(w http.ResponseWriter, req *http.Request) {
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
	ans = append(ans, Answer{ID: "1", Uid:1, Qid:1,Correctness:1})
	ans = append(ans, Answer{ID: "2", Uid:1, Qid:2,Correctness:0})
	ans = append(ans, Answer{ID: "3", Uid:1, Qid:3,Correctness:0})
	router.HandleFunc("/ans", GetAnswer).Methods("GET")
	router.HandleFunc("/ans/{id}", GetAnswers).Methods("GET")
	//router.HandleFunc("/ans/{id}", PostAnswer).Methods("POST")
	//router.HandleFunc("/ans/{id}", DeleteAnswer).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}