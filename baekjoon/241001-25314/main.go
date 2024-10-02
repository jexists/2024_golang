package main

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)

	for i := 0; i < num/4; i++ {
		fmt.Printf("long ")
	}
	fmt.Println("int")
}
