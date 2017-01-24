package main

import (
	"fmt"
)

func factorial (x int) int {
	total := 1
	for x > 0 {
		total *= x
		x--
	}
	return total
}

func recursiveFactorial (x int) int {
	if x == 0 {
		return 1
	}
	return x * recursiveFactorial(x-1)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Factorial: %d = %d Recursive Factorial \n",
			factorial(i), recursiveFactorial(i))
	}
}
