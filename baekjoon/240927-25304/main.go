package main

import (
	"fmt"
)

func main() {
	var totalPrice, num int
	fmt.Scan(&totalPrice, &num)

	var price, count, finalPrice int
	for i := 0; i < num; i++ {
		fmt.Scan(&price, &count)
		finalPrice += (price * count)
	}
	if totalPrice == finalPrice {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
