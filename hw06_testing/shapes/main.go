package main

import (
	"errors"
	"fmt"

	"github.com/cigan43/home_work_basic/hw05_shapes/geometry"
)

type Shape interface {
	Area() (float64, error)
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
	c := geometry.Circle{Radius: 33.0}
	r := geometry.Rectangle{Width: 0, Height: 8}
	t := geometry.Triangle{Base: 5, Height: 7}

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
