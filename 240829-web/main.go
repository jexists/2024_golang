package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age,omitempty"`
	Score int
	meta  string
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 34, "meta"}
	// 대문자만 json변환
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
