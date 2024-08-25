package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map[F, T any](s []F, f func(F) T) []T {
	rst := make([]T, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}
func main2() {
	doubled := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	// 다른 타입이 사용 가능
	doubled1 := Map([]int{1, 2, 3}, func(v int) float32 {
		return float32(v * 2)
	})
	upper := Map([]string{"hello", "world", "abc"}, func(v string) string {
		return strings.ToUpper(v)
	})
	toString := Map([]int{1, 2, 3}, func(v int) string {
		return "str" + strconv.Itoa(v)
	})

	fmt.Println(doubled)
	fmt.Println(doubled1)
	fmt.Println(upper)
	fmt.Println(toString)
}

// T, U, V, S
// T1, T2 => 많이 사용
func Map2[T1, T2 any](s []T2, f func(T2) T1) []T1 {
	rst := make([]T1, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}
func Map3[argument, res any](s []res, f func(res) argument) []argument {
	rst := make([]argument, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}

// 제네릭: 틀
// 제네릭 -> 리팩토링할때 쓰는게 좋음

// 제네릭 타입
type Node[T any] struct {
	val  T
	next *Node[T]
}

// Make it work, make it right, make it fast. - Kent beck
