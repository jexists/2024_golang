package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a, b, c int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d %d", &a, &b, &c)

	fmt.Println(a + b + c)
}
