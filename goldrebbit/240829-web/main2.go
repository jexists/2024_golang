package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "world"
	}
	id, _ := strconv.Atoi(values.Get("id"))

	fmt.Fprint(w, "hello %s! id:%d", name, id)
}

func main2() {
	http.HandleFunc("/", barHandler)

	http.ListenAndServe(":3000", nil)
}
