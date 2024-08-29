package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var fibMap map[int]int = make(map[int]int)

func Fib2(n int) int {
	if f, ok := fibMap[n]; ok {
		return f
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	f := Fib2(n-1) + Fib2(n-2)
	fibMap[n] = f
	return f
}

func main2() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Println(Fib2(50))

	time.Sleep(10 * time.Second)
}

// go tool pprof cpu.prof
// Type: cpu
// Time: Aug 26, 2024 at 10:29pm (KST)
// Duration: 10s, Total samples = 0
// No samples were found with the default sample value type.
// Try "sample_index" command to analyze different sample values.
// Entering interactive mode (type "help" for commands, "o" for options)

// (pprof) top
// Showing nodes accounting for 0, 0% of 0 total
//       flat  flat%   sum%        cum   cum%
