package geometry

import (
	"testing"
	"github.com/cigan43/home_work_basic/hw06_testing/shapes/geometry"
)

func TestAreaTriangle(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestAreaCircle(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestAreaRectangle(t *testing.T) {
	testCases := []struct {
		desc                string
		Width, Height, want float64
		err error
	}{
		{
			desc:   "Good area",
			Width:  4,
			Height: 5,
			want:   20,
			err: nil,
		},
		{
			desc:   "Zero parametr",
			Width:  0,
			Height: 5,
			want: 0,
			err: errors.New("rectangle.area неправильно введены параметры")

		},
		{
			desc:   "negative meaning",
			Width:  4,
			Height: -5,
			want: 0,
			err: errors.New("rectangle.area неправильно введены параметры")
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := (tC.countColumn, tC.firstStep)
			assert.Equal(t, tC.want, got)
		})
	}
}
