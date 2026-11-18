package main

import (
	"fmt"
)

type Area interface {
	Area() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func printArea(a Area) {
	fmt.Println("Area:", a.Area())
}


