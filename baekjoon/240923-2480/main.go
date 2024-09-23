package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	_, _ = fmt.Scanln(&a, &b, &c)
	arr := []int{a, b, c}
	sort.Ints(arr)
	var result = 0
	a, b, c = arr[0], arr[1], arr[2]
	if a == b && b == c {
		result = 10000 + a*1000
	} else if a == b || b == c {
		result = 1000 + b*100
	} else {
		result = c * 100
	}
	fmt.Println(result)
}

// func mySolution() {

// 	var a, b, c int
// 	reader := bufio.NewReader(os.Stdin)
// 	writer := bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()

// 	fmt.Fscanf(reader, "%d %d %d", &a, &b, &c)
// 	var money = 0
// 	if a == b && a == c {
// 		money = 10000 + a*1000
// 	} else if a == b || b == c || a == c {
// 		if a == b {
// 			money = 1000 + a*100
// 		} else if b == c {
// 			money = 1000 + b*100
// 		} else if a == c {
// 			money = 1000 + a*100
// 		}
// 	} else if a != b && a != c {
// 		var bigest = a
// 		if bigest < b {
// 			bigest = b
// 			if bigest < c {
// 				bigest = c
// 			}
// 		} else {
// 			if bigest < c {
// 				bigest = c
// 			}
// 		}
// 		money = bigest * 100
// 	}

// 	fmt.Println(money)
// }
