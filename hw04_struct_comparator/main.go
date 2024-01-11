package main

import "fmt"

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

func (b *Book) SetRate(newrate float32) {
	b.rate = newrate
}

func (b *Book) SetSize(newsize int) {
	b.size = newsize
}

func (b *Book) SetYear(newyear int) {
	b.year = newyear
}

func (b *Book) SetTitle(newtitle string) {
	b.title = newtitle
}

func (b *Book) SetAuthor(newauthor string) {
	b.author = newauthor
}

func (b *Book) SetId(newid uint) {
	b.id = newid
}

func (b *Book) GetRate() float32 {
	return b.rate
}

func (b *Book) GetSize() int {
	return b.size
}
func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetId() uint {
	return b.id
}

func main() {
	p := newBook(1, "Война и мир", "Толстой", 2020, 500, 5.8)
	p.SetRate(6)
	fmt.Println(p.rate)
}
