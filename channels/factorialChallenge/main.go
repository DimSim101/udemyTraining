package main

import (
	"fmt"
)

func main() {
	var number int = 4
	fmt.Println("Number value before factorial is:", number)
	factorialChan := factorial(number)
	for n := range factorialChan {
		number = n
		fmt.Println("Number value is now:", number)
	}

	number = 6
	fmt.Println("Number value before factorial is:", number)
	factorialChan = factorial(number)
	for n := range factorialChan {
		number = n
		fmt.Println("Number value is now:", number)
	}

	number = 8
	fmt.Println("Number pre increment is", number)
	incrementChan := incrementByValue(number, 20)
	for n := range incrementChan {
		number = n
		fmt.Println("Number value post increment is:", number)
	}
}

func factorial(number int)<-chan int {
	out := make(chan int)
	go func() {
		var factorialValue int = 1
		for i := number; i > 0; i-- {
			factorialValue *= i
		}
		out <- factorialValue
		close(out)
	}()
	return out
}

func incrementByValue(number int, value int) <-chan int {
	out := make(chan int)
	go func() {
		var incrementedValue int = number
		for i := 0; i < value; i++ {
			incrementedValue += 1
		}
		out <- incrementedValue
		close(out)
	}()
	return out
}