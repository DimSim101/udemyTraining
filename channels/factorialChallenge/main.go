package main

import (
	"fmt"
)

func main() {
	var number int = 4
	fmt.Println("Number value before factorial is:", number)
	factorialChan := factorial(number)
	for n := range multiplyValues(factorialChan) {
		number = n
		fmt.Println("Number value is now:", number)
	}

	number = 6
	fmt.Println("Number value before factorial is:", number)
	factorialChan = factorial(number)
	for n := range multiplyValues(factorialChan) {
		number = n
		fmt.Println("Number value is now:", number)
	}

	number = 10
	fmt.Println("Number value before factorial is:", number)
	factorialChan = factorial(number)
	for n := range multiplyValues(factorialChan) {
		number = n
		fmt.Println("Number value is now:", number)
	}
}

func factorial(number int)<-chan int {
	out := make(chan int)
	go func() {
		for i := number; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func multiplyValues(c <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		var factorialValue int = 1 // init this to 1 or answer is always
		// multiplied by 0, so always == 0
		for n := range c {
			factorialValue *= n
		}
		out <- factorialValue
		close(out)
	}()
	return out
}