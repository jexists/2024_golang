package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// statefull

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	mux.HandleFunc("/students", PostStudentHandler).Methods("PUT")

	students = make(map[int]Student)

	students[1] = Student{1, "aaa", 16, 84}
	students[1] = Student{2, "bbb", 16, 84}
	lastId = 2
	return mux
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make([]Student, len(students))
	for _, val := range students {
		list = append(list, val)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}

// 썬더클라이언트 Thunder client
