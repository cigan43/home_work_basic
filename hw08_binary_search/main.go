package main

import (
	"errors"
	"fmt"
	"sort"
)

func sortSlice(ar []int) []int {
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	return ar
}

func search(min, max, searchItem int, sortSlice []int) (int, error) {
	half := (max-min)/2 + min

	switch v := sortSlice[half]; {
	case v == searchItem:
		return half, nil
	case half-1 == min || half+1 == max:
		return 0, errors.New("ненаншли елемента в заданном массиве")
	case v > searchItem:
		return search(min, half, searchItem, sortSlice)
	case v < searchItem:
		return search(half, max, searchItem, sortSlice)
	default:
		return 0, errors.New("ненаншли елемента в заданном массиве")
	}
}

func main() {
	a := []int{1, 2, 5, 7, 3, 4, 6, 9, 8}
	searchIndex := 999
	s := sortSlice(a)
	index, err := search(0, len(s), searchIndex, s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(index)
	}
}
