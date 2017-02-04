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

	names := []string{"BLAH", "David", "Mandy", "Namu", "Dave", "Davo"}

	for _, name := range names {
		myName, err := getName(name)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(myName)
		}

		myName, err = getNameIdiomaticErrors(name)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(myName)
		}
	}
}

func factorial(num int) (int, error) {
	if num < 0 {
		return 0, fmt.Errorf("Error: Cannot call factorial() on negative number: %v.", num)
	}

	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result, nil
}

func getName(name string) (string, error) {
	var err error // Here we can create an error type which will be nil
	// and hence can just set its value / leave it as the default (nil)
	// and return it.

	// We can also create any errors which we might use repeatedly
	var ErrMum = errors.New("Hi mum")

	switch name {
	case "David":
		// errors returned by errors.New() are always represented by strings.
		err = errors.New("Hello me")
		return "", err
	case "Mandy":
		return "", ErrMum
	case "Namu":
		return "", ErrMum
	default:
		return "Your name is " + name, err // error value is nil here
	}
	// This can be here after the switch or inside a default case - behaviour identical
	// return "Your name is " + name, err // error value is nil here
}

func getNameIdiomaticErrors(name string) (string, error) {
	// Here we can create any errors we want to use repeatedly.
	var (
		// Simple error string using errors.New()
		ErrMum = errors.New("Hi mum idiomatic")
		// Formatter error string using fmt.Errorf (calls errors.New()
		// behind the scenes after formatting the string).
		ErrMe = fmt.Errorf("Error with name: %v. Hello me idiomatic", name)
		ErrYou = fmt.Errorf("Error with name: %v. Hellloo youu idiomatic", name)
	)
	switch name {
	case "David":
		return "", ErrMe
	case "Mandy":
		return "", ErrMum
	case "Namu":
		return "", ErrMum
	case "Dave":
		return "", ErrYou
	case "Davo":
		return "", ErrYou
	default:
		// idiomatic returning of error that is clearly nil (no error).
		return "Your idiomatic name is " + name, nil
	}
}
