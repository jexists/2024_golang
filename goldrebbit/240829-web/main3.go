package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler4() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello bar")
	})

	return mux
}

// mux 사용
func main4() {

	http.ListenAndServe(":3000", MakeWebHandler())
}
