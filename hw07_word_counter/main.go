package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordSplit(row string) []string {

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_!.' '", r)
	}

	words := strings.FieldsFunc(text, splitFunc)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}

}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_!.' '", r)
	}

	words := strings.FieldsFunc(text, splitFunc)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}
