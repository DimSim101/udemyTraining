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

	myName, err := getName("BLAH")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myName)
	}

	myName, err = getName("David")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myName)
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

func getName(name string) (string, error) {
	if name == "David" {
		// errors returned are always strings.
		return "", errors.New("Hello me")
	}
	return "Your name is " + name, nil
}
