package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// fmt.Fscanln(reader, &a, &b)
	fmt.Fscanf(reader, "%d\n%d\n", &a, &b)
	// fmt.Fscanln → fmt.Fscanf: 입력을 줄바꿈(\n)을 기준으로 정확히 두 줄의 입력을 받을 수 있도록 fmt.Fscanf로 변경했습니다. 이렇게 하면 줄바꿈까지 처리할 수 있어 입력 문제를 방지할 수 있습니다.

	fmt.Println(a * (b % 10))
	fmt.Println(a * ((b % 100) / 10))
	fmt.Println(a * (b / 100))
	fmt.Println(a * b)
}
