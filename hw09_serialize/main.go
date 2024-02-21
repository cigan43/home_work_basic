package main

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
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
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

type Marshaller interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaller interface {
	UnmarshalJSON([]byte) error
}

func (b *Book) MarshalJSON() ([]byte, error) {
	bookMarshal, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return bookMarshal, nil
}

func (b *Book) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

type MessageUnmarshaler interface {
	MessageUnmarshaler([]byte) error
}

type MessageMarshaler interface {
	MessageMarshaler() ([]byte, error)
}

func (b *Book) MessageMarshaler() ([]byte, error) {
	bookMarshal, err := proto.Marshal(b)
	if err != nil {
		return nil, err
	}
	return bookMarshal, nil
}

func (b *Book) UnmarshalJSON(bytes []byte) error {
	return proto.Unmarshal(bytes, b)
}

func main() {
	// Place your code here.
}
