package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

func main() {
	p1 := person{"David", "Aaron", 25}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)


	// Both initialization techniques for a struct result in a struct
	// that has the zero values for all the variables it contains.

	// initialize struct using var default nil value -- same as below
	var p2 person
	fmt.Println(p2)

	// initialize struct to empty value -- same as above
	p3 := person{}
	fmt.Println(p3)
}
