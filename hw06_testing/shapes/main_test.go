package shapes

import (
	"errors"
	"testing"

	"github.com/cigan43/home_work_basic/hw06_testing/shapes"
	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	testCases := []struct {
		desc  string
		shape shapes.Shape
		want  float64
		err   error
	}{
		{"Rectangle area", shapes.Rectangle{4, 5}, 20, nil},
		{"Circle  area", shapes.Circle{5}, 78.5, nil},
		{"Triangle area", shapes.Triangle{4, 5}, 10, nil},
		{"Zero parament, Rectangle area", shapes.Rectangle{0, 5}, 0, errors.New("rectangle.area неправильно введены параметры")},
		{"Zero parametr, Circle  area", shapes.Circle{0}, 0, errors.New("rectangle.area неправильно введены параметры")},
		{"Zero parametr, Triangle area", shapes.Triangle{0, 5}, 0, errors.New("rectangle.area неправильно введены параметры")},
		{"Negative meaning, Rectangle area", shapes.Rectangle{-4, 5}, 0, errors.New("rectangle.area неправильно введены параметры")},
		{"Negative meaning, Circle  area", shapes.Circle{-5}, 0, errors.New("rectangle.area неправильно введены параметры")},
		{"Negative meaning, Triangle area", shapes.Triangle{-4, 5}, 0, errors.New("rectangle.area неправильно введены параметры")},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := tC.shape.Area()
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}
