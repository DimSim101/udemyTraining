package main

import (
	"fmt"
	"sync"
)

func main() {
	const maxChannels = 42
	var intSlice []uint64
	// Put 1,000,000 values inside the slice.
	for j := 0; j < 100000; j++ {
		// Large factorials to make the computation take some time.
		for i := 30; i < 40; i++ {
			intSlice = append(intSlice, uint64(i))
		}
	}

	// Put all values inside the slice into the channel
	ints := gen(intSlice...) // create a channel of ints

	// FAN OUT - PROCESS SINGLE CHANNEL IN MULTIPLE FUNCTIONS / GOROUTINES
	// Create a slice of <-chan uint64 (channels we can receive from)
	// Which we can then add the new channels to when starting another
	// factorial goroutine. This allows us to run as many goroutines as
	// we want to handle the factorial calculation on any number of uint64's.
	chanSlice := make([]<-chan uint64, maxChannels) // Create a slice of
	// channels using make since we know how many channels we want (defined
	// as a const above).
	for i := 0; i < maxChannels; i++ {
		newChan := factorial(ints) // start calculating factorial of ints in channel
		chanSlice[i] = newChan     // Put the channel in the slice
	}

	// FAN IN - GRAB THE CHANNELS AND MERGE THE THEIR VALUES INTO ONE CHANNEL
	// Now we can merge all the channels we created to handle the factorial
	// calculations. Since theyre all inside chanSlice, we can just use
	// the ... syntax to expand them all (since merge is a variadic function).
	var count int
	for n := range merge(chanSlice...) {
		count++
		fmt.Println(count, ":", n) // Print the factorial value taken from the channel
	}
}

func merge(chans ...<-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	var semaphore sync.WaitGroup
	semaphore.Add(len(chans))

	for _, channel := range chans {
		// The reason we have to pass channel into the goroutine below
		// is because the scope of everything still applies even when
		// the routine is off and running.
		// For example, currently on each iteration of the slice of channels,
		// we grab the channel and pass it to the goroutine which then
		// goes off and runs, using that channel to grab values from.
		// If instead we just spun up the goroutine and had
		// for n := range channel, since the scope still applies,
		// at each iteration we are updating channel (moving to the
		// next channel in the slice) and hence all the goroutines running
		// will access that channel. This is not the behaviour we want.
		// Instead we want each goroutine to run for each channel, grab
		// all the values off that channel and load them into the
		// out channel (FAN IN). Hence, we must pass each channel into
		// the goroutine in order for it to be "kept around" so to speak.
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
