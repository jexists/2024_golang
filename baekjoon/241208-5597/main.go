package main

import (
	"fmt"
)

func main() {
	var submitCount = 28
	student := make([]int, submitCount+2)
	for i := 0; i < submitCount; i++ {
		var j int
		fmt.Scan(&j)
		student[j-1] = 1
	}

	for i, v := range student {
		if v == 0 {
			fmt.Println(i + 1)
		}
	}
}

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	var submitCount = 28
// 	student := make([]int, submitCount+2)
// 	for i := 0; i < submitCount; i++ {
// 		var j int
// 		fmt.Scan(&j)
// 		student[j-1] = j
// 	}

// 	notSubmit := make([]int, 0, 2)

// 	for i, v := range student {
// 		if v == 0 {
// 			notSubmit = append(notSubmit, i+1)
// 		}
// 	}

// 	sort.Slice(notSubmit, func(i, j int) bool {
// 		return notSubmit[i] < notSubmit[j]
// 	})

// 	for _, v := range notSubmit {
// 		fmt.Println(v)
// 	}
// }
// func main() {
// 	var submitCount = 2
// 	student := make([]int, 0, submitCount)
// 	for i := 0; i < submitCount; i++ {
// 		var j int
// 		fmt.Scan(&j)
// 		student = append(student, j)
// 	}

// 	sort.Slice(student, func(i, j int) bool {
// 		return student[i] < student[j]
// 	})

// 	fmt.Println(student)
// }
