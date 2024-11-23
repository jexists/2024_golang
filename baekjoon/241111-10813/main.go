package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	for i := 0; i < m; i++ {
		var j, k int
		fmt.Scan(&j, &k)

		l := arr[j-1]
		arr[j-1] = arr[k-1]
		arr[k-1] = l
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}

}
