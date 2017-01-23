package main

import (
	"fmt"
	"strconv"
)

func main() {
	strings(0, "test")
	strings(1, "one", "two")
	strings(2, "1", "2", "3", `4`)
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