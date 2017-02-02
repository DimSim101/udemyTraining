package main

import "fmt"

func main() {

	var intSlice []uint64
	for j := 0; j < 5; j++ {
		for i := 0; i < 20; i++ {
			intSlice = append(intSlice, uint64(i))
		}
	}

	ints := gen(intSlice...)
	values := factorial(ints)
	count := 0
	for n := range values {
		fmt.Println("Factorial of", intSlice[count], "is:", n)
		count++
	}
}

func gen(nums ...uint64) <-chan uint64 {
	out := make(chan uint64)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	go func() {
		for n := range in {
			out <- calculateFactorial(n)
		}
		close(out)
	}()
	return out
}

func calculateFactorial(num uint64) uint64 {
	var factorialValue uint64 = 1
	for i := num; i > 0; i-- {
		factorialValue *= i
	}
	return factorialValue
}
