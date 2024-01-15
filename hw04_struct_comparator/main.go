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

func (b *Book) SetID(newid uint) {
	b.id = newid
}

func (b *Book) Rate() float32 {
	return b.rate
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) ID() uint {
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

func checkStruct(book Book) bool {
	defaultBook := Book{}
	return book == defaultBook
}

func main() {
	var bookValue uint8
	userBook, userBook1 := Book{}, Book{}
	fmt.Println("Вам нужно ввести данные по книге1 разделенные через пробел: id title author year size rate")
	fmt.Scanf("%d %s %s %d %d %g",
		&userBook.id, &userBook.title, &userBook.author, &userBook.year, &userBook.size, &userBook.rate)
	fmt.Println("Вам нужно ввести данные по книге2 разделенные через пробел: id title author year size rate")
	fmt.Scanf("%d %s %s %d %d %g",
		&userBook1.id, &userBook1.title, &userBook1.author, &userBook1.year, &userBook1.size, &userBook1.rate)

	if checkStruct(userBook) || checkStruct(userBook1) {
		fmt.Println("Неправильно заполнили данные по книгам, берем значение по умолчанию")
		userBook.SetID(1)
		userBook.SetTitle("Война и мир")
		userBook.SetAuthor("Толстой")
		userBook.SetYear(2020)
		userBook.SetSize(500)
		userBook.SetRate(5.8)

		userBook.SetID(2)
		userBook.SetTitle("Сторожевая башня")
		userBook.SetAuthor("Стругадские")
		userBook.SetYear(2008)
		userBook.SetSize(300)
		userBook.SetRate(8)
	}

	details := []FieldComapre{year, size, rate}
	for bookIndex, bookType := range details {
		fmt.Print("[", bookIndex, "] - ", bookType.String(), "\n")
	}
	fmt.Fscanln(os.Stdin, &bookValue)
	fmt.Println("Вы выбрали:", details[bookValue])
	comparator := NewComparator(details[bookValue])
	result := comparator.Compare
	fmt.Println("Результат сравнения двух книг:", result(&userBook, &userBook1))
}
