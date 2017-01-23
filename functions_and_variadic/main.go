package main

import (
	"fmt"
	"strconv"
)

func main() {


	var data []string
	test := "test"

	one := "one"
	two := "two"

	first := "1"
	second := "2"
	third := "3"
	fourth := `4`

	strings(0, test)
	strings(1, one, two)
	strings(2, first, second, third, fourth)

	data = []string {"Hello", "My", "Name", "Is", "King", "David"}
	// calling the function as strings(data) fails as the variable data
	// is of type []string (a slice of strings) and gets passed into the
	// string function as a single argument (under the variable strings).
	// This causes a type comparison error as []string and string are not
	// the same. Hence, we must unpack the individual elements of the slice
	// and pass them in as separate arguments to the variadic argument strings
	// in order for this to work. Hence the syntax strings(3, data...)
	strings(3, data...)

	// However, if we swap the type of strings from variadic to be a
	// slice of strings, then the behaviour is identical (the amount of
	// variables in the slice can vary, theyre all of the same type etc.)
	// Even the syntax for looping over them is the same, as the variadic
	// argument (or variable) gets converted into a slice of the declared
	// type anyway. Probably easier to never use variadic functions and
	// instead just pass in a slice of whatever type you want. Maybe they
	// are better for something like multiple interfaces, but even then why
	// not just use a slice of interfaces?
	data = []string{test}
	stringsTwo(0, data)

	data = []string{one, two}
	stringsTwo(1, data)

	data = []string{first, second, third, fourth}
	stringsTwo(2, data)

	data = []string {"Hello", "My", "Name", "Is", "King", "David"}
	stringsTwo(3, data)

	// stringsThree can have any number of ints and strings passed in via
	// each argument being a slice of their respective type
	dataInts := []int {1, 2, 3, 4, 420}
	stringsThree(5, data, dataInts)

}

// Note the below doesn't work as the variadic argument must be the last arg
// func strings(strings ...string, number int) {
// instead it must be func strings(number int, strings ...string) {
// This makes sense as a varying number of arguments cannot have anything after
// it or it will assume they are part of the varying list
// This also means we cannot have a function with two variadic arguments i.e.
// func strings(number int, integers ...int, strings ...string) { will throw an
// error (as the string arguments after the integers get passed in as ints to
// the variadic argument integers (and then fail on type)).

// The variadic argument strings returns a slice (a list) of varying size
// containing strings (slices use an underlying array to store the data)
func strings(number int, strings ...string) {

	fmt.Println()
	fmt.Println("Test", strconv.Itoa(number), "started!")

	for _, value := range strings {
		fmt.Println(value)
	}
	fmt.Println(strings)
	fmt.Printf("Strings type is %T and number is of type %T \n", strings, number)
}

// alternate method using slice of whatever declared type is in the function
// rather than making the function variadic. This also means we can have
// multiple "variadic" arguments here without being restricted when actually
// using variadic functions (see stringsThree)
func stringsTwo(number int, strings []string) {
	fmt.Println()
	fmt.Println("Test2", strconv.Itoa(number), "started!")

	for _, value := range strings {
		fmt.Println(value)
	}
	fmt.Println(strings)
	fmt.Printf("Strings type is %T and number is of type %T \n", strings, number)
}

// As discussed above, this function can have multiple arguments with a varying
// number of inputs (i.e. strings can have any number of strings inside etc.)
func stringsThree(number int, strings []string, integers []int) {
	fmt.Println()
	fmt.Println("Test3", strconv.Itoa(number), "started!")

	for _, value := range strings {
		fmt.Println(value)
	}
	fmt.Println(strings)

	for _, value := range integers {
		fmt.Println(value)
	}
	fmt.Println(integers)
	fmt.Printf("Strings type is %T, Integers is type %T and number is of type %T \n", strings, integers, number)
}