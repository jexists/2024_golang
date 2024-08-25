package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Print[T1 any, T2 any](a T1, b T2) {
	fmt.Println(a, b)
}

func Print2(a, b interface{}) {
	fmt.Println(a, b)
}

// T => 타입이 하나만 제한
// func add[T int | float](a, b T) T {}

func add[T1 int | float32, T2 int | float32](a, b T1) T1 {
	return 0
}

func add1[T1 float32 | float64, T2 float32 | float64](a, b T1) T2 {
	// return a + b
	return T2(a + b)
}

func add3[T any](a, b T) T {
	// return a + b
	// error invalid operation
	return a
}
func add4[T int8 | int16 | int32 | int](a, b T) T {
	// T int8 | int16 | int32 | int == constraints.Integer
	return a + b
	// error invalid operation
}
func add5[T constraints.Integer | constraints.Float](a, b T) T {
	return a + b
	// error invalid operation
}
func add6[T constraints.Integer | constraints.Float | string](a, b T) T {
	return a + b
	// error invalid operation
}
func add7[T Int1](a, b T) T {
	return a + b
	// error invalid operation
}
func add8[T Int2](a, b T) T {
	return a + b
	// error invalid operation
}

func main() {
	Print(1, 2)

	var a float32 = 3
	var b float32 = 6
	c := add1[float32, float32](a, b)
	fmt.Println(c)

	fmt.Println(add3(1, 2))
	fmt.Println(add4(1, 2))
	fmt.Println(add5(1, 2))
	fmt.Println(add5(1.21, 2.23))
	// fmt.Println(add5("a", "b")) error
	fmt.Println(add6("a", "b"))

	var a1 MyInt = 1
	var b1 MyInt = 2
	var a2 int32 = 1
	var b2 int32 = 2
	fmt.Println(add7(a1, b1))
	fmt.Println(add7(a2, b2))
	// fmt.Println(add8(a1, b1)) Error
	fmt.Println(add8(a2, b2))
	// c1 := a1 + b1 // error

	// boxing / unboxing
	var v1 int = 3
	var v2 interface{} = v1 // boxing
	var v3 int = v2.(int)   // unboxing

	fmt.Println("--", v3)
}

// type Float constraints {
// 	~float32 | ~float64
// }

// 타입제한이 인터페이스로 되있는다 (키워드는 인터페이스이지만 인터페이스가 아니라 타입제한이댜)
type Int1 interface {
	~int32 | ~int64 | ~int
	// ~ -> 별칭 타입도 포함가능
}
type Int2 interface {
	int32 | int64 | int
	// ~ -> 별칭 타입도 포함가능
}
type Float1 interface {
	~float32 | ~float64
	// ~ -> 별칭 타입도 포함가능
}

// 위에랑 같은게 아니댜..
type Getter interface {
	Get() int
}

func funcc(a Getter) Getter {
	return a
}

type MyInt int

// // Float
// func funcc1(a Float1) Float1 {
// 	return a
// }
