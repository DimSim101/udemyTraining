package main

import (
	"fmt"
	"strconv"
)

func main() {

	N := 30
	for i := 1; i <= N; i++ {
		output := calcFizzBuzz(i)
		fmt.Println(output)
	}

}

func calcFizzBuzz(number int) string{
	var output string
	if number%3 == 0 {
		output += "Fizz"
	}

	if number%5 == 0 {
		output += "Buzz"
	}

	if output == "" {
		output = strconv.Itoa(number)
	}

	return output
}

/*
BELOW ARE ALTERNATIVE SOLUTIONS FOR SIMPLICITY / CLARITY

// Above could be rewritten to be as below
// if using the strconv package is not wanted

if output != "" {
	fmt.Println(output)
} else {
	fmt.Println(i)
}


// Inefficient solution that might be easier to read

if i % 3 == 0 && i % 5 == 0 {
	fmt.Println("FizzBuzz")
} else if i % 3 == 0 {
	fmt.Println("Fizz")
} else if i % 5 == 0 {
	fmt.Println("Buzz")
} else {
	fmt.Println(i)
}

*/