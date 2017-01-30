package main

import (
	"fmt"
	"math"
)

// This is an empty interface and everything implements it (as it has no methods).
// Thus, we can use empty interfaces to do interesting things like store variables
// of any type.
type empty interface{}

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

// This function has a value receiver (of type square). Unlike pointer receivers,
// value receivers can be used by both the type directly as well as a pointer to
// the type.
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

func info(z shape) {
	fmt.Printf("%T \n", z)
	fmt.Println(z)
	fmt.Println(z.area())
	fmt.Println(z.width())
}

// Here we can see with an empty interface parameter the function can take in
// values of any type.
func emptyInterface(a interface{}) {
	fmt.Printf("%T = %v \n", a, a)
}

func emptyInterfaceVariadic(a ...interface{}) {
	fmt.Printf("%T = %v \n", a, a)
}

// This function has a pointer receiver of type *square. It can only be called
// using a pointer.
// Courtesy of Bill Kennedy:
// Receivers	Values
// (t T)	T and *T
// (t *T)	*T
func (s *square) doubleSide() float64 {
	return 2 * s.sideLen
}

type shape2 interface {
	doubleSide() float64
}

func info2(z shape2) {
	fmt.Println(z.doubleSide())
}

func main() {
	fmt.Println("Starting receiver example...")
	valueReceiver := square{5}
	fmt.Println("Created square:", valueReceiver)
	fmt.Println("Calling function with pointer receiver using address of square:")
	info2(&valueReceiver) // calls doubleSide() which takes in a variable
	// of type pointer to square (*square).
	fmt.Println("Calling function with value receiver using square variable directly:")
	info(valueReceiver)
	fmt.Println("Calling function with value receiver using address of square:")
	info(&valueReceiver)

	fmt.Println()
	s := square{4}
	info(s)

	c := circle{3}
	info(c)

	// This is a slice of empty interfaces and thus the types inside can be
	// anything.
	emptySlice := []empty{s, c, "Testing", 420, 'a', []int{0, 1, 2, 3}}
	fmt.Println(emptySlice)

	ints := []int{0, 1, 2, 3}
	emptyInterface(s)
	emptyInterface(c)
	emptyInterface(emptySlice)
	emptyInterface(ints)

	interfaces := []interface{}{"Blah", "Blah2", 42, []int{9, 8, 7}}
	fmt.Println(interfaces)
	// Here, the slice of interface{} is the argument and thus only
	// takes up one entry in the printed interfaces variable.
	emptyInterfaceVariadic(s, c, emptySlice, ints, interfaces)

	// The blah... syntax cannot be used with other arguments because it
	// expands to an unlimited number of args. We cannot do stuff like
	// emptyInterfaceVariadic(s, c, interfaces...)
	// This causes each value inside the array of inteface{} to get broken
	// into individual arguments. Thus, each element inside the array is
	// a separate element within the array of interfaces that gets printed
	// (rather than being chunked together as above).
	emptyInterfaceVariadic(interfaces...)

}
