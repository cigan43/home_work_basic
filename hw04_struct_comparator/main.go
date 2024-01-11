package main

import (
	"fmt"
	"os"
)

const (
	year FieldComapre = iota
	size
	rate
)

type FieldComapre uint8

type Comparator struct {
	Type FieldComapre
}

func NewComparator(t FieldComapre) *Comparator {
	return &Comparator{
		Type: t,
	}
}

func (c Comparator) Compare(bookOne, bookTwo *Book) bool {
	switch c.Type {
	case year:
		return bookOne.year > bookTwo.year
	case size:
		return bookOne.size > bookTwo.size
	case rate:
		return bookOne.rate > bookTwo.rate
	default:
		return false
	}
}

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

func (s FieldComapre) String() string {
	switch s {
	case size:
		return "size"
	case rate:
		return "rate"
	case year:
		return "year"
	}
	return "unknown"
}

func main() {
	p := newBook(1, "Война и мир", "Толстой", 2020, 500, 5.8)
	p2 := newBook(2, "Сторожевая башня", "Стругадские", 2008, 300, 8)
	var bookValue uint8
	details := []FieldComapre{year, size, rate}
	for bookIndex, bookType := range details {
		fmt.Print("[", bookIndex, "] - ", bookType.String(), "\n")
	}
	fmt.Fscanln(os.Stdin, &bookValue)
	fmt.Println(bookValue)
	fmt.Println("Вы выбрали:", details[bookValue])
	comparator := NewComparator(details[bookValue])
	result := comparator.Compare
	fmt.Println(result(p, p2))
}
