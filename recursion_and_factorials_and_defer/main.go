package main

import (
	"fmt"
	//"time"
	"time"
)

func factorial(x int) int {
	total := 1
	for x > 0 {
		total *= x
		x--
	}
	return total
}

func recursiveFactorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * recursiveFactorial(x-1)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Factorial of %d: %d = %d Recursive Factorial \n",
			i, factorial(i), recursiveFactorial(i))

	}

	// Defer the recursive calculation until after the normal one has finished
	defer fmt.Println("RECURSIVE FACTORIAL OF 42 = ", recursiveFactorial(42))
	// Simulate some kind of wait condition here
	// For example, we might defer some function so that it will always happen
	// after the result of some async function has returned.
	// This would be useful for multithreading / concurrency stuff, as well as
	// doing stuff in the background (same thing) like a web request.
	fmt.Println("FACTORIAL OF 42 = ", factorial(42))
	time.Sleep(5000000000)
}
