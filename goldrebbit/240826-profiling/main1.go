package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func Fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)
}

func main1() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Println(Fib1(50))

	time.Sleep(10 * time.Second)
}

// go tool pprof 파일명
// top

// (pprof) top
// Showing nodes accounting for 55.46s, 99.68% of 55.64s total
// Dropped 28 nodes (cum <= 0.28s)
//       flat  flat%   sum%        cum   cum%
//     55.46s 99.68% 99.68%     55.53s 99.80%  main.Fib
//          0     0% 99.68%     55.53s 99.80%  main.main
//          0     0% 99.68%     55.53s 99.80%  runtime.main
// web
