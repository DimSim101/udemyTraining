package main

import "fmt"

// This is a function that initializes a count to 0 and returns a function.
// This way we can init the variable and manipulate it with the returned function.
// This means we only call wrapper() once, and therefor x is initialized to its
// default 0 once (when called) and then incremented from there using the returned
// function. See below for example.
func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	// Here we get the function result that is returned from wrapper(),
	// and calling wrapper() also initializes the int which increment() accesses.
	// Thus, calling increment() only calls the inner function and doesn't
	// reset the value of x.
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
