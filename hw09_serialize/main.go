package main

import (
	"encoding/json"
	"fmt"
	// "google.golang.org/protobuf/proto"
)

// - Реализуйте структуру Book со следующими полями: ID, Title, Author, Year, Size, Rate (может быть дробным).
// - Реализуйте для нее интерфейсы Marshaller и Unmarshaller  из пакета json.
// - Составьте protobuf спецификацию для Book
// - Реализуйте для нее интерфейс Message из пакета proto.
// - Напишите фукции выполняющие сериализацию/десериализацию слайса объектов.
// - Напишите юнит тесты на реализованные функции;

type Book struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int64   `json:"year"`
	Size   int64   `json:"size"`
	Rate   float64 `json:"rate"`
}

type Marshaller interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaller interface {
	UnmarshalJSON([]byte) error
}

func (b *Book) MarshalJSON() ([]byte, error) {
	fmt.Println(b)
	bookMarshal, err := json.Marshal(b)
	fmt.Println(err, "-------")
	if err != nil {
		return nil, err
	}
	fmt.Println(bookMarshal, "+++")
	return bookMarshal, nil
}

func (b *Book) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, b)
	if err != nil {
		return err
	}
	return nil
}

// type MessageUnmarshaler interface {
// 	MessageUnmarshaler([]byte) error
// }

// type MessageMarshaler interface {
// 	MessageMarshaler() ([]byte, error)
// }

// func (b *Book) MessageMarshaler() ([]byte, error) {
// 	bookMarshal, err := proto.Marshal(b)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bookMarshal, nil
// }

// func (b *Book) UnmarshalJSON(bytes []byte) error {
// 	return proto.Unmarshal(bytes, b)
// }

func SliceBookMarshel(ss []Book) []byte {
	var bookbyte []byte
	for i := range ss {
		a, err := ss[i].MarshalJSON()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(a)
		// append(bookbyte, )
	}
	return bookbyte
}

// func SliceBookUnMarshel(ss []Book) []byte {
// 	var b []Book
// 	bookbyte := []byte
// 	for i := range ss {
// 		fmt.Println(i)
// 		// append(bookbyte, i.UnmarshalJSON(byte(i), b))
// 	}
// 	return bookbyte
// }

func main() {
	myBook := []Book{
		{
			ID:     1,
			Title:  "Pupkin",
			Author: "Pup",
			Year:   2000,
			Size:   300,
			Rate:   6.7,
		},
		{
			ID:     2,
			Title:  "Rfr",
			Author: "kfd;sk",
			Year:   200,
			Size:   100,
			Rate:   5.9,
		},
	}

	a := SliceBookMarshel(myBook)
	// a := SliceBookUnMarshel(myBook)
	fmt.Println(a)
	// proto.Message()
}
