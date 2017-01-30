package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	width() float64
}

type square struct {
	sideLen float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.sideLen * s.sideLen
}

func (s square) width() float64 {
	return s.sideLen
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) width() float64 {
	return c.radius * 2
}

func info (z shape) {
	fmt.Printf("%T \n", z)
	fmt.Println(z)
	fmt.Println(z.area())
	fmt.Println(z.width())
}

func main()  {
	s := square{4}
	info(s)

	c := circle{3}
	info(c)
}