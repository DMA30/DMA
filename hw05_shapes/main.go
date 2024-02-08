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

func newCircle(radius float64) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func (r *Rectangle) GetWidth() float64 {
	return r.width
}

func (r *Rectangle) GetHeight() float64 {
	return r.height
}

func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}

func (r *Rectangle) SetHeight(height float64) {
	r.height = height
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Triangle struct {
	base   float64
	height float64
}

func newTriangle(base float64, height float64) *Triangle {
	return &Triangle{
		base:   base,
		height: height,
	}
}

func (t Triangle) GetHeight() float64 {
	return t.height
}

func (t Triangle) GetBase() float64 {
	return t.base
}

func (t *Triangle) SetHeight(height float64) {
	t.height = height
}

func (t *Triangle) SetBase(base float64) {
	t.base = base
}

func (t Triangle) area() float64 {
	return t.base * t.height / 2
}

func calculateArea(s any) (float64, error) {
	area, ok := s.(Shape)
	if !ok {
		return 0, fmt.Errorf("не реализует интерфейс Shape")
	}
	return area.area(), nil
}

func main() {
	circle := newCircle(5)
	rectangle := NewRectangle(10, 5)
	triangle := newTriangle(8, 6)

	fmt.Println(calculateArea(circle))
	fmt.Println(calculateArea(rectangle))
	fmt.Println(calculateArea(triangle))
}
