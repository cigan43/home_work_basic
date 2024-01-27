package shapes

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	testCases := []struct {
		desc  string
		shape Shape
		want  float64
		er    error
	}{
		{"Rectangle area", Rectangle{4, 5}, 20, nil},
		{"Circle  area", Circle{5}, 78.5, nil},
		{"Triangle area", Triangle{4, 5}, 10, nil},
		{"Zero parament, Rectangle area", Rectangle{0, 5}, 0, errors.New("rectangle.area неправильно введены параметры")},
		{"Zero parametr, Circle  area", Circle{0}, 0, errors.New("circle.area неправильно введены параметры")},
		{"Zero parametr, Triangle area", Triangle{0, 5}, 0, errors.New("triangle.area неправильно введены параметры")},
		{"Negative meaning, Rectangle area", Rectangle{-4, 5}, 0, errors.New("rectangle.area неправильно введены параметры")},
		{"Negative meaning, Circle  area", Circle{-5}, 0, errors.New("circle.area неправильно введены параметры")},
		{"Negative meaning, Triangle area", Triangle{-4, 5}, 0, errors.New("triangle.area неправильно введены параметры")},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := tC.shape.Area()
			if err != nil {
				assert.Equal(t, tC.er, err)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestCalculateArea(t *testing.T) {
	testCases := []struct {
		desc  string
		inter any
		want  float64
		er    error
	}{
		{
			desc:  "Good circle",
			inter: Shape(Circle{Radius: 33.0}),
			want:  3419.46,
			er:    nil,
		},
		{
			desc:  "Good rectangle",
			inter: Shape(Rectangle{Width: 5, Height: 8}),
			want:  40.0,
			er:    nil,
		},
		{
			desc:  "Good triangle",
			inter: Shape(Triangle{Base: 5, Height: 7}),
			want:  17.5,
			er:    nil,
		},
		{
			desc:  "Error interface",
			inter: []struct{ square float64 }{{10.0}},
			want:  0.0,
			er:    errors.New("не интерфайс Shape"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := CalculateArea(tC.inter)
			if err != nil {
				fmt.Println(err, got)
				assert.Equal(t, tC.er, err)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}
