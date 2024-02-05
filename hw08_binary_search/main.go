package main

// 1.Вычисляется среднее значение массива.
// 2.Значение полученного элемента сравнивается с искомым (ключом).
// 		Если оно меньше, дальнейший поиск для возрастающего массива выполняется слева от центрального элемента.
// 		В противном случае ключ ищется справа.
// 3.В случае совпадения среднего значения с искомым поиск прекращается.
// 		Пользователю возвращается индекс совпавшего элемента.
// 4.Дальнейшие итерации первых двух шагов повторяются вплоть до нахождения ключа.
// 5.Если в результате очередного деления остался лишь один элемент,
// 		и он не совпадает с искомым, пользователю возвращается значение -1.

// принимать масси и значения поиска

// arr := []int{1, 3, 6, 7, 13, 15, 16, 19, 24, 28, 29}
// search = 19

import (
	"errors"
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

func search(min, max, search_item int, sort_slice []int) (int, error) {
	half := (max-min)/2 + min
	fmt.Println("!", half, min, max, sort_slice)
	fmt.Println(half, sort_slice[half], search_item)
	if sort_slice[half] == search_item {
		return half, nil
	} else if sort_slice[half] > search_item { // искомый элемент с лева
		fmt.Println("-", min, half, sort_slice[half])
		search(min, half, search_item, sort_slice)
	} else if sort_slice[half] < search_item { // искомый элемент с права
		fmt.Println("+", half, max)
		search(half, max, search_item, sort_slice)
	}
	return 0, errors.New("ненаншли елемента в заданном массиве")
}

func main() {
	a := []int{1, 2, 5, 7, 3, 4, 6, 9, 8}
	searchIndex := 6
	s := sortSlice(a)
	index, err := search(0, len(s), searchIndex, s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(index)
	}
}
