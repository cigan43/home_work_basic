package shapes

import (
	"errors"
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
				// t.Errorf("ERROR: %v", err)
				assert.Equal(t, tC.er, err)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}
