package main

import (
	"fmt"
	"github.com/DimSim101/udemyTraining/03_visibility_and_ptrs/vis"
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
	fmt.Println(p2.age)

	// initialize struct to empty value -- same as above
	p3 := person{}
	fmt.Println(p3)
	fmt.Println(p3.age)
	fmt.Println()

	// Pointer to p1 is just a pointer to the first variable in the struct
	// Presumably this is just the start of the memory chunk in which
	// all the struct variables are placed
	fmt.Printf("%p \n", &p1)
	fmt.Printf("%p - %p - %p \n\n", &p1.first, &p1.last, &p1.age)

	fmt.Printf("%p \n", &p2)
	fmt.Printf("%p - %p - %p \n\n", &p2.first, &p2.last, &p2.age)

	fmt.Printf("%p \n", &p3)
	fmt.Printf("%p - %p - %p \n", &p3.first, &p3.last, &p3.age)

	// visibility rules apply to naming structs just like any other variable
	// capital first letter = visible outside package.
	test := vis.VisiblePerson{}
	fmt.Println(test)
}
