package main

<<<<<<< Updated upstream
=======
import (
	"encoding/json"
	"fmt"

	"github.com/cigan43/home_work_basic/hw09_serialize/module"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// - Реализуйте структуру Book со следующими полями: ID, Title, Author, Year, Size, Rate (может быть дробным).
// - Реализуйте для нее интерфейсы Marshaller и Unmarshaller  из пакета json.
// - Составьте protobuf спецификацию для Book
// - Реализуйте для нее интерфейс Message из пакета proto.
// - Напишите фукции выполняющие сериализацию/десериализацию слайса объектов.
// - Напишите юнит тесты на реализованные функции;
type ProtoBook module.Book

type Book struct {
	Id     int64   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int64   `json:"year"`
	Size   int64   `json:"size"`
	Rate   float64 `json:"rate"`
}

type Marshaller interface {
	Marshal() ([]byte, error)
	MessageMarshaler() ([]byte, error)
}

type Unmarshaller interface {
	Unmarshal([]byte) error
	MessageUnmarshal([]byte) error
}

func (b *Book) Marshal() ([]byte, error) {
	bookMarshal, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return bookMarshal, nil
}

type Message interface {
	String() string
	ProtoReflect()
	GetID() string
	GetRate() float32
	GetYear() uint32
	GetSize() uint32
	GetTitle() string
	GetAuthor() string
}

type ProtoMessage interface { ///зачем это надо я не понимаю
	ProtoReflect() Message
}

func (b *Book) Unmarshal(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

// type MessageUnmarshaler interface {
// 	MessageUnmarshaler([]byte) error
// }

type MessageMarshaler interface {
	MessageMarshaler() ([]byte, error)
}

func (b *ProtoBook) MessageMarshaler() ([]byte, error) {
	bookMarshal, err := proto.Marshal(protoreflect.ProtoMessage(b))
	if err != nil {
		return nil, err
	}
	return bookMarshal, nil
}

func (b *ProtoBook) MessageUnmarshal(bytes []byte) error {
	return proto.Unmarshal(bytes, b)
}

func SliceBookMarshel(ss []Book) []byte {
	var bookbyte []byte
	for i := range ss {
		a, err := ss[i].Marshal()
		if err != nil {
			fmt.Println(err)
		}
		append(bookbyte, a...)
	}
	return bookbyte
}

func SliceBookMarshelProto(ss []module.Book) []byte {
	var bookbyte []byte
	for i := range ss {
		a, err := ss[i].MessageMarshaler()
		if err != nil {
			fmt.Println(err)
		}
		append(bookbyte, a...)
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

>>>>>>> Stashed changes
func main() {
	// Place your code here.
}
