package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myAnswer()

	referenceAnswer()
}

func myAnswer() {

	var hour, mins int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d", &hour, &mins)

	// 1시간 = 60분
	var oneHour = 60
	// 알림시간 45분
	var alarm = 45

	// 현재 분이 알림시간보다 작은 경우
	if mins < alarm {
		if hour == 0 {
			hour = 24
		}
		hour -= 1
		mins = oneHour - alarm + mins
	} else {
		mins = mins - alarm
	}

	fmt.Printf("%d %d\n", hour, mins)
}

func referenceAnswer() {
	var hour, mins int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d", &hour, &mins)

	mins -= 45 // 45분전
	if mins < 0 {
		mins += 60
		hour -= 1

		if hour < 0 {
			hour = 23
		}
	}
	fmt.Printf("%d %d\n", hour, mins)
}
