package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	arr := make([]int, n)
	for i := 0; i < m; i++ {
		var i, j, k int
		fmt.Scan(&i, &j, &k)

		for m := i; m <= j; m++ {
			arr[m-1] = k
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}

}

// func main() {
// 	var n, m int
// 	fmt.Scan(&n, &m)

// 	arr := make([]int, n)
// 	for i := 0; i < m; i++ {
// 		var i, j, k int
// 		fmt.Scan(&i, &j, &k)
// 		for i := i - 1; i <= j-1; i++ {
// 			arr[i] = k
// 		}
// 	}
// 	// fmt.Println(arr)
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Print(arr[i])
// 		if i == len(arr)-1 {
// 			fmt.Print("\n")
// 		} else {
// 			fmt.Print(" ")
// 		}
// 	}
// }
