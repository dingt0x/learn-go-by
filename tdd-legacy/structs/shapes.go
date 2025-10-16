package main

import "math"

func Perimeter(rectangle Rectangle) float64 {
    return (rectangle.height + rectangle.width) * 2
}

func Area(rectangle Rectangle) float64 {
    return rectangle.width * rectangle.height
}

type Rectangle struct {
    width, height float64
}

type Circle struct {
    Radius float64
}

func (r Rectangle) Area() float64 {
    return r.height * r.width
}

func (c Circle) Area() float64 {
    return c.Radius * c.Radius * math.Pi
}

type Shape interface {
    Area() float64
}

type Triangle struct {
    Base, Height float64
}

func (t Triangle) Area() float64 {
    return t.Base * t.Height * 0.5
}
