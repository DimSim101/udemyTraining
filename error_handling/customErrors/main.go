package main

import (
	"errors"
	"fmt"
)

func main() {
	for i := -5; i < 5; i++ {
		fmt.Println("Calling factorial() on:", i)
		val, err := factorial(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Factorial result is:", val)
		}
	}
}

func factorial(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Error: Cannot call factorial() on a negative number.")
	}

	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result, nil
}
