package geometry

import (
	"errors"
	"math"
)

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height, Base float64
}

func (r Rectangle) Area() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return 0.0, errors.New("rectangle.area неправильно введены параметры")
	}
	return r.Width * r.Height, nil
}

func (c Circle) Area() (float64, error) {
	if c.Radius <= 0 {
		return 0.0, errors.New("circle.area неправильно введены параметры")
	}
	return math.Pow(c.Radius, 2) * 3.14, nil
}

func (t Triangle) Area() (float64, error) {
	if t.Base <= 0 || t.Height <= 0 {
		return 0.0, errors.New("triangle.area неправильно введены параметры")
	}
	return 0.5 * t.Height * t.Base, nil
}
