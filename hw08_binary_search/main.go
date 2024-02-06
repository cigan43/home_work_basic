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
	case v > searchItem:
		return search(min, half, searchItem, sortSlice)
	case v < searchItem:
		return search(half, max, searchItem, sortSlice)
	default:
		return 0, errors.New("ненаншли елемента в заданном массиве")
	}
}

// }
// if sortSlice[half] == searchItem {
// 	return half, nil
// } else if sortSlice[half] > searchItem { // искомый элемент с лева
// 	return search(min, half, searchItem, sortSlice)
// } else if sortSlice[half] < searchItem { // искомый элемент с права
// 	return search(half, max, searchItem, sortSlice)
// }
// return 0, errors.New("ненаншли елемента в заданном массиве")
// }

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
