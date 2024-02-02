package main

import (
	"fmt"
	"sort"
)

type Range struct {
	min, max, num int
	nums          []int
}

func sortSlice(ar []int) []int {
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	return ar
}

func main() {
	a := []int{1, 2, 5, 7, 3, 4, 6, 9, 8}
	// searchIndex := 5
	s := sortSlice(a)
	fmt.Println(s)
}
