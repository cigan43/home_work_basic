package main

import (
	"fmt"
	"strings"
)

const (
	black = "#"
	white = " "
)

func CreateString(count_column int, first int) string {
	var ready_string string
	for i := 0; i < count_column; i++ {
		if first%2 == 0 {
			if len(ready_string)%2 == 0 {
				ready_string = ready_string + black
			} else {
				ready_string = ready_string + white
			}
		} else {
			if len(ready_string)%2 == 0 {
				ready_string = ready_string + white
			} else {
				ready_string = ready_string + black
			}

		}
	}
	return ready_string
}

func main() {
	var row, column int
	var board string
	var size string
	fmt.Println("В ведите количество стор и ")
	// fmt.Scan(&row, &column)
	fmt.Scanf("%d", &size)
	&row, &column = strings.Split(size, "x")

	for i := 0; row > i; i++ {
		board = board + "\n" + CreateString(column, i)
	}
	fmt.Println(board)
}
