package main

import (
	"fmt"
)

const (
	black = "#"
	white = " "
)

func CreateString(countColumn int, firstStep int) string {
	var readyString string
	for i := 0; i < countColumn; i++ {
		if firstStep%2 == 0 {
			if len(readyString)%2 == 0 {
				readyString += black
			} else {
				readyString += white
			}
		} else {
			if len(readyString)%2 == 0 {
				readyString += white
			} else {
				readyString += black
			}
		}
	}
	return readyString
}

func main() {
	var row, column int
	var board string
	fmt.Scanf("%dx%d", &row, &column)

	for i := 0; row > i; i++ {
		board = board + "\n" + CreateString(column, i)
	}
	fmt.Println(board)
}
