package main

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof"

	// 사용하지않아도 안지워지게하기 위해서 _ 추가
	"time"
)

// 프로파일링 보기위해 링크
// http://localhost:8080/debug/pprof

func main() {
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- http.StatusOK
	}()

	select {
	case status := <-ch:
		w.WriteHeader(status)
	case <-time.After(200 * time.Millisecond):
		w.WriteHeader(http.StatusRequestTimeout)
	}
}

// 부하발생
// .\hey.exe http://localhost:8080/log
