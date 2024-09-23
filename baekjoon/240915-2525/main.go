package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var hh, mm, time int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n%d", &hh, &mm, &time)

	mins := hh*60 + mm + time
	mm = mins % 60
	hh = mins / 60
	if hh >= 24 {
		hh -= 24
	}

	fmt.Printf("%d %d\n", hh, mm)
}

func myAnswer() {
	var hour, mins, oven int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n%d", &hour, &mins, &oven)

	mins += oven
	if mins >= 60 {
		extraHour := mins / 60
		mins = mins % 60
		hour += extraHour
	}
	if hour >= 24 {
		hour -= 24
	}

	fmt.Printf("%d %d\n", hour, mins)
}

func reference() {
	var hh, mm, time int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n%d", &hh, &mm, &time)
	hh += time / 60
	mm += time % 60

	if mm >= 60 {
		hh += 1
		mm -= 60
	}

	if hh >= 24 {
		hh -= 24
	}

	fmt.Printf("%d %d\n", hh, mm)
}
