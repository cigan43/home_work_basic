package main

type Book struct {
	id     uint
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func newBook(id uint, title, author string, year, size int, rate float32) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func main() {
	p := newBook(1, "Война и мир", "Толстой", 2020, 500, 5.8)
}
