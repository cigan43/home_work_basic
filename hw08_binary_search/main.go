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

func search(min, max, search_item int, sort_slice []int) (int, error) {
	half := (max-min)/2 + min
	if sort_slice[half] == search_item {
		return half, nil
	} else if sort_slice[half] > search_item { // искомый элемент с лева
		return search(min, half, search_item, sort_slice)
	} else if sort_slice[half] < search_item { // искомый элемент с права
		return search(half, max, search_item, sort_slice)
	}
	return 0, errors.New("ненаншли елемента в заданном массиве")
}

func main() {
	a := []int{1, 2, 5, 7, 3, 4, 6, 9, 8}
	searchIndex := 7
	s := sortSlice(a)
	index, err := search(0, len(s), searchIndex, s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(index)
	}
}
