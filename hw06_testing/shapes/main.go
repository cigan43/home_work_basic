package shapes

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() (float64, error)
}
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

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0.0, errors.New("не интерфайс Shape")
	}

	area, err := shape.Area()
	if err != nil {
		return 0.0, fmt.Errorf("%w", err)
	}
	return area, nil
}

func main() {
	// Place your code here.
	c := Circle{Radius: 33.0}
	r := Rectangle{Width: 0, Height: 8}
	t := Triangle{Base: 5, Height: 7}

	areaC, err := calculateArea(Shape(c))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(areaC)
	}

	areaR, err := calculateArea(Shape(r))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(areaR)
	}

	areaT, err := calculateArea(Shape(t))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(areaT)
	}
}
