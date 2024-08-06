package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return shape.Area(), nil
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 8, Height: 6}

	shapes := []any{circle, rectangle, triangle, "not a shape"}

	for _, shape := range shapes {
		switch v := shape.(type) {
		case Circle:
			fmt.Printf("Круг: радиус %.2f\n", v.Radius)
		case Rectangle:
			fmt.Printf("Прямоугольник: ширина %.2f, высота %.2f\n", v.Width, v.Height)
		case Triangle:
			fmt.Printf("Треугольник: основание %.2f, высота %.2f\n", v.Base, v.Height)
		default:
			fmt.Printf("Ошибка: переданный объект не является фигурой\n")
			continue
		}

		area, err := calculateArea(shape)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Площадь: %.8f\n", area)
		}
	}
}

//ready
