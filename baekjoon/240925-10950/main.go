package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)

	var a, b int
	for i := 0; i < num; i++ {
		_, _ = fmt.Scanln(&a, &b)

		fmt.Println(a + b)
	}
}
