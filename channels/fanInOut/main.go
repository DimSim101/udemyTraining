package main

import (
	"fmt"
	"sync"
)

func main() {

	const MAX_CHANNELS  = 42
	var intSlice []uint64
	for j := 0; j < 5; j++ {
		for i := 0; i < 20; i++ {
			intSlice = append(intSlice, uint64(i))
		}
	}

	ints := gen(intSlice...) // create a channel of ints

	// FAN OUT - PROCESS SINGLE CHANNEL IN MULTIPLE FUNCTIONS / GOROUTINES
	// Create a slice of <-chan uint64 (channels we can receive from)
	// Which we can then add the new channels to when starting another
	// factorial goroutine. This allows us to run as many goroutines as
	// we want to handle the factorial calculation on any number of uint64's.
	var chanSlice []<-chan uint64
	for i := 0; i < MAX_CHANNELS; i++ {
		newChan := factorial(ints) // start calculating factorial of ints in channel
		chanSlice = append(chanSlice, newChan) // append the channel to the slice
	}

	// FAN IN - GRAB THE CHANNELS AND MERGE THE THEIR VALUES INTO ONE CHANNEL
	// Now we can merge all the channels we created to handle the factorial
	// calculations. Since theyre all inside chanSlice, we can just use
	// the ... syntax to expand them all (since merge is a variadic function).
	for n := range merge(chanSlice...) {
		fmt.Println(n) // Print the factorial value taken from the channel
	}
}

func merge(chans ...<-chan uint64) <-chan uint64{
	out := make(chan uint64)
	var semaphore sync.WaitGroup
	semaphore.Add(len(chans))

	for _, channel := range chans {
		go func(ch <-chan uint64) {
			for n := range ch {
				out <- n
			}
			semaphore.Done()
		}(channel)
	}

	go func() {
		semaphore.Wait()
		close(out)
	}()

	return out
}

func gen(nums ...uint64) <-chan uint64{
	out := make(chan uint64)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan uint64)<-chan uint64 {
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