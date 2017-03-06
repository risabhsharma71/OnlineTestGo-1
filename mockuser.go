package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Name struct {
	ID     string `json:"id,omitempty"`
	FName string `json:"fname,omitempty"`
	LName string `json:"lname,omitempty"`
	Phone int `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Test string `json:"test,omitempty"`

}

var name []Name

func GetName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range name {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Name{})
}

func GetNames(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(name)
}

func PostName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var nm Name
	_ = json.NewDecoder(req.Body).Decode(&nm)
	nm.ID = params["id"]
	name = append(name, nm)
	json.NewEncoder(w).Encode(name)
}

func DeleteName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range name {
		if item.ID == params["id"] {
			name = append(name[:index], name[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(name)
}

func main() {
	router := mux.NewRouter()
	name = append(name, Name{ID: "1", FName: "Sweta", LName: "Vahia",Phone:42423,Email:"ewe@fsd",Test:"java"})
	name = append(name, Name{ID: "2", FName: "Shishir", LName: "Vahia",Phone:42423,Email:"ewe@fsd",Test:"java"})
	name = append(name, Name{ID: "3", FName: "Aadya", LName: "Vahia",Phone:42423,Email:"ewe@fsd"})
	router.HandleFunc("/name", GetNames).Methods("GET")
	router.HandleFunc("/name/{id}", GetName).Methods("GET")
	router.HandleFunc("/name/{id}", PostName).Methods("POST")
	router.HandleFunc("/name/{id}", DeleteName).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

