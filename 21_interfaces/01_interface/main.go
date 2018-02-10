package main

import (
	"fmt"
	"math"
)

// Square type
type Square struct {
	side float64
}

// attach area method to the Square type
func (z Square) area() float64 {
	return z.side * z.side
}

// Circle type
type Circle struct {
	radius float64
}

// attach area method to the Circle type
func (z Circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

// Shape interface is implemented by any type with an area method
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(s)
	c := Circle{6}
	info(c)
}
