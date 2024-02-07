package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) area() float64 {
	return t.base * t.height / 2
}

func calculateArea(s Shape) (float64, error) {
	area := s.area()
	if area < 0 {
		return 0, fmt.Errorf("обьект не реализует интерфейс Shape")
	}
	return area, nil
}

func main() {
	circle := Circle{5}
	rectangle := Rectangle{10, 5}
	triangle := Triangle{8, 6}

	circleArea, err := calculateArea(circle)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Площадь круга:", circleArea)
	}

	rectangleArea, err := calculateArea(rectangle)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Площадь прямоугольника:", rectangleArea)
	}

	triangleArea, err := calculateArea(triangle)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Площадь треугольника:", triangleArea)
	}
}
