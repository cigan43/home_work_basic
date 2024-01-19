package chessboard

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
		if len(readyString)%2 == 0 && firstStep%2 == 0 {
			readyString += black
			continue
		} else if len(readyString)%2 != 0 && firstStep%2 == 0 {
			readyString += white
			continue
		}
		if len(readyString)%2 == 0 && firstStep%2 != 0 {
			readyString += white
			continue
		} else if len(readyString)%2 != 0 && firstStep%2 != 0 {
			readyString += black
			continue
		}
	}
	return readyString
}

func main() {
	var row, column int
	var board string
	fmt.Scanf("%dx%d", &row, &column)

	for i := 0; row > i; i++ {
		board += "\n" + CreateString(column, i)
	}
	fmt.Println(board)
}
